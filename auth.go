package meet

import (
	"context"
	"errors"

	"github.com/hayavo/hayavo-meet-go/internal/httpclient"
	"github.com/hayavo/hayavo-meet-go/models"
)

type AuthService struct {
	http *httpclient.Client
}

func NewAuth(client *httpclient.Client) *AuthService {
	return &AuthService{
		http: client,
	}
}

func (a *AuthService) Authenticate(
	ctx context.Context,
	req *models.AuthRequest,
) (*models.AuthResponse, error) {

	if req.Login == "" {
		return nil, errors.New("username is required")
	}

	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	var resp models.AuthResponse

	err := a.http.Post(
		ctx,
		"/rtm/client/token",
		req,
		&resp,
		true,
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
