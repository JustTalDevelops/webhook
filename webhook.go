package webhook

// Webhook is the webhook object sent to Discord in hook requests.
type Webhook struct {
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	Content   string  `json:"content,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

// Embed is the embed object in the webhook.
type Embed struct {
	Author      Author  `json:"author,omitempty"`
	Footer      Footer  `json:"footer,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
	Thumbnail   Image   `json:"thumbnail,omitempty"`
	Image       Image   `json:"image,omitempty"`
	URL         string  `json:"url,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
	Color       int64   `json:"color,omitempty"`
}

// Author is the author object in the embed.
type Author struct {
	Name    string `json:"name,omitempty"`
	URI     string `json:"uri,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// Field is a field in the embed.
type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// Footer is the footer object in the embed.
type Footer struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// Image is an image that can be used in the embed.
type Image struct {
	URL string `json:"url"`
}
