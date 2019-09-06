package repository

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/itcupnvyk/backend-class/database/model"
)

// UserRepository interface
type UserRepository interface {
	Create(*model.User) error
	Update(*model.User) error
	FindAll() ([]model.User, error)
	FindByID(ID int64) (model.User, error)
}

// NewUserRepository return new user repository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

type userRepositoryImpl struct {
	db *sql.DB
}

func (r *userRepositoryImpl) Create(user *model.User) (err error) {
	query := `
		INSERT INTO users(id, name) VALUES(?, ?)
	`
	_, err = r.db.Exec(query, user.ID, user.Name)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *userRepositoryImpl) Update(user *model.User) (err error) {
	query := `
		UPDATE users SET name = ? WHERE id = ?
	`
	_, err = r.db.Exec(query, user.Name, user.ID)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *userRepositoryImpl) FindAll() (users []model.User, err error) {
	query := `
		SELECT id, name FROM users
	`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return
}

func (r *userRepositoryImpl) FindByID(ID int64) (user model.User, err error) {
	query := `
		SELECT id, name FROM users WHERE id = ?
	`
	err = r.db.QueryRow(query, ID).Scan(&user.ID, &user.Name)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
