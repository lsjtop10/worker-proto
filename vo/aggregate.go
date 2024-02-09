package vo

type ProductAggregate struct {
}

type Action int32

const (
	Buy  Action = 1
	Sell Action = -1
	Hold Action = 0
)
