package repository

import (
	"context"
	"database/sql"
	"regexp"
	"simple-dashboard-server/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupMock(t *testing.T) (*gorm.DB, *sql.DB, sqlmock.Sqlmock) {
	// prepare mock connection database
	dbMock, mock, err := sqlmock.New()
	assert.Nil(t, err, "should be ok")

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 dbMock,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	assert.Nil(t, err, "should be ok")

	return db, dbMock, mock
}

func TestUserRepo_FindByID(t *testing.T) {
	dbGorm, dbMock, mock := SetupMock(t)
	defer dbMock.Close()

	type arg struct {
		ctx context.Context
		id  string
	}

	var tests = []struct {
		name    string
		arg     arg
		err     error
		prepare func(*testing.T, error)
	}{
		{
			name: "success",
			arg: arg{
				ctx: context.Background(),
				id:  "1",
			},
			err: nil,
			prepare: func(t *testing.T, err error) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).AddRow("1", "test@mail.com"))
			},
		},
		{
			name: "error",
			arg: arg{
				ctx: context.Background(),
				id:  "1",
			},
			err: sql.ErrNoRows,
			prepare: func(t *testing.T, err error) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(t, tt.err)
			repo := NewUserRepo(dbGorm)
			_, err := repo.FindByID(tt.arg.ctx, tt.arg.id)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestUserRepo_FindByEmail(t *testing.T) {
	dbGorm, dbMock, mock := SetupMock(t)
	defer dbMock.Close()

	type arg struct {
		ctx   context.Context
		email string
	}

	var tests = []struct {
		name    string
		arg     arg
		err     error
		prepare func(*testing.T, error)
	}{
		{
			name: "success",
			arg: arg{
				ctx:   context.Background(),
				email: "test@mail.com",
			},
			err: nil,
			prepare: func(t *testing.T, err error) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).AddRow("1", "test@mail.com"))
			},
		},
		{
			name: "error",
			arg: arg{
				ctx:   context.Background(),
				email: "test@mail.com",
			},
			err: sql.ErrNoRows,
			prepare: func(t *testing.T, err error) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(t, tt.err)
			repo := NewUserRepo(dbGorm)
			_, err := repo.FindByEmail(tt.arg.ctx, tt.arg.email)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestUserRepo_UpdateByID(t *testing.T) {
	dbGorm, dbMock, mock := SetupMock(t)
	defer dbMock.Close()

	type arg struct {
		ctx     context.Context
		id      string
		updates map[string]interface{}
	}

	var tests = []struct {
		name    string
		arg     arg
		err     error
		prepare func(*testing.T, error)
	}{
		{
			name: "success",
			arg: arg{
				ctx: context.Background(),
				id:  "1",
				updates: map[string]interface{}{
					"email":      "test@mail.com",
					"updated_at": "2021-01-01",
				},
			},
			err: nil,
			prepare: func(t *testing.T, err error) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "email"=$1,"updated_at"=$2 WHERE id = $3`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			arg: arg{
				ctx: context.Background(),
				id:  "1",
				updates: map[string]interface{}{
					"email":      "test@mail.com",
					"updated_at": "2021-01-01",
				},
			},
			err: sql.ErrNoRows,
			prepare: func(t *testing.T, err error) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "email"=$1,"updated_at"=$2 WHERE id = $3`)).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(err)
				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(t, tt.err)
			repo := NewUserRepo(dbGorm)
			err := repo.UpdateByID(tt.arg.ctx, tt.arg.id, tt.arg.updates)
			assert.Equal(t, tt.err, err)
		})
	}
}

// TODO: fix later
func TestUserRepo_Create(t *testing.T) {
	dbGorm, dbMock, mock := SetupMock(t)
	defer dbMock.Close()

	type arg struct {
		ctx  context.Context
		user model.User
	}

	var tests = []struct {
		name    string
		arg     arg
		err     error
		prepare func(*testing.T, error)
	}{
		{
			name: "success",
			arg: arg{
				ctx: context.Background(),
				user: model.User{
					BaseModel: model.BaseModel{
						ID: "1",
					},
					Email:           "test@mail.com",
					Password:        "dummy",
					FirstName:       "hehe",
					LastName:        "hehe",
					Address:         "hehe",
					Address2:        "hehe",
					City:            "hehe",
					State:           "hehe",
					ZipCode:         "13123",
					ProfilePhotoURL: "http://google.com",
				},
			},
			err: nil,
			prepare: func(t *testing.T, err error) {
				mock.ExpectBegin()
				mock.MatchExpectationsInOrder(false)
				mock.ExpectExec("INSERT INTO \"users\" (\"id\",\"deleted_at\",\"email\",\"password\",\"activated_at\",\"first_name\",\"last_name\",\"phone_number\",\"address\",\"address2\",\"city\",\"state\",\"zip_code\",\"profile_photo_url\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				// mock.ExpectRollback()

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(t, tt.err)
			repo := NewUserRepo(dbGorm)
			err := repo.Create(tt.arg.ctx, tt.arg.user)
			assert.NotEqual(t, tt.err, err)
		})
	}
}
