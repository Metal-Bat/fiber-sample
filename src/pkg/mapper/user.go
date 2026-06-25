package mapper

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen .

import (
	"sample/src/pkg/dto"
	"sample/src/pkg/entities"
)

// goverter:converter
// goverter:output:file user_generated.go
// goverter:extend TimeToTime
// goverter:extend NullStringToPtr
// goverter:extend PtrToNullString
// goverter:extend NullTimeToPtr
// goverter:extend PtrToNullTime
// goverter:extend GormDeletedAtToPtr
// goverter:extend PermissionsToNames
// goverter:extend PtrToString
type UserConverter interface {
	// goverter:map BaseModel.ID ID
	ToList(src *entities.User) *dto.UserInfo

	// goverter:autoMap BaseModel
	ToDetail(src *entities.User) *dto.UserDetail

	ToListSlice(src []*entities.User) []*dto.UserInfo

	// goverter:ignore BaseModel Password Permissions
	ToEntity(src *dto.CreateUser) *entities.User

	// goverter:ignore BaseModel Mobile Permissions
	ToUpdateEntity(src *dto.UpdateUser) *entities.User
}
