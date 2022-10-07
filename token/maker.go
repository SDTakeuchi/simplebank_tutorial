package token

import "time"

// an interface for managing tokens
type Maker interface {
	// creates a new token for a specific usernmame and duration
	// *Payload is neccessary to refresh token (cuz it needs token ID)
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// verifies if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
