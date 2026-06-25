package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"sample/src/api/presenter"
	"sample/src/initializers"
	"sample/src/pkg/base"
	"sample/src/pkg/dto"
	"sample/src/pkg/utils"
	"sample/src/pkg/utils/customErrors"
)

func GetAll[L, D, C, U any](svc base.Service[L, D, C, U], filterable []string) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "handler."+svc.Name()+".GetAll")
		defer span.End()

		var page dto.PaginationStructure
		if err := c.Bind().Query(&page); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}
		if err := utils.ValidatePaginationQuery(c, &page, filterable); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}

		items, count, err := svc.GetAll(c.Context(), page)
		if err != nil {
			span.RecordError(err)
			var Err *customErrors.InternalError
			if errors.As(err, &Err) {
				return presenter.ReturnInternalError(c, Err)
			}
			return presenter.ReturnNotFound(c)
		}
		return presenter.SuccessfulPaginatedResponse(c, page, items, count)
	}
}

func GetOne[L, D, C, U any](svc base.Service[L, D, C, U]) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "handler."+svc.Name()+".GetOne")
		defer span.End()

		index, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return presenter.ReturnNotFound(c)
		}
		item, err := svc.GetOne(c.Context(), uint(index))
		if err != nil {
			span.RecordError(err)
			var Err *customErrors.InternalError
			if errors.As(err, &Err) {
				return presenter.ReturnInternalError(c, Err)
			}
			return presenter.ReturnNotFound(c)
		}
		return presenter.SuccessfulResponse(c, item)
	}
}

func Create[L, D, C, U any](svc base.Service[L, D, C, U]) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "handler."+svc.Name()+".Create")
		defer span.End()

		var payload C
		if err := c.Bind().Body(&payload); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}
		item, err := svc.Create(c.Context(), &payload)
		if err != nil {
			span.RecordError(err)
			var Err *customErrors.InternalError
			if errors.As(err, &Err) {
				return presenter.ReturnInternalError(c, Err)
			}
			return presenter.InputSerializerError(c)
		}
		return presenter.SuccessfulResponse(c, item)
	}
}

func Update[L, D, C, U any](svc base.Service[L, D, C, U]) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "handler."+svc.Name()+".Update")
		defer span.End()

		index, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return presenter.ReturnNotFound(c)
		}
		var payload U
		if err := c.Bind().Body(&payload); err != nil {
			span.RecordError(err)
			return presenter.InputSerializerError(c)
		}
		if err := svc.Update(c.Context(), uint(index), &payload); err != nil {
			span.RecordError(err)
			var Err *customErrors.InternalError
			if errors.As(err, &Err) {
				return presenter.ReturnInternalError(c, Err)
			}
			return presenter.ReturnNotFound(c)
		}
		return presenter.SuccessfulUpdateResponse(c)
	}
}

func Delete[L, D, C, U any](svc base.Service[L, D, C, U]) fiber.Handler {
	return func(c fiber.Ctx) error {
		_, span := initializers.Tracer.Start(c.Context(), "handler."+svc.Name()+".Delete")
		defer span.End()

		index, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return presenter.ReturnNotFound(c)
		}
		if err := svc.Delete(c.Context(), uint(index)); err != nil {
			span.RecordError(err)
			var Err *customErrors.InternalError
			if errors.As(err, &Err) {
				return presenter.ReturnInternalError(c, Err)
			}
			return presenter.ReturnNotFound(c)
		}
		return presenter.SuccessfulDeleteResponse(c)
	}
}
