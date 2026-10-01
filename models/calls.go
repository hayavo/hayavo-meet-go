package models

type Media struct {
	Audio  *bool `json:"audio,omitempty"`
	Video  *bool `json:"video,omitempty"`
	Screen *bool `json:"screen,omitempty"`
}

type StartCallRequest struct {
	CalleeID string `json:"callee_id"`
	Media    Media  `json:"media"`
}

type StartCallResponse struct {
	CallID string `json:"call_id"`
	Room   string `json:"room"`
}

type CallTokenRequest struct {
	CallID      string            `json:"call_id"`
	DisplayName string            `json:"display_name,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

type GuestCallTokenRequest struct {
	CallID      string            `json:"call_id"`
	AppID       string            `json:"app_id"`
	DisplayName string            `json:"display_name,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

type CallTokenData struct {
	RTC    string `json:"rtc"`
	Signal string `json:"signal"`
}

type CallTokenResponse struct {
	Token string        `json:"token"`
	Data  CallTokenData `json:"data"`
}

type EndCallRequest struct {
	CallID string `json:"call_id"`
}

type MuteUserRequest struct {
	Identity string `json:"identity"`
	Mute     bool   `json:"mute"`
}

type VideoOffRequest struct {
	Identity string `json:"identity"`
	Disable  bool   `json:"disable"`
}

type IdentityRequest struct {
	Identity string `json:"identity"`
}

type CommonResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
