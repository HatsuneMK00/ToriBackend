package entity

import (
	"database/sql"
)

type UserSummary struct {
	ID            uint
	StudentNumber sql.NullString
}
