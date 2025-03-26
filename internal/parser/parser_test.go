package parser

import (
	"github.com/mespinosago/unitag/internal/model"
	"testing"
)

func TestParser_GetBrowser(t *testing.T) {
	tests := []struct {
		name   string
		header string
		want   model.Browser
	}{
		{
			name:   "get chrome",
			header: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36",
			want:   model.BrowserChrome,
		},
		{
			name:   "get safari",
			header: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.3.1 Safari/605.1.15",
			want:   model.BrowserSafari,
		},
		{
			name: "get default",
			want: model.BrowserChrome,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{}
			if got := p.GetBrowser(tt.header); got != tt.want {
				t.Errorf("GetBrowser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetLanguage(t *testing.T) {
	tests := []struct {
		name   string
		header string
		want   model.Language
	}{
		{
			name:   "get english",
			header: "en-GB,en-US;q=0.9,en;q=0.8,es;q=0.7",
			want:   model.LanguageEnglish,
		},
		{
			name:   "get french",
			header: "fr-FR,en-US;q=0.9,en;q=0.8,es;q=0.7",
			want:   model.LanguageFrench,
		},
		{
			name:   "get spanish",
			header: "es-SP,en-US;q=0.9,en;q=0.8,es;q=0.7",
			want:   model.LanguageSpanish,
		},
		{
			name:   "get default",
			header: "",
			want:   model.LanguageEnglish,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{}
			if got := p.GetLanguage(tt.header); got != tt.want {
				t.Errorf("GetLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetOS(t *testing.T) {
	tests := []struct {
		name   string
		header string
		want   model.OS
	}{
		{
			name:   "get android",
			header: "Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Mobile Safari/537.36 EdgA/123.0.0.0",
			want:   model.OSAndroid,
		},
		{
			name: "get mac for the rest of options",
			want: model.OSMac,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{}
			if got := p.GetOS(tt.header); got != tt.want {
				t.Errorf("GetOS() = %v, want %v", got, tt.want)
			}
		})
	}
}
