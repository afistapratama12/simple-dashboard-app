package service

import (
	"context"
	"simple-dashboard-server/api/request"
	"simple-dashboard-server/config"
	mocks "simple-dashboard-server/mocks/repository"
	"simple-dashboard-server/model"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthService_Login(t *testing.T) {
	env := config.ENV{
		BaseClientURL: "http://test.test",
	}

	type args struct {
		ctx context.Context
		req request.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		err     error
		prepare func(*testing.T, error) AuthService
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.LoginRequest{
					Email:      "test@mail.com",
					Password:   "pratama12",
					KeepSignIn: true,
				},
			},

			err: nil,
			prepare: func(t *testing.T, err error) AuthService {
				userRepo := mocks.UserRepo{}
				userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(model.User{
					BaseModel: model.BaseModel{
						ID: "1",
					},
					Email:       "test@mail.com",
					Password:    "$2a$12$4dFEkMNeAhEOcFIYiqxj8.Wk19qVFqgu/r2RQNXVJ.yC.yNsFyqI2",
					ActivatedAt: func() *time.Time { t := time.Now(); return &t }(),
				}, nil)

				return NewAuthService(env, &userRepo, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.prepare(t, tt.err)
			_, err := s.Login(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestAuthService_Register(t *testing.T) {
	env := config.ENV{
		BaseClientURL: "http://test.test",
	}

	type args struct {
		ctx context.Context
		req request.RegisterRequest
	}
	tests := []struct {
		name    string
		args    args
		err     error
		prepare func(*testing.T, error) AuthService
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.RegisterRequest{
					Email:    "test@mail.com",
					Password: "pratama12",
				},
			},
			err: nil,
			prepare: func(t *testing.T, err error) AuthService {
				userRepo := mocks.UserRepo{}
				userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(model.User{}, nil)
				userRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

				notifRepo := mocks.NotifRepo{}

				notifRepo.On("NotifEmail", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				return NewAuthService(env, &userRepo, &notifRepo)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.prepare(t, tt.err)
			err := s.Register(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestAuthService_VerifyEmail(t *testing.T) {
	env := config.ENV{
		BaseClientURL: "http://test.test",
	}

	ValidateToken = func(tknStr string) (*model.Claims, error) {
		return &model.Claims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			},
			UserID: "1",
			Email:  "test@mail.com",
		}, nil
	}

	type args struct {
		ctx context.Context
		req request.VerifyEmailRequest
	}
	tests := []struct {
		name    string
		args    args
		err     error
		prepare func(*testing.T, error) AuthService
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.VerifyEmailRequest{
					Token: "token",
				},
			},
			err: nil,
			prepare: func(t *testing.T, err error) AuthService {
				userRepo := mocks.UserRepo{}

				userRepo.On("FindByID", mock.Anything, mock.Anything).Return(model.User{
					BaseModel: model.BaseModel{
						ID: "1",
					},
					Email: "test@mail.com",
				}, nil)

				userRepo.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				return NewAuthService(env, &userRepo, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.prepare(t, tt.err)
			_, err := s.VerifyEmail(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestAuthService_NotifForgotPassword(t *testing.T) {
	env := config.ENV{
		BaseClientURL: "http://test.test",
	}

	type args struct {
		ctx context.Context
		req request.ResetPasswordRequest
	}
	tests := []struct {
		name    string
		args    args
		err     error
		prepare func(*testing.T, error) AuthService
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.ResetPasswordRequest{
					Email: "test@mail.com",
				},
			},
			err: nil,
			prepare: func(t *testing.T, err error) AuthService {
				userRepo := mocks.UserRepo{}
				userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(model.User{
					BaseModel: model.BaseModel{
						ID: "1",
					},
					Email: "test@mail.com",
				}, nil)

				notifRepo := mocks.NotifRepo{}
				notifRepo.On("NotifEmail", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				return NewAuthService(env, &userRepo, &notifRepo)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.prepare(t, tt.err)
			err := s.NotifForgotPassword(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestAuthService_ResetPassword(t *testing.T) {
	env := config.ENV{
		BaseClientURL: "http://test.test",
	}

	ValidateToken = func(tknStr string) (*model.Claims, error) {
		return &model.Claims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			},
			UserID: "1",
			Email:  "test@mail.com",
		}, nil
	}

	type args struct {
		ctx context.Context
		req request.ResetPasswordConfirmRequest
	}
	tests := []struct {
		name    string
		args    args
		err     error
		prepare func(*testing.T, error) AuthService
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.ResetPasswordConfirmRequest{
					Token:    "token",
					Password: "pratama12",
				},
			},
			err: nil,
			prepare: func(t *testing.T, err error) AuthService {
				userRepo := mocks.UserRepo{}
				userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(model.User{
					BaseModel: model.BaseModel{
						ID: "1",
					},
					Email:       "test@mail.com",
					ActivatedAt: func() *time.Time { t := time.Now(); return &t }(),
				}, nil)

				userRepo.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				return NewAuthService(env, &userRepo, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.prepare(t, tt.err)
			err := s.ResetPassword(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.err, err)
		})
	}
}
