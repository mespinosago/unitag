package parser

import (
	"github.com/mespinosago/unitag/internal/model"
	"strings"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

// GetLanguage just takes the first language listed in the header (for simplicity)
func (p *Parser) GetLanguage(header string) model.Language {
	languages := strings.Split(header, ",")
	for _, language := range languages {
		if language == "" {
			continue
		}
		l := strings.Split(language, ";")
		if len(l) == 0 {
			continue
		}
		return model.Language(l[0][:2]) // returns just the generic alpha-2 language code
	}
	// Default option
	return model.LanguageEnglish
}

// GetOS returns the OS from the header, for simplicity, just Android and "Mac"
func (p *Parser) GetOS(header string) model.OS {
	if strings.Contains(header, "Android") {
		return model.OSAndroid
	}
	return model.OSMac
}

// GetBrowser returns the browser from the header, for simplicity, just Chrome and Safari
func (p *Parser) GetBrowser(header string) model.Browser {
	switch {
	// Be aware the Chrome requests also contain `Safari`
	case strings.Contains(header, "Chrome"):
		return model.BrowserChrome
	// If Chrome is not contained but Safari it is then, Safari
	case strings.Contains(header, "Safari"):
		return model.BrowserSafari
	}
	return model.BrowserChrome
}
