package models

type Url struct {
	UrlProfile string `json:"spotify"`
}

type Image struct {
	ImageUrl string `json:"url"`
}

type Profile struct {
	Name   string  `json:"display_name"`
	Urls   Url     `json:"external_urls"`
	Images []Image `json:"images"`
	Email  string  `json:"email"`
}
