package model

type Language string

const (
	LanguageEnglish Language = "en"
	LanguageFrench  Language = "fr"
	LanguageSpanish Language = "es"
)

type OS string

const (
	OSMac     OS = "mac"
	OSAndroid OS = "android"
)

type Browser string

const (
	BrowserChrome Browser = "chrome"
	BrowserSafari Browser = "safari"
)

type Options struct {
	Language Language
	OS       OS
	Browser  Browser
}
