package request

type CreateProfile struct {
	ID        uint   `json:"id"`
	ImageURL  string `json:"image_url"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
