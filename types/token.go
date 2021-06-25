package types

type TokensReturn struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenMetadataReturn struct {
	UUID   string
	UserId string
}
