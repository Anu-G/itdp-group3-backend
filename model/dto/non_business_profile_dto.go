package dto

import "itdp-group3-backend/model/entity"

type NonBusinessProfileRequest struct {
	AccountID    string `json:"account_id"`
	ProfileImage string `json:"profile_image"`
	ProfileBio   string `json:"profile_bio"`
	DisplayName  string `json:"display_name"`
}

type NonBusinessProfileResponse struct {
	NonBusinessProfile entity.NonBusinessProfile `json:"non_business_profile"`
	PhoneNumber        string                    `json:"phone_number"`
}
