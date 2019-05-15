package core

// IModelReader is an abstraction of reading Person data
type IModelReader interface {
	GetPeople() ([]Person, error)
	GetMatches() ([]Match, error)
}

// IPersonIDProvider is an abstraction of something that provides the person id
type IPersonIDProvider interface {
	GetPersonIDByAlias(alias string) (int, error)
}
