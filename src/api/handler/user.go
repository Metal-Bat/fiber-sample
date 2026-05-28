package handler

import (
	"sample/src/api/presenter"
	"sample/src/initializers"
	"sample/src/pkg/dto"
	"sample/src/pkg/internals/user"
	"sample/src/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func GetUsers(service user.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "presenter.GetUsers")
		defer span.End()

		var page dto.PaginationStructure
		if err := c.Bind().Query(&page); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}

		filterable := []string{"id", "mobile", "national_code", "created_at", "updated_at"}
		if err := utils.ValidatePaginationQuery(c, &page, filterable); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}

		items, count, err := service.GetUsers(c, page)
		if err != nil {
			span.RecordError(err)
			return presenter.ReturnNotFound(c)
		}

		return presenter.SuccessfulPaginatedResponse(c, page, items, count)
	}
}

func GetUser(service user.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "presenter.GetUser")
		defer span.End()

		idParam := c.Params("id")

		index, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			return presenter.ReturnNotFound(c)
		}

		item, err := service.GetUser(c, uint(index))
		if err != nil {
			span.RecordError(err)
			return presenter.ReturnNotFound(c)
		}

		return presenter.SuccessfulResponse(c, item)
	}
}

func CreateUser(service user.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "presenter.CreateUser")
		defer span.End()

		var payload dto.CreateUser
		if err := c.Bind().Body(&payload); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}

		item, err := service.CreateUser(c, &payload)
		if err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}

		return presenter.SuccessfulResponse(c, item)
	}
}

func UpdateUser(service user.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "presenter.UpdateUser")
		defer span.End()

		idParam := c.Params("id")
		index, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			return presenter.ReturnNotFound(c)
		}

		var payload dto.UpdateUser
		if err := c.Bind().Body(&payload); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}

		err = service.UpdateUser(c, uint(index), &payload)
		if err != nil {
			span.RecordError(err)
			return presenter.ReturnNotFound(c)
		}

		return presenter.SuccessfulUpdateResponse(c)
	}
}

func DeleteUser(service user.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "presenter.DeleteUser")
		defer span.End()

		idParam := c.Params("id")
		index, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			return presenter.ReturnNotFound(c)
		}

		if err := service.DeleteUser(c, uint(index)); err != nil {
			span.RecordError(err)
			return presenter.ReturnNotFound(c)
		}

		return presenter.SuccessfulDeleteResponse(c)
	}
}

func Login(service user.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "presenter.Login")
		defer span.End()

		var input presenter.LoginInput
		if err := c.Bind().Body(&input); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}

		result, err := service.Login(c, input.Mobile, input.Password)
		if err != nil {
			return presenter.InvalidUserPassError(c)
		}
		return presenter.SuccessfulResponse(c, result)
	}

}
