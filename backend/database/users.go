package database

import (
	"context"
	"errors"
	"log"
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

/********************************** GET **********************************/

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

// GetUsers : get all the user
func GetUsers(ctx context.Context) ([]User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var users []User
	var user User
	var i int

	i = 0
	rows, err := pool.QueryContext(ctx,
		`SELECT user_id, tenant_id, username, password, email, created_at, updated_at, last_login
		FROM users `)

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.TenantID, &user.Username, &user.Password, &user.Email,
			&user.CreatedAt, &user.UpdatedAt, &user.LastLogin)
		if err != nil {
			return users, err
		}
		users = append(users, user)
		i = i + 1
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}

/********************************** GET **********************************/

/********************************** UPDATE **********************************/

// UpdateUser : update a user
func UpdateUser(ctx context.Context, userID int, tenantID int,
	username string, password string, email string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if (tenantID == 0) && (username == "") && (password == "") && (email == "") {
		return errors.New("invalid input fields (database/users.go")
	}

	// modify zoneID
	if tenantID != 0 {
		result, err := pool.ExecContext(ctx, `
			UPDATE users SET tenant_id = $1 
			WHERE user_id = $2
		`, tenantID, userID)

		if err != nil {
			return errors.New("error update user tenant_id")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : user tenant_id - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify username
	if username != "" {
		result, err := pool.ExecContext(ctx, `
			UPDATE users SET username = $1 
			WHERE user_id = $2
		`, username, userID)

		if err != nil {
			return errors.New("error update users username")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : users username - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify password
	if password != "" {
		result, err := pool.ExecContext(ctx, `
			UPDATE users SET password = $1 
			WHERE user_id = $2
		`, password, userID)

		if err != nil {
			return errors.New("error update users password")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : users password - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	// modify email
	if email != "" {
		result, err := pool.ExecContext(ctx, `
			UPDATE users SET email = $1 
			WHERE user_id = $2
		`, email, userID)

		if err != nil {
			return errors.New("error update user email")
		}

		// verify if there is one ou more rows affected
		rows, err := result.RowsAffected()
		if err != nil {
			return errors.New("error : user email - rows affected")
		}
		// checks the number of rows affected
		if rows < 0 {
			log.Fatalf("expected to affect 1 row, affected %d", rows)
		}
	}

	return nil
}

/********************************** UPDATE **********************************/
