package godrinth

type Config interface {
	Profiles() []Profile
}
