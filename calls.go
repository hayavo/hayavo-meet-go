package meet

import (
	"context"
	"errors"
	"fmt"

	"github.com/hayavo/hayavo-meet-go/config"
	"github.com/hayavo/hayavo-meet-go/internal/httpclient"
	"github.com/hayavo/hayavo-meet-go/models"
)

// CallService handles call-related APIs.
type CallService struct {
	http   *httpclient.Client
	config *config.Config
}

// NewCalls creates a new Call service.
func NewCalls(client *httpclient.Client, cfg *config.Config) *CallService {
	return &CallService{
		http:   client,
		config: cfg,
	}
}

// Start starts a new call.
func (c *CallService) Start(
	ctx context.Context,
	req *models.StartCallRequest,
) (*models.StartCallResponse, error) {

	audio := true
	video := true
	screen := false

	if req.Media.Audio != nil {
		audio = *req.Media.Audio
	}

	if req.Media.Video != nil {
		video = *req.Media.Video
	}

	if req.Media.Screen != nil {
		screen = *req.Media.Screen
	}

	if video && !c.config.AllowVideo {
		return nil, errors.New("video feature not allowed")
	}

	if audio && !c.config.AllowAudio {
		return nil, errors.New("audio feature not allowed")
	}

	if screen && !c.config.AllowScreen {
		return nil, errors.New("screen sharing not allowed")
	}

	req.Media.Audio = &audio
	req.Media.Video = &video
	req.Media.Screen = &screen

	var resp models.StartCallResponse

	if err := c.http.Post(
		ctx,
		"/rtc/call/start",
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HostToken generates a host token.
func (c *CallService) HostToken(
	ctx context.Context,
	req *models.CallTokenRequest,
) (*models.CallTokenResponse, error) {

	if req.CallID == "" {
		return nil, errors.New("call_id is required")
	}

	var resp models.CallTokenResponse

	if err := c.http.Post(
		ctx,
		"/rtc/call/token",
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	resp.Data.RTC = c.http.GetMediaPath()
	resp.Data.Signal = c.http.GetSignal()

	return &resp, nil
}

func (c *CallService) GuestToken(
	ctx context.Context,
	req *models.GuestCallTokenRequest,
) (*models.CallTokenResponse, error) {

	if req.AppID == "" {
		return nil, errors.New("app_id is required")
	}

	if req.CallID == "" {
		return nil, errors.New("call_id is required")
	}

	var resp models.CallTokenResponse

	if err := c.http.Post(
		ctx,
		"/create/call/guest/token",
		req,
		&resp,
		false,
	); err != nil {
		return nil, err
	}

	resp.Data.RTC = c.http.GetMediaPath()

	return &resp, nil
}

func (c *CallService) End(
	ctx context.Context,
	callID string,
) (*models.CommonResponse, error) {

	req := models.CallTokenRequest{
		CallID: callID,
	}

	var resp models.CommonResponse

	if err := c.http.Post(
		ctx,
		"/rtc/call/end",
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *CallService) MuteUser(
	ctx context.Context,
	callID string,
	identity string,
	mute bool,
) (*models.CommonResponse, error) {

	req := models.MuteUserRequest{
		Identity: identity,
		Mute:     mute,
	}

	var resp models.CommonResponse

	if err := c.http.Post(
		ctx,
		fmt.Sprintf("/rtc/call/%s/mute", callID),
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *CallService) VideoOffUser(
	ctx context.Context,
	callID string,
	identity string,
	disable bool,
) (*models.CommonResponse, error) {

	req := models.VideoOffRequest{
		Identity: identity,
		Disable:  disable,
	}

	var resp models.CommonResponse

	if err := c.http.Post(
		ctx,
		fmt.Sprintf("/rtc/call/%s/video-off", callID),
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *CallService) ScreenOffUser(
	ctx context.Context,
	callID string,
	identity string,
) (*models.CommonResponse, error) {

	req := models.IdentityRequest{
		Identity: identity,
	}

	var resp models.CommonResponse

	if err := c.http.Post(
		ctx,
		fmt.Sprintf("/client/calls/%s/screen-off", callID),
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *CallService) KickUser(
	ctx context.Context,
	callID string,
	identity string,
) (*models.CommonResponse, error) {

	req := models.IdentityRequest{
		Identity: identity,
	}

	var resp models.CommonResponse

	if err := c.http.Post(
		ctx,
		fmt.Sprintf("/rtc/call/%s/kick", callID),
		req,
		&resp,
		true,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
