package mapper

import (
	"database/sql"
	"time"

	"gorm.io/gorm"

	"sample/src/pkg/entities"
)

func TimeToTime(t time.Time) time.Time { return t }

func NullStringToPtr(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

func PtrToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func NullTimeToPtr(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

func PtrToNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func GormDeletedAtToPtr(t gorm.DeletedAt) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

func PermissionsToNames(perms []entities.Permission) []string {
	names := make([]string, len(perms))
	for i, p := range perms {
		names[i] = p.Name
	}
	return names
}

func PtrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
