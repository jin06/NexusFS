package command

type OPT uint8

const (
	READ  OPT = 0
	WRITE OPT = 1
	OPEN  OPT = 2
)

type Command struct {
	Opt OPT
}
