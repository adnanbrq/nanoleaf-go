package nanoleaf

import (
	"fmt"
	"net/http"
)

// NanoIdentity identity
type NanoIdentity struct {
	nano     *Nanoleaf
	endpoint string
}

// newNanoIdentity returns a new instance of NanoIdentity
func newNanoIdentity(nano *Nanoleaf) *NanoIdentity {
	return &NanoIdentity{
		nano:     nano,
		endpoint: fmt.Sprintf("%s/%s/identify", nano.url, nano.token),
	}
}

// Flash let the light panels flash green twice
func (i *NanoIdentity) Flash() error {
	resp, err := i.nano.client.R().Put(i.endpoint)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return ErrUnexpectedResponse
	}

	return nil
}
