package meet

import (
	"context"

	"github.com/hayavo/hayavo-meet-go/internal/httpclient"
	"github.com/hayavo/hayavo-meet-go/models"
)

// AdminService handles meethayavo administrative APIs.
type AdminService struct {
	http *httpclient.Client
}

// NewAdmin creates a new Admin service.
func NewAdmin(client *httpclient.Client) *AdminService {
	return &AdminService{
		http: client,
	}
}

// Kick removes a participant from a room.
func (a *AdminService) Kick(
	ctx context.Context,
	roomName string,
	identity string,
) (*models.CommonResponse, error) {

	req := models.AdminUserRequest{
		RoomName: roomName,
		Identity: identity,
	}

	var resp models.CommonResponse

	if err := a.http.Post(
		ctx,
		"/rtc/av/kick",
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Mute mutes a participant in a room.
func (a *AdminService) Mute(
	ctx context.Context,
	roomName string,
	identity string,
) (*models.CommonResponse, error) {

	req := models.AdminUserRequest{
		RoomName: roomName,
		Identity: identity,
	}

	var resp models.CommonResponse

	if err := a.http.Post(
		ctx,
		"/rtc/video/mute",
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RoomCleaner ends the room and removes all participants.
func (a *AdminService) RoomCleaner(
	ctx context.Context,
	roomName string,
) (*models.CommonResponse, error) {

	req := models.RoomCleanerRequest{
		RoomName: roomName,
	}

	var resp models.CommonResponse

	if err := a.http.Post(
		ctx,
		"/rtc/video/end",
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
