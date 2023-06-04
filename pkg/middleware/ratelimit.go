package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiter *rate.Limiter
	last    time.Time
}

type RateLimiterStore struct {
	store    map[string]*RateLimiter
	capacity int
	duration time.Duration
	mu       sync.Mutex
}

func NewRateLimiterStore() *RateLimiterStore {
	return &RateLimiterStore{
		store:    make(map[string]*RateLimiter),
		duration: 1 * time.Minute,
		capacity: 100,
	}
}

func (s *RateLimiterStore) getLimiter(ip string) *rate.Limiter {
	s.mu.Lock()
	defer s.mu.Unlock()

	limiter, ok := s.store[ip]
	if !ok {
		limiter = &RateLimiter{
			limiter: rate.NewLimiter(rate.Every(s.duration), s.capacity),
		}
		s.store[ip] = limiter
	}
	return limiter.limiter
}

func (s *RateLimiterStore) RateCheckLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		limiter := s.getLimiter(ip)

		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("You exceeded your rate limit"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
