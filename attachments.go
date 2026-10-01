package meet

import (
	"context"
	"errors"
	"net/http"

	"github.com/hayavo/hayavo-meet-go/internal/httpclient"
	"github.com/hayavo/hayavo-meet-go/models"
)

const maxAttachmentSize int64 = 1024 * 1024 * 1024

type AttachmentService struct {
	http *httpclient.Client
}

func NewAttachments(client *httpclient.Client) *AttachmentService {
	return &AttachmentService{
		http: client,
	}
}

func (a *AttachmentService) Upload(
	ctx context.Context,
	req *models.UploadAttachmentRequest,
) (*models.UploadAttachmentResponse, error) {

	if req == nil {
		return nil, errors.New("upload request is nil")
	}

	if req.Token == "" {
		return nil, errors.New("token is required")
	}

	if req.RoomID == "" {
		return nil, errors.New("room_id is required")
	}

	if req.File == nil {
		return nil, errors.New("file is required")
	}

	if req.FileName == "" {
		return nil, errors.New("file_name is required")
	}

	if req.MimeType == "" {
		req.MimeType = "application/octet-stream"
	}

	if req.FileSize > maxAttachmentSize {
		return nil, errors.New("maximum file size is 1024MB")
	}

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+req.Token)

	var resp models.UploadAttachmentResponse

	err := a.http.Upload(
		ctx,
		"/attachments/upload",
		req.File,
		req.FileName,
		req.MimeType,
		map[string]string{
			"room_id": req.RoomID,
		},
		headers,
		&resp,
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (a *AttachmentService) Download(
	ctx context.Context,
	attachmentID string,
	saveDir string,
	token string,
) (*models.DownloadAttachmentResponse, error) {

	headers := http.Header{}

	if token != "" {
		headers.Set("Authorization", "Bearer "+token)
	}

	path, filename, err := a.http.Download(
		ctx,
		"/attachments/"+attachmentID+"/download",
		saveDir,
		headers,
	)

	if err != nil {
		return nil, err
	}

	return &models.DownloadAttachmentResponse{
		Success:  true,
		Path:     path,
		Filename: filename,
	}, nil
}
