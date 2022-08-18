package entity

type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	AtExpires   int64
}

type AccessDetail struct {
	AccessUuid string `json:"accessUuid"`
	Role       uint   `json:"role"`
	Username   string `json:"userName"`
}
