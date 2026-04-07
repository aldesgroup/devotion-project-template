// Generated file, do not edit!
package model

import (
	"sync"

	g "github.com/aldesgroup/goald"
)

// static, reflect-free access to the definition of the Contact model
type contactModel struct {
	g.IBusinessObjectModel
	firstName *g.StringField
	lastName  *g.StringField
}

// this is the main way to refer to the Contact model in the applicative code
func Contact() *contactModel {
	return contact
}

// internal variables
var (
	contact     *contactModel
	contactOnce sync.Once
)

// fully describing each of this class' properties & relationships
func newContactModel() *contactModel {
	newModel := &contactModel{IBusinessObjectModel: g.NewBusinessObjectModel()}
	newModel.firstName = g.NewStringField(newModel, "FirstName", false)
	newModel.lastName = g.NewStringField(newModel, "LastName", false)

	return newModel
}

// making sure the Contact model exists at app startup
func init() {
	contactOnce.Do(func() {
		contact = newContactModel()
	})

	// this helps dynamically access to the Contact model
	g.RegisterModel("Contact", contact)
}

// accessing all the Contact class' properties and relationships

func (c *contactModel) FirstName() *g.StringField {
	return c.firstName
}

func (c *contactModel) LastName() *g.StringField {
	return c.lastName
}
