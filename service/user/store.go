package user

import (
	"database/sql"
	"enfermeria_go/types"
	"fmt"
)

// Storage es el repositorio de usuarios
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.User) error {
	//_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)
	_, err := s.db.Exec("INSERT INTO public.users (id, firstname, lastname, email, \"password\", created_at, updated_at) VALUES(nextval('users_id_seq'::regclass), $1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)",
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {

	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Create_at,
		&user.Update_at,
		//&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
