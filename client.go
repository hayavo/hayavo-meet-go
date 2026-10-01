package meet

import (
	"github.com/hayavo/hayavo-meet-go/config"
	"github.com/hayavo/hayavo-meet-go/internal/httpclient"
)

type Client struct {
	Auth        *AuthService
	Rooms       *RoomService
	Calls       *CallService
	Admin       *AdminService
	Attachments *AttachmentService
}

func New(cfg *config.Config) *Client {

	http := httpclient.New(cfg)

	return &Client{
		Auth:        NewAuth(http),
		Rooms:       NewRooms(http),
		Calls:       NewCalls(http, cfg),
		Admin:       NewAdmin(http),
		Attachments: NewAttachments(http),
	}
}
