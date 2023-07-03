package inputs

type UpdateReviewInput struct {
	Title   string  `json:"title" binding:"required"`
	Comment string  `json:"comment" binding:"required"`
	Rating  float64 `json:"rating" binding:"required,min=0,max=5"`
}
