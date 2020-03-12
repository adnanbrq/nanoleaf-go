package nanoleaf

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// Nanoleaf nanoleaf object
type Nanoleaf struct {
	client   *resty.Client
	url      string
	token    string
	Identity *NanoIdentity
	Auth     *NanoAuth
	Effects  *NanoEffects
	State    *NanoState
	Stream   *NanoStream
	Layout   *NanoLayout
}

// jsonPayload is used to append a json body in requests targeting the nanoleaf api
type jsonPayload map[string]interface{}

// ControllerInfo as a go struct
type ControllerInfo struct {
	Name            string `json:"name"`
	Serial          string `json:"serialNo"`
	Manufacturer    string `json:"manufacturer"`
	FirmwareVersion string `json:"firmwareVersion"`
	Model           string `json:"model"`
	State           struct {
		On struct {
			Value bool `json:"value"`
		} `json:"on"`
		Brightness struct {
			Value int `json:"value"`
			Max   int `json:"max"`
			Min   int `json:"min"`
		} `json:"brightness"`
		Hue struct {
			Value int `json:"value"`
			Max   int `json:"max"`
			Min   int `json:"min"`
		} `json:"hue"`
		Sat struct {
			Value int `json:"value"`
			Max   int `json:"max"`
			Min   int `json:"min"`
		} `json:"sat"`
		Ct struct {
			Value int `json:"value"`
			Max   int `json:"max"`
			Min   int `json:"min"`
		} `json:"ct"`
		ColorMode string `json:"colorMode"`
	} `json:"state"`
	Effects struct {
		Active string   `json:"select"`
		List   []string `json:"effectsList"`
	} `json:"effects"`
	PanelLayout struct {
		Layout            PanelLayout `json:"layout"`
		GlobalOrientation `json:"globalOrientation"`
	} `json:"panelLayout"`
	Rythm struct {
		Connected       bool   `json:"rythmConnected"`
		Active          bool   `json:"rythmActive"`
		ID              bool   `json:"rythmId"`
		HardwareVersion string `json:"hardwareVersion"`
		FirmwareVersion string `json:"firmwareVersion"`
		AuxAvailable    bool   `json:"auxAvailable"`
		Mode            string `json:"rythmMode"`
		Pos             string `json:"rythmPos"`
	} `json:"rythm"`
}

// NewNanoleaf created a
func NewNanoleaf(url string) *Nanoleaf {
	n := &Nanoleaf{
		client: resty.New(),
		url:    url,
	}

	n.Auth = newNanoAuth(n)
	n.Stream = newNanoStream(n)

	return n
}

// SetToken used internally to override token
func (n *Nanoleaf) SetToken(token string) {
	n.token = token

	n.Identity = newNanoIdentity(n)
	n.Effects = newNanoEffects(n)
	n.State = newNanoState(n)
	n.Layout = newNanoLayout(n)
}

// GetToken returns the current token
func (n *Nanoleaf) GetToken() string {
	return n.token
}

// IsConnected checks if we have a connection so far
func (n *Nanoleaf) IsConnected() bool {
	return n.token != ""
}

// GetControllerInfo returns controllerInfo
func (n *Nanoleaf) GetControllerInfo() (*ControllerInfo, error) {
	url := fmt.Sprintf("%s/%s", n.url, n.token)
	resp, err := n.client.R().Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return nil, ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, ErrUnexpectedResponse
	}

	var controllerInfo ControllerInfo
	if err := json.Unmarshal(resp.Body(), &controllerInfo); err != nil {
		fmt.Println(err)
		return nil, ErrParsingJSON
	}

	return &controllerInfo, nil
}
