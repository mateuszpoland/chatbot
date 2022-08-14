package structs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"os"
)

//input msg from messenger
type InputMessage struct {
	Entry []struct {
		ID string `json:"id"`
		Time int64 `json:"time"`
		Messaging []struct {
			Sender struct {
				ID string `json:"id"`
			} `json:"sender"`
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Timestamp int64 `json:"timestamp"`
			Message struct {
				Mid string `json:"mid"`
				Text string `json:"text"`
			} `json:"message"`
		} `json:"messaging"`
	}
}

// message recipient
type Recipient struct {
	ID string `json:"id"`
}

type OutputMessage struct {
	Text string `json:"text,omitempty"`
}

type Button struct {
	Type string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
	URL string `json:"url,omitempty"`
}

type Element struct {
	Title string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	DefaultAction DefaultAction `json:"default_action,omitempty"`
	Title string `json:"buttons,omitempty"`
}

type DefaultAction struct {
	Type string `json:"type,omitempty"`
	URL string `json:"url,omitempty"`
	WebViewHeightRattio string `json:"web_view_height_ratio,omitempty"`
}

//attachment to send
type Attachment struct {
	Attachment struct {
	 Type    string `json:"type,omitempty"`
	 Payload struct {
	  TemplateType string    `json:"template_type,omitempty"`
	  Elements     []Element `json:"elements,omitempty"`
	 } `json:"payload,omitempty"`
	} `json:"attachment,omitempty"`
}

// Full response
type ResponseAttachment struct {
	Recipient Recipient  `json:"recipient"`
	Message   Attachment `json:"message,omitempty"`
}

// Full response
type ResponseMessage struct {
	Recipient Recipient `json:"recipient"`
	Message OutputMessage   `json:"message,omitempty"`
}