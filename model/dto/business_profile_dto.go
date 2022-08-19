package dto

import (
	"time"
)

type BusinessProfileRequest struct {
	AccountID    uint   `json:"account_id"`
	CategoryID   uint   `json:"category_id"`
	Address      string `json:"address"`
	ProfileImage string `json:"profile_image"`
	ProfileBio   string `json:"profile_bio"`
	GmapsLink    string `json:"gmaps_link"`
	BusinessHours []OpeningHour `json:"business_hours"`
	BusinessLinks []ShoppingLink `json:"business_links"`
}

type OpeningHour struct {
	Day       int       `json:"day"`
	OpenHour  time.Time `json:"open_hour"`
	CloseHour time.Time `json:"close_hour"`
}

type ShoppingLink struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}
