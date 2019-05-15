package core

// IPersonIDProvider is an abstraction of something that provides the person id
type IPersonIDProvider interface {
	GetPersonIDByAlias(alias string) (int, error)
}
