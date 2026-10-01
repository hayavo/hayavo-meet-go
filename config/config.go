package config

import "time"

type Config struct {
	APIKey    string
	AppID     string
	AppSecret string

	Platform   string
	Identifier string

	AppName string
	Domain  string

	BaseURL string

	Timeout time.Duration

	Debug bool

	AllowVideo  bool
	AllowAudio  bool
	AllowScreen bool
}
