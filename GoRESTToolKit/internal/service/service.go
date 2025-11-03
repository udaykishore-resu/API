package service

import (
	"errors"
	"go-rest-toolkit/internal/db"
	"go-rest-toolkit/internal/model"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateUser(user *model.User) error {
	query := "INSERT INTO users (id, username, email, password, roles) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.DB.Exec(query, user.ID, user.Username, user.Email, user.Password, user.Roles)
	return err
}

func (s *Service) GetUser(id string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, username, email, roles FROM users WHERE id=$1"
	err := db.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Roles)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *Service) AuthenticateUser(email, password string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, username, email, password, roles FROM users WHERE email=?"
	row := db.DB.QueryRow(query, email)
	var hashedPassword string
	var roles string

	err := row.Scan(&user.ID, &user.Username, &user.Email, &hashedPassword, &roles)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Password compare
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	user.Roles = strings.Split(roles, ",")
	return user, nil
}

func (s *Service) CreateRefreshToken(userID string) (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	refreshToken := base64.URLEncoding.EncodeToString(tokenBytes)

	expiresAt := time.Now().Add(24 * time.Hour) // Valid 24 hours or as per your policy

	query := "INSERT INTO refresh_tokens (token, user_id, expires_at) VALUES (?, ?, ?)"
	_, err = db.DB.Exec(query, refreshToken, userID, expiresAt)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (s *Service) ValidateRefreshToken(token string) (string, error) {
	var userID string
	var expiresAt time.Time

	query := "SELECT user_id, expires_at FROM refresh_tokens WHERE token=?"
	err := db.DB.QueryRow(query, token).Scan(&userID, &expiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid refresh token")
		}
		return "", err
	}

	if time.Now().After(expiresAt) {
		return "", errors.New("refresh token expired")
	}

	return userID, nil
}
