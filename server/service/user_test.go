package service

import (
	"context"
	"simple-dashboard-server/api/request"
	mocks "simple-dashboard-server/mocks/repository"
	"simple-dashboard-server/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserservice_EditUserLogin(t *testing.T) {
	type args struct {
		ctx context.Context
		req request.EditUserRequest
	}
	tests := []struct {
		name    string
		args    args
		err     error
		prepare func(*testing.T, error) UserService
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.EditUserRequest{
					ID:        "1",
					FirstName: "test",
					LastName:  "test",
				},
			},
			err: nil,
			prepare: func(t *testing.T, err error) UserService {
				userRepo := mocks.UserRepo{}
				userRepo.On("FindByID", mock.Anything, mock.Anything).Return(model.User{
					BaseModel: model.BaseModel{
						ID: "1",
					},
					Email:       "test@mail.com",
					ActivatedAt: func() *time.Time { t := time.Now(); return &t }(),
				}, nil)

				userRepo.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything).Return(nil)

				return NewUserService(&userRepo)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.prepare(t, tt.err)

			err := s.EditUserLogin(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestUserservice_GetProfileUserLogin(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		args    args
		err     error
		prepare func(*testing.T, error) UserService
	}{
		{
			name: "success",
			args: args{
				ctx:    context.Background(),
				userId: "1",
			},
			err: nil,
			prepare: func(t *testing.T, err error) UserService {
				userRepo := mocks.UserRepo{}
				userRepo.On("FindByID", mock.Anything, mock.Anything).Return(model.User{
					BaseModel: model.BaseModel{
						ID: "1",
					},
					Email:       "test@mail.com",
					ActivatedAt: func() *time.Time { t := time.Now(); return &t }(),

					FirstName: "test",
					LastName:  "test",
				}, nil)

				return NewUserService(&userRepo)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.prepare(t, tt.err)

			_, err := s.GetProfileUserLogin(tt.args.ctx, tt.args.userId)
			assert.Equal(t, tt.err, err)
		})
	}
}
