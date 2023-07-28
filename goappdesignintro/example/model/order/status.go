package order

type Status int

const (
	StatusNew Status = iota
	StatusInProgress
	StatusFail
	StatusReady
)
