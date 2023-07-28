package user

import (
	"time"

	"github.com/evermos/boilerplate-go/shared/nuuid"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type Role int

const (
	Admin Role = iota
	Regular
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time   `db:"updated"`
	DeletedAt null.Time   `db:"deleted"`
	UpdatedBy nuuid.NUUID `db:"updated_by"`
	DeletedBy nuuid.NUUID `db:"deleted_by"`
}

type PayloadUser struct {
	Username string
	Email    string
	Role     string
}

type UserResponseFormat struct {
	ID        uuid.UUID `json:"userId"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt null.Time `json:"deletedAt"`
	UpdatedBy uuid.UUID `json:"updatedBy"`
	DeletedBy uuid.UUID `json:"deletedBy"`
}
