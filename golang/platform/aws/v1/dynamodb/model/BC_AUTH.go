package model

type BC_AUTH struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Active       bool   `json:"active"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

func (_ *BC_AUTH) GetTableName() string {
	return "BC_AUTH"
}
