package utils

import (
	"encoding/json"
	"sample/src/initializers"
	"sample/src/pkg/dto"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func ValidatePaginationQuery(c fiber.Ctx, page *dto.PaginationStructure, allowed []string) error {
	_, span := initializers.Tracer.Start(c.Context(), "utils.ValidatePaginationQuery")
	defer span.End()

	rawFilters := c.Query("filters")
	if rawFilters != "" {

		normalized := "[" + rawFilters + "]"
		if err := json.Unmarshal([]byte(normalized), &page.Filters); err != nil {
			span.RecordError(err)
			return fmt.Errorf("invalid filters format")
		}
	}

	allowedMap := map[string]struct{}{}
	for _, f := range allowed {
		allowedMap[f] = struct{}{}
	}

	for _, f := range page.Filters {
		if _, ok := allowedMap[f.Field]; !ok {
			return fmt.Errorf("field %s is not filterable", f.Field)
		}

		if _, ok := dto.OperatorSQLMap[f.Operation]; !ok {
			return fmt.Errorf("operator %s is invalid", f.Operation)
		}
	}

	return nil
}

func ApplyFilters[T any](
	c fiber.Ctx,
	q gorm.ChainInterface[T],
	page dto.PaginationStructure,
) gorm.ChainInterface[T] {

	for _, f := range page.Filters {
		sqlOp := dto.OperatorSQLMap[f.Operation]

		val := f.Value
		if f.Operation == dto.OpLIKE {
			val = "%" + fmt.Sprint(val) + "%"
		}

		q = q.Where(
			fmt.Sprintf("%s %s ?", f.Field, sqlOp),
			val,
		)
	}

	return q
}

func ApplyLimitsAndOffset[T any](
	c fiber.Ctx,
	q gorm.ChainInterface[T],
	page dto.PaginationStructure,
) gorm.ChainInterface[T] {
	_, span := initializers.Tracer.Start(c.Context(), "utils.ApplyLimits")
	defer span.End()
	return q.Offset(page.Offset).Limit(page.Limit)
}
