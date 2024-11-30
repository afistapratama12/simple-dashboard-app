package service

import (
	"context"
	"errors"
	"simple-dashboard-server/api/request"
	"simple-dashboard-server/api/response"
	"simple-dashboard-server/repository"
	"time"
)

type UserService interface {
	EditUserLogin(ctx context.Context, req request.EditUserRequest) error
	GetProfileUserLogin(ctx context.Context, userId string) (res response.UserResponse, err error)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) EditUserLogin(ctx context.Context, req request.EditUserRequest) error {
	user, err := s.userRepo.FindByID(ctx, req.ID)
	if err != nil && err.Error() != "record not found" {
		return err
	}

	if user.ID == "" {
		return errors.New("user not found")
	}

	// update user
	update := map[string]interface{}{
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"updated_at": time.Now(),
	}

	if req.PhoneNumber != "" {
		update["phone_number"] = req.PhoneNumber
	}

	if req.Address != "" {
		update["address"] = req.Address
	}

	if req.Address2 != "" {
		update["address2"] = req.Address2
	}

	if req.City != "" {
		update["city"] = req.City
	}

	if req.State != "" {
		update["state"] = req.State
	}

	if req.ZipCode != "" {
		update["zip_code"] = req.ZipCode
	}

	if err := s.userRepo.UpdateByID(ctx, req.ID, update); err != nil {
		return err
	}

	return nil

}

func (s *userService) GetProfileUserLogin(ctx context.Context, userId string) (res response.UserResponse, err error) {
	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil && err.Error() != "record not found" {
		return
	}

	if user.ID == "" {
		err = errors.New("user not found")
		return
	}

	res.Serialize(user)

	return
}
