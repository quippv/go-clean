package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/quippv/go-clean/internal/entity"
	"github.com/quippv/go-clean/utils"
)

type PostgresRepository struct {
	db *sql.DB
}

// creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func (r *PostgresRepository) GetUserById(id uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresRepository) CreateUser(user *entity.User) (uuid.UUID, error) {
	id, createdAt, updatedAt := utils.GenerateIDAndUnixMillis()
	query := "INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, id, user.Name, user.Email, user.Password, createdAt, updatedAt)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
