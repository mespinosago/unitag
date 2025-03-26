package service

import (
	"errors"
	"github.com/mespinosago/unitag/internal/model"
)

var (
	// default urls by code
	defaultURLs = map[string]string{
		"r7TH8k": "https://www.unitag.io",
	}
	urlGoogle = map[model.Language]string{
		model.LanguageEnglish: "https://www.google.com",
		model.LanguageFrench:  "https://www.google.fr",
		model.LanguageSpanish: "https://www.google.es",
	}
	urlApple = map[model.Language]string{
		model.LanguageEnglish: "https://www.apple.com",
		model.LanguageFrench:  "https://www.apple.com/fr",
		model.LanguageSpanish: "https://www.apple.com/es",
	}
	urlsByBrowser = map[model.Browser]map[model.Language]string{
		model.BrowserChrome: urlGoogle,
		model.BrowserSafari: urlApple,
	}
	urlsByOS = map[model.OS]map[model.Language]string{
		model.OSAndroid: urlGoogle,
		model.OSMac:     urlApple,
	}
)

var ErrCodeNotFound = errors.New("not found")

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetURL(code string, options model.Options) (string, error) {
	var (
		urlsByLang map[model.Language]string
		url        string
	)
	defaultURL, ok := defaultURLs[code]
	if !ok {
		return "", ErrCodeNotFound
	}
	// first try getting the url by browser
	if urlsByLang, ok = urlsByBrowser[options.Browser]; ok {
		if url, ok = urlsByLang[options.Language]; ok {
			return url, nil
		}
	}
	// then try getting the url by OS
	if urlsByLang, ok = urlsByOS[options.OS]; ok {
		if url, ok = urlsByLang[options.Language]; ok {
			return url, nil
		}
	}
	return defaultURL, nil
}
