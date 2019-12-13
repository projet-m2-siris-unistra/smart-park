package database

import (
	"context"
	"errors"
	"log"
	"database/sql"
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

// UserResponse returns the id of the updated / created object 
type UserResponse struct {
	UserID int `json:"user_id"`
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
func GetUsers(ctx context.Context, limite int, offset int) ([]User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var users []User
	var user User
	var i int

	limite, offset = CheckArgUser(limite, offset)

	i = 0

	rows, err := pool.QueryContext(ctx,
		`SELECT user_id, tenant_id, username, password, email, created_at, updated_at, last_login
		FROM users LIMIT $1 OFFSET $2`, limite, offset)

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
	username string, password string, email string) (UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user UserResponse

	user.UserID = -1

	if (tenantID == 0) && (username == "") && (password == "") && (email == "") {
		return user, errors.New("invalid input fields (database/users.go")
	}

	// modify zoneID
	if tenantID != 0 {
		err := pool.QueryRowContext(ctx, `
			UPDATE users SET tenant_id = $1 
			WHERE user_id = $2 RETURNING user_id
		`, tenantID, userID).Scan(&user.UserID)
		
		if err == sql.ErrNoRows {
			log.Printf("no user with id %d\n", userID)
			return user, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return user, err
		}
		
	}

	// modify username
	if username != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE users SET username = $1 
			WHERE user_id = $2 RETURNING user_id
		`, username, userID).Scan(&user.UserID)

		if err == sql.ErrNoRows {
			log.Printf("no user with id %d\n", userID)
			return user, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return user, err
		}
	}

	// modify password
	if password != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE users SET password = $1 
			WHERE user_id = $2 RETURNING user_id
		`, password, userID).Scan(&user.UserID)

		if err == sql.ErrNoRows {
			log.Printf("no user with id %d\n", userID)
			return user, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return user, err
		}
	}

	// modify email
	if email != "" {
		err := pool.QueryRowContext(ctx, `
			UPDATE users SET email = $1 
			WHERE user_id = $2 RETURNING user_id
		`, email, userID).Scan(&user.UserID)

		if err == sql.ErrNoRows {
			log.Printf("no user with id %d\n", userID)
			return user, err
		}

		if err != nil {
			log.Printf("query error: %v\n", err)
			return user, err
		}
	}

	return user, nil
}

/********************************** UPDATE **********************************/

/********************************** OPTIONS **********************************/


// CountUser : count number of rows
func CountUser(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	var count int

	count = -1

	row := pool.QueryRow("SELECT COUNT(*) FROM users")
	err := row.Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

// CheckArgUser : check limit and offset arguments
func CheckArgUser(limite int, offset int) (int, int) {

	if limite == 0 {
		limite = 20
	}

	if offset == 0 {
		offset = 0
	}

	return limite, offset
}

/********************************** OPTIONS **********************************/