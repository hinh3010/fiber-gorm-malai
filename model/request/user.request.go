package request

type UserCreateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address"`
	Age     uint8  `json:"age"`
}
