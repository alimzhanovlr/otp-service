package sms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
type Client struct {
	baseURL string
	logger  *slog.Logger
	client  HTTPDoer
}

func (c *Client) Send(ctx context.Context, in Payload) error {
	b, _ := json.Marshal(in)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.baseURL+"/sms", bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var r struct {
			Message string `json:"message"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&r)
		return fmt.Errorf("failed to send sms: %s", r.Message)
	}
	return nil
}
