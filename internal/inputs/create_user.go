package inputs

type CreateUserInput struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Addresses []struct {
		City       string `json:"city" binding:"required"`
		State      string `json:"state" binding:"required"`
		Country    string `json:"country" binding:"required"`
		PostalCode string `json:"postalCode" binding:"required"`
		Line1      string `json:"line1" binding:"required"`
		Line2      string `json:"line2"`
	} `json:"addresses"`
}
