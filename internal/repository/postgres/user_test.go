package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/quippv/go-clean/internal/entity"
	"github.com/quippv/go-clean/internal/repository/postgres"
	"github.com/quippv/go-clean/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &entity.User{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	id, createdAt, updatedAt := utils.GenerateIDAndUnixMillis()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(id, user.Name, user.Email, user.Password, createdAt, updatedAt)

	mock.ExpectQuery("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = \\$1").
		WithArgs(id).
		WillReturnRows(rows)

	repo := postgres.NewUserRepository(db)
	userSelected, err := repo.GetUserById(id)

	assert.NoError(t, err)
	assert.NotNil(t, userSelected)
	assert.Equal(t, id, userSelected.ID)
	assert.Equal(t, user.Name, userSelected.Name)
	assert.Equal(t, user.Email, userSelected.Email)
	assert.Equal(t, user.Password, userSelected.Password)
	assert.Equal(t, createdAt, userSelected.CreatedAt)
	assert.Equal(t, updatedAt, userSelected.UpdatedAt)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &entity.User{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(sqlmock.AnyArg(), user.Name, user.Email, user.Password, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

	userRepo := postgres.NewUserRepository(db)

	_, err = userRepo.CreateUser(user)
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}
