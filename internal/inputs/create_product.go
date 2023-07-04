package inputs

type CreateProductInput struct {
	Name           string `json:"name" binding:"required"`
	Type           string `json:"type" binding:"required"`
	Specifications []struct {
		Name  string `json:"name" binding:"required"`
		Value string `json:"value" binding:"required"`
	} `json:"specifications"`
	Customizations []struct {
		Name    string `json:"name" binding:"required"`
		Options []struct {
			Name          string `json:"name" binding:"required"`
			PriceIncrease struct {
				Amount   int64  `json:"amount" binding:"required"`
				Currency string `json:"currency" binding:"required"`
			} `json:"priceIncrease" binding:"required"`
		} `json:"options" binding:"required"`
	} `json:"customizations"`
	BasePrice struct {
		Amount   int64  `json:"amount" binding:"required"`
		Currency string `json:"currency" binding:"required"`
	} `json:"basePrice" binding:"required"`
}
