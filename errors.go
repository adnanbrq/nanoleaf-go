package nanoleaf

import "errors"

var (
	// ErrParsingJSON occurs if there was an error failed to parse json
	ErrParsingJSON = errors.New("Failed to parse JSON")

	// ErrAuthFailed occurs if authentication failed
	ErrAuthFailed = errors.New("Authentication failed")

	// ErrAuthNotReady occurs if the nanoleaf api has not been set to accept new authentications
	ErrAuthNotReady = errors.New("Nanoleaf don't accept new Authentications. Please activate the Authentication process")

	// ErrUnauthorized occurs if a request has been made without a valid token
	ErrUnauthorized = errors.New("Unauthorized. Please authorize before sending requests")

	// ErrUnexpectedResponse occurs if nanoleafs send something else than expected
	ErrUnexpectedResponse = errors.New("Received an unexpected response from Nanoleafs")

	// ErrEffectNotFound occurs if given effect was not found
	ErrEffectNotFound = errors.New("Effect not Found")

	// ErrInvalidVersion occurs if given extControl Version does not match v1
	ErrInvalidVersion = errors.New("Invalid version given. Please use v1")
)
