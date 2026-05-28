package dto

type Operator string

const (
	OpEQ   Operator = "eq"
	OpLIKE Operator = "like"
	OpGT   Operator = "gt"
	OpGTE  Operator = "gte"
	OpLT   Operator = "lt"
	OpLTE  Operator = "lte"
)

var OperatorSQLMap = map[Operator]string{
	OpEQ:   "=",
	OpLIKE: "LIKE",
	OpGT:   ">",
	OpGTE:  ">=",
	OpLT:   "<",
	OpLTE:  "<=",
}

type FilterStructure struct {
	Field     string   `json:"field"`
	Operation Operator `json:"operator"`
	Value     any      `json:"value"`
}

type PaginationStructure struct {
	Filters []FilterStructure `query:"filters"`
	Limit   int               `query:"limit"`
	Offset  int               `query:"offset"`
}
