package models

type AdminUserRequest struct {
	RoomName string `json:"room_name"`
	Identity string `json:"identity"`
}

type RoomCleanerRequest struct {
	RoomName string `json:"room_name"`
}
