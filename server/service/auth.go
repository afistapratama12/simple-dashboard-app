package service

import (
	"context"
	"errors"
	"fmt"
	"simple-dashboard-server/api/request"
	"simple-dashboard-server/api/response"
	"simple-dashboard-server/config"
	"simple-dashboard-server/helper"
	"simple-dashboard-server/model"
	"simple-dashboard-server/repository"
	"simple-dashboard-server/template"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, req request.LoginRequest) (res response.LoginResponse, err error)
	Register(ctx context.Context, req request.RegisterRequest) error
	VerifyEmail(ctx context.Context, req request.VerifyEmailRequest) (res response.LoginResponse, err error)
	NotifForgotPassword(ctx context.Context, req request.ResetPasswordRequest) error
	ResetPassword(ctx context.Context, req request.ResetPasswordConfirmRequest) error
}

type authService struct {
	env       config.ENV
	userRepo  repository.UserRepo
	notifRepo repository.NotifRepo
}

func NewAuthService(env config.ENV, userRepo repository.UserRepo, notifRepo repository.NotifRepo) AuthService {
	return &authService{
		env:       env,
		userRepo:  userRepo,
		notifRepo: notifRepo,
	}
}

func (s *authService) Login(ctx context.Context, req request.LoginRequest) (res response.LoginResponse, err error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return res, err
	}

	if user.ActivatedAt == nil {
		return res, errors.New("user not activated")
	}

	// validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return res, err
	}

	// generate token, done
	token, err := helper.GenerateToken(user.ID, user.Email, req.KeepSignIn)
	if err != nil {
		return res, err
	}

	res.Token = token
	res.UserID = user.ID
	res.Email = user.Email

	return res, nil
}

func (s *authService) Register(ctx context.Context, req request.RegisterRequest) error {
	// validate email tidak duplicate
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil && err.Error() != "record not found" {
		return err
	}

	if user.ID != "" {
		return errors.New("email already registered")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// create user
	user = model.User{
		BaseModel: model.BaseModel{
			ID: uuid.New().String(),
		},
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}

	token, err := helper.GenerateToken(user.ID, req.Email, false)
	if err != nil {
		return err
	}

	// send notif email
	err = s.notifRepo.NotifEmail(fmt.Sprintf("Verifikasi Email %s", req.FirstName),
		[]string{req.Email},
		template.GenerateEmailRegister(s.env.BaseClientURL, token))
	if err != nil {
		return err
	}

	return nil
}

var ValidateToken = helper.ValidateToken

func (s *authService) VerifyEmail(ctx context.Context, req request.VerifyEmailRequest) (res response.LoginResponse, err error) {
	// verify token
	claims, err := ValidateToken(req.Token)
	if err != nil {
		return
	}

	user, err := s.userRepo.FindByID(ctx, claims.UserID)
	if err != nil {
		return
	}

	if user.ActivatedAt != nil {
		return res, errors.New("email already verified")
	}

	if err := s.userRepo.UpdateByID(ctx, user.ID, map[string]interface{}{
		"activated_at": time.Now(),
		"updated_at":   time.Now(),
	}); err != nil {
		return res, err
	}

	token, err := helper.GenerateToken(user.ID, user.Email, false)
	if err != nil {
		return
	}

	res.Token = token
	res.UserID = user.ID
	res.Email = user.Email

	return
}

func (s *authService) NotifForgotPassword(ctx context.Context, req request.ResetPasswordRequest) error {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil && err.Error() != "record not found" {
		return err
	}

	if user.ID == "" {
		return errors.New("email not found")
	}

	token, err := helper.GenerateToken(user.ID, req.Email, false)
	if err != nil {
		return err
	}

	err = s.notifRepo.NotifEmail("Reset Password",
		[]string{req.Email},
		template.GenerateEmailForgotPassword(s.env.BaseClientURL, token))
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ResetPassword(ctx context.Context, req request.ResetPasswordConfirmRequest) error {
	// verify token
	claims, err := ValidateToken(req.Token)
	if err != nil {
		return err
	}

	user, err := s.userRepo.FindByEmail(ctx, claims.Email)
	if err != nil && err.Error() != "record not found" {
		return err
	}

	if user.ID == "" {
		return errors.New("user not found")
	}

	if user.ActivatedAt == nil {
		return errors.New("user not activated")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := s.userRepo.UpdateByID(ctx, user.ID, map[string]interface{}{
		"password":   string(hashedPassword),
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}

	return nil
}
