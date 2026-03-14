package model

import (
	"net/url"
	"strings"
)

// ShortenRequest represents the request to shorten a URL
type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

// ShortenResponse represents the response after shortening a URL
type ShortenResponse struct {
	ShortCode string `json:"short_code"`
	ShortURL  string `json:"short_url"`
	Original  string `json:"original_url"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// Validate checks if the URL is valid
func (r *ShortenRequest) Validate() error {
	// Trim whitespace
	r.URL = strings.TrimSpace(r.URL)

	// Check if empty
	if r.URL == "" {
		return ErrEmptyURL
	}

	// Parse URL
	parsedURL, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return ErrInvalidURL
	}

	// Check if scheme is http or https
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return ErrInvalidScheme
	}

	// Check if host exists
	if parsedURL.Host == "" {
		return ErrInvalidHost
	}

	return nil
}

// Custom errors
type ValidationError struct {
	message string
}

func (e *ValidationError) Error() string {
	return e.message
}

var (
	ErrEmptyURL      = &ValidationError{"URL cannot be empty"}
	ErrInvalidURL    = &ValidationError{"Invalid URL format"}
	ErrInvalidScheme = &ValidationError{"URL must use http or https scheme"}
	ErrInvalidHost   = &ValidationError{"URL must have a valid host"}
)
