package customErrors

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// https://www.postgresql.org/docs/current/errcodes-appendix.html
type PGCode string

const (
	// Class 08 — Connection Exception
	PGConnectionException                  PGCode = "08000"
	PGConnectionDoesNotExist               PGCode = "08003"
	PGConnectionFailure                    PGCode = "08006"
	PGSQLClientUnableToEstablishConnection PGCode = "08001"
	PGSQLServerRejectedEstablishment       PGCode = "08004"
	PGTransactionResolutionUnknown         PGCode = "08007"

	// Class 22 — Data Exception
	PGStringDataRightTruncation PGCode = "22001"
	PGNumericValueOutOfRange    PGCode = "22003"
	PGInvalidDatetimeFormat     PGCode = "22007"
	PGDatetimeFieldOverflow     PGCode = "22008"
	PGInvalidEscapeSequence     PGCode = "22025"
	PGZeroDivision              PGCode = "22012"

	// Class 23 — Integrity Constraint Violation
	PGIntegrityConstraintViolation PGCode = "23000"
	PGRestrictViolation            PGCode = "23001"
	PGNotNullViolation             PGCode = "23502"
	PGForeignKeyViolation          PGCode = "23503"
	PGUniqueViolation              PGCode = "23505"
	PGCheckViolation               PGCode = "23514"
	PGExclusionViolation           PGCode = "23P01"

	// Class 40 — Transaction Rollback
	PGTransactionRollback  PGCode = "40000"
	PGSerializationFailure PGCode = "40001"
	PGDeadlockDetected     PGCode = "40P01"

	// Class 42 — Syntax Error or Access Rule Violation
	PGUndefinedColumn       PGCode = "42703"
	PGUndefinedTable        PGCode = "42P01"
	PGDuplicateTable        PGCode = "42P07"
	PGDuplicateColumn       PGCode = "42701"
	PGInsufficientPrivilege PGCode = "42501"

	// Class 53 — Insufficient Resources
	PGInsufficientResources PGCode = "53000"
	PGDiskFull              PGCode = "53100"
	PGOutOfMemory           PGCode = "53200"
	PGTooManyConnections    PGCode = "53300"

	// Class 57 — Operator Intervention
	PGQueryCanceled    PGCode = "57014"
	PGAdminShutdown    PGCode = "57P01"
	PGCannotConnectNow PGCode = "57P03"

	// Class 58 — System Error
	PGIOError PGCode = "58030"
)

var pgErrorMap = map[PGCode]InternalError{
	// Connection
	PGConnectionException:                  {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_connection"},
	PGConnectionDoesNotExist:               {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_connection"},
	PGConnectionFailure:                    {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_connection"},
	PGSQLClientUnableToEstablishConnection: {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_connection"},
	PGSQLServerRejectedEstablishment:       {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_connection"},

	// Data
	PGStringDataRightTruncation: {HTTPStatus: http.StatusUnprocessableEntity, MessageKey: "db_error_value_too_long"},
	PGNumericValueOutOfRange:    {HTTPStatus: http.StatusUnprocessableEntity, MessageKey: "db_error_value_out_of_range"},
	PGInvalidDatetimeFormat:     {HTTPStatus: http.StatusUnprocessableEntity, MessageKey: "db_error_invalid_datetime"},
	PGDatetimeFieldOverflow:     {HTTPStatus: http.StatusUnprocessableEntity, MessageKey: "db_error_invalid_datetime"},
	PGZeroDivision:              {HTTPStatus: http.StatusUnprocessableEntity, MessageKey: "db_error_zero_division"},

	// Integrity
	PGIntegrityConstraintViolation: {HTTPStatus: http.StatusConflict, MessageKey: "db_error_constraint_violation"},
	PGRestrictViolation:            {HTTPStatus: http.StatusConflict, MessageKey: "db_error_constraint_violation"},
	PGNotNullViolation:             {HTTPStatus: http.StatusUnprocessableEntity, MessageKey: "db_error_not_null_violation"},
	PGForeignKeyViolation:          {HTTPStatus: http.StatusConflict, MessageKey: "db_error_foreign_key_violation"},
	PGUniqueViolation:              {HTTPStatus: http.StatusConflict, MessageKey: "db_error_unique_violation"},
	PGCheckViolation:               {HTTPStatus: http.StatusUnprocessableEntity, MessageKey: "db_error_check_violation"},
	PGExclusionViolation:           {HTTPStatus: http.StatusConflict, MessageKey: "db_error_constraint_violation"},

	// Transaction
	PGSerializationFailure: {HTTPStatus: http.StatusConflict, MessageKey: "db_error_serialization_failure"},
	PGDeadlockDetected:     {HTTPStatus: http.StatusConflict, MessageKey: "db_error_deadlock"},

	// Access
	PGUndefinedColumn:       {HTTPStatus: http.StatusInternalServerError, MessageKey: "db_error_internal"},
	PGUndefinedTable:        {HTTPStatus: http.StatusInternalServerError, MessageKey: "db_error_internal"},
	PGInsufficientPrivilege: {HTTPStatus: http.StatusForbidden, MessageKey: "db_error_insufficient_privilege"},

	// Resources
	PGInsufficientResources: {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_insufficient_resources"},
	PGDiskFull:              {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_insufficient_resources"},
	PGOutOfMemory:           {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_insufficient_resources"},
	PGTooManyConnections:    {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_too_many_connections"},

	// Operator
	PGQueryCanceled:    {HTTPStatus: http.StatusRequestTimeout, MessageKey: "db_error_query_canceled"},
	PGAdminShutdown:    {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_connection"},
	PGCannotConnectNow: {HTTPStatus: http.StatusServiceUnavailable, MessageKey: "db_error_connection"},
}

func MapError(err error) *InternalError {
	if err == nil {
		return nil
	}

	// GORM sentinel errors
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &InternalError{HTTPStatus: http.StatusNotFound, MessageKey: "entity not found", Err: err}
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return &InternalError{HTTPStatus: http.StatusConflict, MessageKey: "db_error_unique_violation", Err: err}
	}
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return &InternalError{HTTPStatus: http.StatusConflict, MessageKey: "db_error_foreign_key_violation", Err: err}
	}
	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		return &InternalError{HTTPStatus: http.StatusUnprocessableEntity, MessageKey: "db_error_check_violation", Err: err}
	}

	// PostgreSQL native SQLSTATE via pgconn
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if template, ok := pgErrorMap[PGCode(pgErr.Code)]; ok {
			return &InternalError{
				HTTPStatus: template.HTTPStatus,
				MessageKey: template.MessageKey,
				Err:        err,
			}
		}
	}

	return &InternalError{HTTPStatus: http.StatusInternalServerError, MessageKey: "db_error_internal", Err: err}
}
