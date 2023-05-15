package config

type Role byte

const (
	Client Role = 0
	Server Role = 1
)

type Config struct {
	Role       Role
	ServerAddr string
}
