package repository

import (
	"database/sql"
	"errors"

	"taskapi/internal/model"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

var (
	ErrUserEmailExists = errors.New("email already exists")
	ErrUserNotFound    = errors.New("user not found")
)

func (r *UserRepo) Create(user *model.User) error {
	// id, email, name, password_hash, create_at, update_at
	err := r.db.QueryRow(
		`insert into users (email, name, password_hash) values ($1, $2, $3)
		returning id, created_at, updated_at`,
		user.Email,
		user.Name,
		user.PasswordHash).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		// 如果pg返回错误码23505，表示email已经注册过了。
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrUserEmailExists
		}
		return err
	}
	return nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRow(
		`select id, email, name, password_hash, created_at, updated_at
		from users
		where email = $1`, email).
		Scan(&user.ID, &user.Email, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
