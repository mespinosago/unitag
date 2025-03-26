package service

import (
	"github.com/mespinosago/unitag/internal/model"
	"testing"
)

func TestService_GetURL(t *testing.T) {
	tests := []struct {
		name        string
		code        string
		options     model.Options
		expectedURL string
		expectedErr error
	}{

		{
			name:        "get invalid code error",
			expectedErr: ErrCodeNotFound,
		},
		{
			name: "get url from browser / language",
			code: "r7TH8k",
			options: model.Options{
				Language: model.LanguageFrench,
				Browser:  model.BrowserChrome,
			},
			expectedURL: "https://www.google.fr",
		},
		{
			name: "default url if language does not exist for browser",
			code: "r7TH8k",
			options: model.Options{
				Browser: model.BrowserChrome,
			},
			expectedURL: "https://www.unitag.io",
		},
		{
			name: "get url from OS / language",
			code: "r7TH8k",
			options: model.Options{
				Language: model.LanguageSpanish,
				OS:       model.OSMac,
			},
			expectedURL: "https://www.apple.com/es",
		},
		{
			name: "default url if language does not exist for OS",
			code: "r7TH8k",
			options: model.Options{
				OS: model.OSAndroid,
			},
			expectedURL: "https://www.unitag.io",
		},
		{
			name:        "default url for empty options",
			code:        "r7TH8k",
			expectedURL: "https://www.unitag.io",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			got, err := s.GetURL(tt.code, tt.options)
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("GetURL() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if got != tt.expectedURL {
				t.Errorf("GetURL() got = %v, expectedURL %v", got, tt.expectedURL)
			}
		})
	}
}
