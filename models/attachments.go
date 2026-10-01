package models

import (
	"io"
)

type UploadAttachmentResponse struct {
	AttachmentID string `json:"attachment_id"`

	FileName string `json:"file_name"`

	FileSize int64 `json:"file_size"`

	MimeType string `json:"mime_type"`
}

type DownloadAttachmentResponse struct {
	Success  bool   `json:"success"`
	Path     string `json:"path"`
	Filename string `json:"filename"`
}

type UploadAttachmentRequest struct {
	Token    string    `json:"token"`
	RoomID   string    `json:"room_id"`
	FileName string    `json:"file_name"`
	File     io.Reader `json:"file"`
	FileSize int64     `json:"file_size"`
	MimeType string    `json:"mime_type"`
}
