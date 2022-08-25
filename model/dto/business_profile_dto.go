package dto

import "itdp-group3-backend/model/entity"

type BusinessProfileRequest struct {
	AccountID     string         `json:"account_id"`
	CategoryID    string         `json:"category_id"`
	Address       string         `json:"address"`
	ProfileImage  string         `json:"profile_image"`
	ProfileBio    string         `json:"profile_bio"`
	GmapsLink     string         `json:"gmaps_link"`
	DisplayName   string         `json:"display_name"`
	BusinessHours []OpeningHour  `json:"business_hours"`
	BusinessLinks []ShoppingLink `json:"business_links"`
}

type OpeningHour struct {
	Day       string `json:"day"`
	OpenHour  string `json:"open_hour"`
	CloseHour string `json:"close_hour"`
}

type ShoppingLink struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

type BusinessProfileResponse struct {
	BusinessProfile entity.BusinessProfile `json:"business_profile"`
	PhoneNumber     string                 `json:"phone_number"`
}

type FAQRequest struct {
	FAQID     string `json:"faq_id"`
	AccountID string `json:"account_id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
}
