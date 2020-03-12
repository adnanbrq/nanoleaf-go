package nanoleaf

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NanoLayout layouts
type NanoLayout struct {
	nano     *Nanoleaf
	endpoint string
}

// GlobalOrientation globalOrientation as a go struct
type GlobalOrientation MinMaxValue

// PanelPositionData positionData as a go struct
type PanelPositionData struct {
	ID int `json:"panelId"`
	X  int `json:"x"`
	Y  int `json:"y"`
	Z  int `json:"z"`
}

// PanelLayout panelLayout as a go struct
type PanelLayout struct {
	Panels       int                 `json:"numPanels"`
	SideLength   int                 `json:"sideLength"`
	PositionData []PanelPositionData `json:"positionData"`
}

// newNanoLayout returns a new instance of NanoLayout
func newNanoLayout(nano *Nanoleaf) *NanoLayout {
	return &NanoLayout{
		nano:     nano,
		endpoint: fmt.Sprintf("%s/%s/panelLayout", nano.url, nano.token),
	}
}

// GetGlobalOrientation returns the global orientation
func (l *NanoLayout) GetGlobalOrientation() (*GlobalOrientation, error) {
	url := fmt.Sprintf("%s/globalOrientation", l.endpoint)
	resp, err := l.nano.client.R().Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return nil, ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, ErrUnexpectedResponse
	}

	var globalOrientation GlobalOrientation

	if err := json.Unmarshal(resp.Body(), &globalOrientation); err != nil {
		return nil, ErrParsingJSON
	}

	return &globalOrientation, nil
}

// SetGlobalOrientation sets the global orientation
func (l *NanoLayout) SetGlobalOrientation(value int) error {
	url := fmt.Sprintf("%s/globalOrientation", l.endpoint)
	body := jsonPayload{"globalOrientation": jsonPayload{"value": value}}
	resp, err := l.nano.client.R().SetHeader("Content-Type", "application/josn").SetBody(body).Put(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusNoContent {
		return ErrUnexpectedResponse
	}

	return nil
}

// GetLayout returns the layout of nanoleafs
func (l *NanoLayout) GetLayout() (*PanelLayout, error) {
	url := fmt.Sprintf("%s/layout", l.endpoint)
	resp, err := l.nano.client.R().Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return nil, ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, ErrUnexpectedResponse
	}

	var panelLayout PanelLayout

	if err := json.Unmarshal(resp.Body(), &panelLayout); err != nil {
		return nil, ErrParsingJSON
	}

	return &panelLayout, nil
}
