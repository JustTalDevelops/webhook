package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
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

// Send sends the webhook provided to the hook.
func (h *Hook) Send(webhook Webhook) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(webhook)
	if err != nil {
		return fmt.Errorf("error encoding the webhook: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, h.uri, nil)
	if err != nil {
		return fmt.Errorf("error while building a webhook request: %w", err)
	}

	if len(webhook.Files) < 1 {
		req.Header.Add("Content-Type", "application/json")
		req.Body = io.NopCloser(&buf)
	} else {
		var body bytes.Buffer
		mw := multipart.NewWriter(&body)
		_ = mw.WriteField("payload_json", buf.String())

		for i, f := range webhook.Files {
			w, err := mw.CreateFormFile(fmt.Sprintf("file%d", i), f.Filename)
			if err != nil {
				return err
			}

			if _, err = io.Copy(w, f.Body); err != nil {
				return err
			}
		}

		_ = mw.Close()

		req.Header.Set("Content-Type", mw.FormDataContentType())
		req.Body = io.NopCloser(&body)
	}

	resp, err := h.Client().Do(req)
	if err != nil {
		return fmt.Errorf("webhook: error when sending: %w", err)
	}
	return resp.Body.Close()
}
