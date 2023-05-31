package models

import "time"

type TokenDto struct {
	Access_Token string    `json:"access_token"`
	Expiration   time.Time `json:"expiration"`
}
