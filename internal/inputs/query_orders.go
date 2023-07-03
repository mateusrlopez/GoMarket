package inputs

type QueryOrdersInput struct {
	UserID string `form:"userId"`
	Status string `form:"status"`
}
