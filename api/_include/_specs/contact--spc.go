// Generated file, do not edit!
package specs

import (
	"sync"

	g "github.com/aldesgroup/goald"
)

// static, reflect-free access to the definition of the Contact specs
type contactSpecs struct {
	g.IBusinessObjectSpecs
	firstName *g.StringField
	lastName  *g.StringField
}

// this is the main way to refer to the Contact specs in the applicative code
func Contact() *contactSpecs {
	return contact
}

// internal variables
var (
	contact     *contactSpecs
	contactOnce sync.Once
)

// fully describing each of this class' properties & relationships
func newContactSpecs() *contactSpecs {
	newSpecs := &contactSpecs{IBusinessObjectSpecs: g.NewBusinessObjectSpecs()}
	newSpecs.firstName = g.NewStringField(newSpecs, "FirstName", false)
	newSpecs.lastName = g.NewStringField(newSpecs, "LastName", false)

	return newSpecs
}

// making sure the Contact specs exists at app startup
func init() {
	contactOnce.Do(func() {
		contact = newContactSpecs()
	})

	// this helps dynamically access to the Contact specs
	g.RegisterSpecs("Contact", contact)
}

// accessing all the Contact class' properties and relationships

func (c *contactSpecs) FirstName() *g.StringField {
	return c.firstName
}

func (c *contactSpecs) LastName() *g.StringField {
	return c.lastName
}
