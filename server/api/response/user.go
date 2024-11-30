package response

import "simple-dashboard-server/model"

type UserResponse struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	Address         string `json:"address"`
	Address2        string `json:"address2"`
	City            string `json:"city"`
	State           string `json:"state"`
	ZipCode         string `json:"zip_code"`
	ProfilePhotoURL string `json:"profile_photo_url"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (u *UserResponse) Serialize(user model.User) {
	u.ID = user.ID
	u.Email = user.Email
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.PhoneNumber = user.PhoneNumber
	u.Address = user.Address
	u.Address2 = user.Address2
	u.City = user.City
	u.State = user.State
	u.ZipCode = user.ZipCode
	u.ProfilePhotoURL = user.ProfilePhotoURL
	u.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	u.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")
}
