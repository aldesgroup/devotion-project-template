package reference

import (
	"github.com/aldesgroup/devotion-project-template/_include/reference/model"
	"github.com/aldesgroup/goald"
)

type Contact struct {
	goald.BusinessObject
	FirstName string `json:"firstName" io:"i*" desc:"the contact's first name"`
	LastName  string `json:"lastName"  io:"i*" desc:"the contact's last name"`
}

func init() {
	model.Contact().SetDescription("A contact")
	model.Contact().SetNotPersisted()
}
