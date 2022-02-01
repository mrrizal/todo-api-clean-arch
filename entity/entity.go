package entity

import (
	"errors"
	"fmt"
	"net/mail"
	"time"
)

type ActivityGroups struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var Priorities = []string{"very-low", "low", "medium", "high", "very-high"}

func (input ActivityGroups) Validate() error {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	if input.Title == "" {
		return errors.New("title cannot be empty")
	}
	return nil
}

type Todo struct {
	ID              int        `json:"id"`
	ActivityGroupID int        `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        bool       `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}

func contains(priorities []string, priority string) bool {
	for _, val := range priorities {
		if val == priority {
			return true
		}
	}
	return false
}

func (input Todo) Validate() error {
	if input.ActivityGroupID == 0 {
		return errors.New("activity group cannot be empty")
	}

	if input.Title == "" {
		return errors.New("title cannot be empty")
	}

	if !contains(Priorities, input.Priority) {
		return fmt.Errorf("invalid value, priority must be one of %v", input.Priority)
	}

	return nil
}
