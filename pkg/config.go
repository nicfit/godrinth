package godrinth

type Config interface {
	Profiles() []string
}
