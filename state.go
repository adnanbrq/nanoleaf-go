package nanoleaf

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NanoState nanoleaf state
type NanoState struct {
	nano     *Nanoleaf
	endpoint string
}

// MinMaxValue represents a common response type from nanoleaf api
type MinMaxValue struct {
	Value int `json:"value"`
	Max   int `json:"max"`
	Min   int `json:"min"`
}

// Brightness nanoleaf Brightness
type Brightness MinMaxValue

// Hue nanoleaf Hue
type Hue MinMaxValue

// Saturation nanoleaf Saturation
type Saturation MinMaxValue

// ColorTemprature nanoleaf ColorTemprature
type ColorTemprature MinMaxValue

// newNanoState returns a new instance of State
func newNanoState(nano *Nanoleaf) *NanoState {
	return &NanoState{
		nano:     nano,
		endpoint: fmt.Sprintf("%s/%s/state", nano.url, nano.token),
	}
}

// IsOn checks if the nanoleafs are currently on
func (s *NanoState) IsOn() (bool, error) {
	url := fmt.Sprintf("%s/on", s.endpoint)
	resp, err := s.nano.client.R().Get(url)

	if err != nil {
		return false, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return false, ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return false, ErrUnexpectedResponse
	}

	return true, nil
}

// SetOn "on" means nanoleafs do light up and "off" for the opposite
func (s *NanoState) SetOn(state bool) error {
	body := jsonPayload{"on": jsonPayload{"value": state}}
	resp, err := s.nano.client.R().SetHeader("Content-Type", "application/json").SetBody(body).Put(s.endpoint)

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

// GetBrightness returns the current brightness
func (s *NanoState) GetBrightness() (Brightness, error) {
	var brightness Brightness
	url := fmt.Sprintf("%s/brightness", s.endpoint)
	resp, err := s.nano.client.R().Get(url)

	if err != nil {
		return brightness, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return brightness, ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return brightness, ErrUnexpectedResponse
	}

	if err := json.Unmarshal(resp.Body(), &brightness); err != nil {
		return brightness, ErrParsingJSON
	}

	return brightness, nil
}

// SetBrightness sets the Light Panels Brightness over given time (ms)
func (s *NanoState) SetBrightness(value, time int) error {
	body := jsonPayload{
		"brightness": jsonPayload{
			"value":    value,
			"duration": time,
		},
	}

	resp, err := s.nano.client.R().SetHeader("Content-Type", "application/json").SetBody(body).Put(s.endpoint)

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

// GetHue returns the current brightness
func (s *NanoState) GetHue() (Hue, error) {
	var hue Hue
	url := fmt.Sprintf("%s/hue", s.endpoint)
	resp, err := s.nano.client.R().Get(url)

	if err != nil {
		return hue, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return hue, ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return hue, ErrUnexpectedResponse
	}

	if err := json.Unmarshal(resp.Body(), &hue); err != nil {
		return hue, ErrParsingJSON
	}

	return hue, nil
}

// SetHue sets the hue or increments it
func (s *NanoState) SetHue(value int, isIncremental bool) error {
	var body jsonPayload

	if isIncremental {
		body = jsonPayload{"hue": jsonPayload{"increment": value}}
	} else {
		body = jsonPayload{"hue": jsonPayload{"value": value}}
	}

	resp, err := s.nano.client.R().SetHeader("Content-Type", "application/json").SetBody(body).Put(s.endpoint)

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

// GetSaturation returns the current brightness
func (s *NanoState) GetSaturation() (Saturation, error) {
	var saturation Saturation
	url := fmt.Sprintf("%s/sat", s.endpoint)
	resp, err := s.nano.client.R().Get(url)

	if err != nil {
		return saturation, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return saturation, ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return saturation, ErrUnexpectedResponse
	}

	if err := json.Unmarshal(resp.Body(), &saturation); err != nil {
		return saturation, ErrParsingJSON
	}

	return saturation, nil
}

// SetSaturation sets the saturation or increments it
func (s *NanoState) SetSaturation(value int, isIncremental bool) error {
	var body jsonPayload

	if isIncremental {
		body = jsonPayload{"sat": jsonPayload{"increment": value}}
	} else {
		body = jsonPayload{"sat": jsonPayload{"value": value}}
	}

	resp, err := s.nano.client.R().SetHeader("Content-Type", "application/json").SetBody(body).Put(s.endpoint)

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

// GetColorTemp returns the current color temprature
func (s *NanoState) GetColorTemp() (ColorTemprature, error) {
	var colorTemp ColorTemprature
	url := fmt.Sprintf("%s/ct", s.endpoint)
	resp, err := s.nano.client.R().Get(url)

	if err != nil {
		return colorTemp, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return colorTemp, ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return colorTemp, ErrUnexpectedResponse
	}

	if err := json.Unmarshal(resp.Body(), &colorTemp); err != nil {
		return colorTemp, ErrParsingJSON
	}

	return colorTemp, nil
}

// SetColorTemp sets the color temprature or increments it
func (s *NanoState) SetColorTemp(value int, isIncremental bool) error {
	var body jsonPayload

	if isIncremental {
		body = jsonPayload{"ct": jsonPayload{"increment": value}}
	} else {
		body = jsonPayload{"ct": jsonPayload{"value": value}}
	}

	resp, err := s.nano.client.R().SetHeader("Content-Type", "application/json").SetBody(body).Put(s.endpoint)

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

// GetColorMode returns the current color temprature
func (s *NanoState) GetColorMode() (string, error) {
	colorMode := ""
	url := fmt.Sprintf("%s/colorMode", s.endpoint)
	resp, err := s.nano.client.R().Get(url)

	if err != nil {
		return colorMode, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		return colorMode, ErrUnauthorized
	}

	if resp.StatusCode() != http.StatusOK {
		return colorMode, ErrUnexpectedResponse
	}

	colorMode = resp.String()
	return colorMode, nil
}
