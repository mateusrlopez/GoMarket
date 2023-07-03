package inputs

type CreateReviewInput struct {
	Title     string  `json:"title" binding:"required"`
	Comment   string  `json:"comment" binding:"required"`
	Rating    float64 `json:"rating" binding:"required,min=0,max=5"`
	UserID    string  `json:"userId" binding:"required"`
	ProductID string  `json:"productId" binding:"required"`
}
