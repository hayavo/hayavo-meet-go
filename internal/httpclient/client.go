package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/hayavo/hayavo-meet-go/config"
	"github.com/hayavo/hayavo-meet-go/internal/validator"
)

type Client struct {
	baseURL  string
	client   *http.Client
	config   *config.Config
	validate *validator.Validator
}

func New(cfg *config.Config) *Client {
	return &Client{
		baseURL:  "https://api.meet.hayavo.com/v1",
		config:   cfg,
		validate: validator.New(cfg),
		client: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

func (c *Client) Post(
	ctx context.Context,
	endpoint string,
	request any,
	response any,
	authRequired bool,
) error {

	headers, err := c.buildHeaders(authRequired)
	if err != nil {
		return err
	}

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+endpoint,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	req.Header = headers

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {

		data, _ := io.ReadAll(resp.Body)

		return errors.New(string(data))
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

func (c *Client) buildHeaders(authRequired bool) (http.Header, error) {

	headers := make(http.Header)

	headers.Set("User-Agent", "hayavo-meet-go-sdk/1.0.0")
	headers.Set("X-SDK-Version", "1.0.0")
	headers.Set("Content-Type", "application/json")

	if authRequired {

		authHeaders, err := c.validate.BuildHeaders()
		if err != nil {
			return nil, err
		}

		for key, value := range authHeaders {
			headers.Set(key, value)
		}
	}

	return headers, nil
}

func (c *Client) Upload(
	ctx context.Context,
	endpoint string,
	file io.Reader,
	filename string,
	mimeType string,
	fields map[string]string,
	headers http.Header,
	response any,
) error {

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	for key, value := range fields {
		if err := writer.WriteField(key, value); err != nil {
			return err
		}
	}

	// Add file
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}

	if _, err := io.Copy(part, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+endpoint,
		&body,
	)

	if err != nil {
		return err
	}

	// SDK headers
	req.Header.Set("User-Agent", "hayavo-meet-go/1.0.0")
	req.Header.Set("X-SDK-Version", "1.0.0")

	// multipart content type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// custom headers (Authorization, etc.)
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {

		data, _ := io.ReadAll(resp.Body)

		return errors.New(string(data))
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

func (c *Client) Download(
	ctx context.Context,
	endpoint string,
	saveDir string,
	headers http.Header,
) (
	string,
	string,
	error,
) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.baseURL+endpoint,
		nil,
	)

	if err != nil {
		return "", "", err
	}

	req.Header.Set("User-Agent", "hayavo-meet-go/1.0.0")
	req.Header.Set("X-SDK-Version", "1.0.0")

	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {

		data, _ := io.ReadAll(resp.Body)

		return "", "", errors.New(string(data))
	}

	finalURL := resp.Request.URL.String()

	u, err := url.Parse(finalURL)
	if err != nil {
		return "", "", err
	}

	filename := filepath.Base(u.Path)

	if filename == "" || filename == "." {
		filename = "download"
	}

	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return "", "", err
	}

	fullPath := filepath.Join(saveDir, filename)

	out, err := os.Create(fullPath)
	if err != nil {
		return "", "", err
	}

	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return "", "", err
	}

	return fullPath, filename, nil
}

func (c *Client) GetMediaPath() string {
	return fmt.Sprintf("%s://%s.%s.%s", "wss", "edge", "meet", "hayavo.com")

}

func (c *Client) GetSignal() string {
	return fmt.Sprintf("%s://%s-%s-%s-%s-%s.%s/%s", "wss", "signal", "core", "node", "83m2q", "meet", "hayavo.com", "ws")

}
