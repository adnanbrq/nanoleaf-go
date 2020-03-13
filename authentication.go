package nanoleaf

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NanoAuth authorization
type NanoAuth struct {
	nano *Nanoleaf
}

// addUserResponse mimics the response when adding a new user
type addUserResponse struct {
	Token string `json:"auth_token"`
}

// newNanoAuth returns a new instance of NanoAuth
func newNanoAuth(nano *Nanoleaf) *NanoAuth {
	return &NanoAuth{nano}
}

// Authenticate will try to authenticate and set the token
func (a *NanoAuth) Authenticate() error {
	url := fmt.Sprintf("%s/new", a.nano.url)
	resp, err := a.nano.client.R().Post(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusForbidden {
		return ErrAuthNotReady
	}

	if resp.StatusCode() != http.StatusOK {
		return ErrUnexpectedResponse
	}

	var res addUserResponse
	if err := json.Unmarshal(resp.Body(), &res); err != nil {
		return ErrParsingJSON
	}

	a.nano.SetToken(res.Token)
	return nil
}

// Unauthenticate will try to invalidate current token
func (a *NanoAuth) Unauthenticate() error {
	url := fmt.Sprintf("%s/%s", a.nano.url, a.nano.token)
	resp, err := a.nano.client.R().Delete(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusNoContent {
		return ErrUnexpectedResponse
	}

	a.nano.SetToken("")
	return nil
}
