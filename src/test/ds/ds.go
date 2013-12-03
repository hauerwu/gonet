package ds

import (
	"test/ds/mongodb"
)

type DAO interface{
	Initialize(url string) error
	Finalize()
	Select(key interface{},result interface{}) error
	Update(key interface{},value interface{}) error
	Delete(key interface{}) error
	Insert(value interface{}) error
}

func New(name string) DAO{
	var d DAO

	switch name{
	case "mongodb":
		d = *mongodb.New()
	default:
		d = nil
	}

	return d
}














