package models

type RoomTokenData struct {
	RTC    string `json:"rtc"`
	Signal string `json:"signal"`
}

type CreateRoomRequest struct {
	RoomName        string `json:"room_name"`
	Mode            string `json:"mode,omitempty"`
	MaxParticipants int    `json:"max_participants,omitempty"`
}

type CreateRoomResponse struct {
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
	Mode     string `json:"mode"`
	Status   string `json:"status"`
}

type HostTokenRequest struct {
	RoomName    string            `json:"room_name" binding:"required"`
	DisplayName string            `json:"display_name,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

type GuestTokenRequest struct {
	RoomName    string            `json:"room_name"`
	AppID       string            `json:"app_id"`
	DisplayName string            `json:"display_name,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

type RoomTokenResponse struct {
	Token    string        `json:"token"`
	Identity string        `json:"identity"`
	Role     string        `json:"role"`
	Data     RoomTokenData `json:"data"`
}
