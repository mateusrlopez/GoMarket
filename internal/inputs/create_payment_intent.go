package inputs

type CreatePaymentIntentInput struct {
	OrderID string `json:"orderId" binding:"required"`
}
