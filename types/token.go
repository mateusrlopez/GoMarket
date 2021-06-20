package types

type TokensReturn struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenMetadataReturn struct {
	UUID   string
	UserId float64
}
