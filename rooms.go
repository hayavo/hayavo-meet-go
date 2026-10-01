package meet

import (
	"context"
	"errors"

	"github.com/hayavo/hayavo-meet-go/internal/httpclient"
	"github.com/hayavo/hayavo-meet-go/models"
)

// RoomService handles room-related APIs.
type RoomService struct {
	http *httpclient.Client
}

// NewRooms creates a new Room service.
func NewRooms(client *httpclient.Client) *RoomService {
	return &RoomService{
		http: client,
	}
}

// Create creates a new room.
func (r *RoomService) Create(
	ctx context.Context,
	req *models.CreateRoomRequest,
) (*models.CreateRoomResponse, error) {

	if req.RoomName == "" {
		return nil, errors.New("room_name is required")
	}

	if req.Mode == "" {
		req.Mode = "hm_default_meet_room"
	}

	if req.MaxParticipants == 0 {
		req.MaxParticipants = 5
	}

	var resp models.CreateRoomResponse

	if err := r.http.Post(
		ctx,
		"/rtc/room/create",
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *RoomService) HostToken(
	ctx context.Context,
	req *models.HostTokenRequest,
) (*models.RoomTokenResponse, error) {

	if req.RoomName == "" {
		return nil, errors.New("room_name is required")
	}

	var resp models.RoomTokenResponse

	if err := r.http.Post(
		ctx,
		"/rtc/room/token/host",
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	resp.Data.RTC = r.http.GetMediaPath()
	resp.Data.Signal = r.http.GetSignal()

	return &resp, nil
}

func (r *RoomService) GuestToken(
	ctx context.Context,
	req *models.GuestTokenRequest,
) (*models.RoomTokenResponse, error) {

	if req.RoomName == "" {
		return nil, errors.New("room_name is required")
	}

	if req.AppID == "" {
		return nil, errors.New("app_id is required")
	}

	var resp models.RoomTokenResponse

	if err := r.http.Post(
		ctx,
		"/create/room/token/guest",
		req,
		&resp,
		false,
	); err != nil {
		return nil, err
	}

	resp.Data.RTC = r.http.GetMediaPath()

	return &resp, nil
}
