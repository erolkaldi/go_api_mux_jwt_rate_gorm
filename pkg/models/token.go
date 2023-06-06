package models

import "time"

type Token struct {
	Access_Token string    `json:"access_token"`
	Expiration   time.Time `json:"expiration"`
}
