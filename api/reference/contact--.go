package reference

import (
	"github.com/aldesgroup/devotion-project-template/_include/reference/model"
	"github.com/aldesgroup/goald"
)

type Contact struct {
	goald.BusinessObject
	FirstName string
	LastName  string
}

func init() {
	model.Contact().SetNotPersisted()
}
