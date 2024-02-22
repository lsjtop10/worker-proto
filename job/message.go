package job

type Stage int

const (
	All Stage = iota + 1
	Fetch
	ModelExec
	ModelResAnalyze
	Transmit
)

type Flag int

const (
	ChannelClosed Flag = iota + 1
)

type Message struct {
	Producer Stage
	Consumer Stage

	Msg Flag
}
