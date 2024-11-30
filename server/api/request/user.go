package request

type EditUserRequest struct {
	ID              string `json:"-"`
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	PhoneNumber     string `json:"phone_number"`
	Address         string `json:"address"`
	Address2        string `json:"address2"`
	City            string `json:"city"`
	State           string `json:"state"`
	ZipCode         string `json:"zip_code"`
	ProfilePhotoURL string `json:"profile_photo_url"`
}
