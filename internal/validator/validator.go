package validator

import (
	"errors"

	"github.com/hayavo/hayavo-meet-go/config"
)

type Validator struct {
	config *config.Config
}

func New(cfg *config.Config) *Validator {
	return &Validator{
		config: cfg,
	}
}

func (v *Validator) Validate() error {

	if v.config.APIKey == "" {
		return errors.New("API key is required")
	}

	if v.config.AppID == "" {
		return errors.New("App ID is required")
	}

	if v.config.AppSecret == "" {
		return errors.New("App secret is required")
	}

	if v.config.Platform == "" {
		return errors.New("App platform is required")
	}

	if v.config.Identifier == "" {
		return errors.New("App identifier is required")
	}

	return nil
}

func (v *Validator) BuildHeaders() (map[string]string, error) {

	if err := v.Validate(); err != nil {
		return nil, err
	}

	headers := map[string]string{
		"X-API-Key":        v.config.APIKey,
		"X-App-ID":         v.config.AppID,
		"X-App-Secret":     v.config.AppSecret,
		"X-App-Platform":   v.config.Platform,
		"X-App-Identifier": v.config.Identifier,
	}

	if v.config.AppName != "" {
		headers["X-App-Name"] = v.config.AppName
	}

	return headers, nil
}
