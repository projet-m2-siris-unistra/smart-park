package database

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v3"
)

// User describes a user in the database
type User struct {
	UserID    int       `json:"user_id"`
	TenantID  int       `json:"tenant_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	LastLogin null.Time `json:"last_login"`
	Timestamps
}

// GetUser fetches the user by its ID
func GetUser(ctx context.Context, userID int) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user User

	err := pool.QueryRowContext(ctx, `
		SELECT user_id, tenant_id, username, password, email, created_at, updated_at, last_login
		FROM users 
		WHERE user_id = $1
	`, userID).
		Scan(&user.UserID, &user.TenantID, &user.Username, &user.Password, &user.Email,
			&user.CreatedAt, &user.UpdatedAt, &user.LastLogin)

	if err != nil {
		return user, err
	}

	return user, nil
}
