package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rana-touseef11/go-chi-postgresql/internal/model"
	"github.com/rana-touseef11/go-chi-postgresql/pkg/constant"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])
	// for rows.Next() {
	// 	var u model.User
	// 	err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone)
	if err != nil {
		return nil, err
	}
	// 	users = append(users, u)
	// }
	return users, nil
}

func (r *UserRepository) Login(ctx context.Context, u model.User) (*model.User, error) {
	rows, err := r.db.Query(ctx,
		`SELECT * FROM users
		 WHERE email=$1`,
		u.Email)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, u model.User) (*model.User, error) {
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO users (name, email, phone)
		 VALUES ($1, $2, $3)
		 RETURNING id, created_at, updated_at`,
		u.Name,
		u.Email,
		u.Phone,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	rows, err := r.db.Query(ctx,
		`SELECT *
		 FROM users
		 WHERE id = $1`,
		id)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, id string, u model.User) (*model.User, error) {
	if err := r.db.QueryRow(ctx,
		`UPDATE users
	 	 SET name=$1, email=$2, phone=$3, address=$4
	 	 WHERE id=$5
		 RETURNING id`,
		u.Name,
		u.Email,
		u.Phone,
		u.Address,
		id).Scan(&u.ID); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	cmd, err := r.db.Exec(ctx,
		`UPDATE users
	 	 SET status=$1
	 	 WHERE id=$2 AND status!=$1`,
		constant.UserStatusDeleted,
		id)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("user not found or already deleted")
	}

	return nil
}
