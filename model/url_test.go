package model

import (
	"testing"
)

func TestShortenRequestValidate(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
		errType error
	}{
		{
			name:    "valid http url",
			url:     "http://example.com",
			wantErr: false,
		},
		{
			name:    "valid https url",
			url:     "https://example.com/path?query=value",
			wantErr: false,
		},
		{
			name:    "empty url",
			url:     "",
			wantErr: true,
			errType: ErrEmptyURL,
		},
		{
			name:    "url with whitespace",
			url:     "  https://example.com  ",
			wantErr: false,
		},
		{
			name:    "invalid scheme",
			url:     "ftp://example.com",
			wantErr: true,
			errType: ErrInvalidScheme,
		},
		{
			name:    "no scheme",
			url:     "example.com",
			wantErr: true,
			errType: ErrInvalidURL,
		},
		{
			name:    "invalid format",
			url:     "not-a-url",
			wantErr: true,
			errType: ErrInvalidURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &ShortenRequest{URL: tt.url}
			err := req.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.errType != nil && err != tt.errType {
				t.Errorf("Validate() error = %v, want %v", err, tt.errType)
			}
		})
	}
}
