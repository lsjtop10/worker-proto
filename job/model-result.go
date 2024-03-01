package job

type Action int32

const (
	Buy  Action = -1
	None Action = 0
	Sell Action = 1
)

type UpDownProbDist [3]float32

type Prices []float32
