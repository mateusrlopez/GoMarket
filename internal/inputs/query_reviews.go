package inputs

type QueryReviewsInput struct {
	UserID    string `form:"userId"`
	ProductID string `form:"productId"`
}
