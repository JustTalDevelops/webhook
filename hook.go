package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Hook is the toplevel object. It contains everything required to send webhooks. It's optimised for extendability, so
// it has an embedded http.Client. You can modify the client yourself if you want to change the defaults.
type Hook struct {
	client *http.Client
	uri    string
}

// New returns a new hook with the designated URL.
func New(uri string) *Hook {
	return &Hook{
		uri:    uri,
		client: http.DefaultClient,
	}
}

// Client returns the http.Client used by the hook.
func (h *Hook) Client() *http.Client {
	return h.client
}

// URI returns the URI of the hook.
func (h *Hook) URI() string {
	return h.uri
}

// Send sends the webhook provided to the hook.
func (h *Hook) Send(webhook Webhook) error {
	enc, err := json.Marshal(webhook)
	if err != nil {
		return fmt.Errorf("hook send: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, h.uri, bytes.NewReader(enc))
	if err != nil {
		return fmt.Errorf("hook send: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := h.Client().Do(req)
	if err != nil {
		return fmt.Errorf("hook send: %w", err)
	}
	return resp.Body.Close()
}
