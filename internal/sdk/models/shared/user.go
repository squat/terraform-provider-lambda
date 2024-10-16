// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UserStatus - Status of the user's account
type UserStatus string

const (
	UserStatusActive      UserStatus = "active"
	UserStatusDeactivated UserStatus = "deactivated"
)

func (e UserStatus) ToPointer() *UserStatus {
	return &e
}
func (e *UserStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "deactivated":
		*e = UserStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserStatus: %v", v)
	}
}

// User - Information about a user in your team
type User struct {
	// Unique identifier for the user
	ID string `json:"id"`
	// Email address of the user
	Email string `json:"email"`
	// Status of the user's account
	Status UserStatus `json:"status"`
}

func (o *User) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *User) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *User) GetStatus() UserStatus {
	if o == nil {
		return UserStatus("")
	}
	return o.Status
}
