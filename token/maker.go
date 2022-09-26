package token

import "time"

// an interface for managing tokens
type Maker interface {
	// creates a new token for a specific usernmame and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// verifies if the token is valid or not
	VefifyToken(token string) (*Payload, error)
}