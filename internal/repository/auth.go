package repository

import (
	"database/sql"
	"miras/internal/models"
)

type Auth struct {
	db *sql.DB
}

func newAuthRepo(db *sql.DB) *Auth {
	return &Auth{db: db}
}
func (r *Auth) CreateUser(user models.Register) error {

	query := `INSERT INTO users(username,email,hash_password,role) VALUES($1,$2,$3,$4)`

	_, err := r.db.Exec(query, &user.Username, &user.Email, &user.Password, "user")
	if err != nil {
		return err
	}

	return nil
}

func (r *Auth) GetUserById(id int) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id=$1`

	row := r.db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *Auth) SelectUser(user models.Login) (models.User, error) {

	var result models.User

	query := `SELECT * FROM users WHERE email=$1 `

	row := r.db.QueryRow(query, &user.Email)

	err := row.Scan(&result.ID, &result.Username, &result.Email, &result.Password, &result.Role)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func (r *Auth) CreateSession(session models.Session) error {
	query := `INSERT INTO sessions(user_id,token,expired_date) VALUES($1,$2,$3)`
	_, err := r.db.Exec(query, session.UserID, session.Token, session.ExpiredDate)
	if err != nil {
		return err
	}
	return nil

}

func (r *Auth) GetSessionByToken(token string) (models.Session, error) {

	var session models.Session

	query := `SELECT * FROM sessions WHERE token=$1`

	row := r.db.QueryRow(query, token)
	err := row.Scan(&session.ID, &session.UserID, &session.Token, &session.ExpiredDate)

	if err != nil {
		return models.Session{}, err
	}
	return session, nil
}

func (r *Auth) GetSessionByUserID(id int64) (models.Session, error) {

	var session models.Session

	query := `SELECT * FROM sessions WHERE user_id=$1`

	row := r.db.QueryRow(query, id)
	err := row.Scan(&session.ID, &session.UserID, &session.Token, &session.ExpiredDate)

	if err != nil {
		return models.Session{}, err
	}
	return session, nil
}

func (r *Auth) UpdateToken(token, date string, userID int) error {

	query := `UPDATE sessions SET token=$1 ,expaired_date=$2 WHERE user_id=$3`

	_, err := r.db.Exec(query, token, date, userID)
	if err != nil {
		return err
	}
	return nil
}
