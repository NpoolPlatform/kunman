package cruder

type Any interface{}

const (
	EQ       = "eq"
	NEQ      = "neq"
	GT       = "gt"
	GTE      = "gte"
	LT       = "lt"
	LTE      = "lte"
	IN       = "in"
	LIKE     = "like"
	NIN      = "nin"
	BETWEEN  = "between"
	NBETWEEN = "nbetween"
	OVERLAP  = "overlap"
	NOVERLAP = "noverlap"
)

type Cond struct {
	Op  string
	Val Any
}
