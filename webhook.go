package webhook

// Webhook is the webhook object sent to Discord in hook requests.
type Webhook struct {
	Username  string  `json:"username"`
	AvatarURL string  `json:"avatar_url"`
	Content   string  `json:"content"`
	Embeds    []Embed `json:"embeds"`
}

// Embed is the embed object in the webhook.
type Embed struct {
	Author      Author  `json:"author"`
	Footer      Footer  `json:"footer"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Timestamp   string  `json:"timestamp"`
	Thumbnail   Image   `json:"thumbnail"`
	Image       Image   `json:"image"`
	URL         string  `json:"url"`
	Fields      []Field `json:"fields"`
	Color       int64   `json:"color"`
}

// Author is the author object in the embed.
type Author struct {
	Name    string `json:"name"`
	URI     string `json:"uri"`
	IconURL string `json:"icon_url"`
}

// Field is a field in the embed.
type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// Footer is the footer object in the embed.
type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

// Image is an image that can be used in the embed.
type Image struct {
	URL string `json:"url"`
}
