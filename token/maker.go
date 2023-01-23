package token

import (
	"time"
)

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration (the token
	// should only be valid within an amount of time)
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	// return the playload data stored inside the Token's body if valid
	VerifyToken(token string) (*Payload, error)
}