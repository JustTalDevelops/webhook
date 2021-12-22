# webhook

webhook is a bare-bones package for working with Discord webhooks. This version is a fork of the original to change
a few things to make it more useful for my own use. This implementation does not support attachments.

## Example

Below is a simple example to send plain text to a webhook:
```go
webhook.New(hookURI).Send(webhook.Webhook{
    Username:  "Captain'Hook",
    Content:   "This is the content of the message, it's plain text",
})
```
