package reference

import (
	specs "github.com/aldesgroup/devotion-project-template/_include/_specs"
	"github.com/aldesgroup/goald"
)

type Contact struct {
	goald.BusinessObject
	FirstName string
	LastName  string
}

func init() {
	specs.Contact().SetNotPersisted()
}
