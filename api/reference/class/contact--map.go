// Generated file, do not edit!
package class

import (
	"github.com/aldesgroup/corego"
	"github.com/aldesgroup/devotion-project-template/reference"
	"github.com/aldesgroup/goald"
)

// getting a property's value as a string, without using reflection
func (thisClass *ContactClass) GetValueAsString(bo goald.IBusinessObject, propertyName string) string {
	switch propertyName {
	case "FirstName":
		return bo.(*reference.Contact).FirstName
	case "ID":
		return core.Int64ToString(int64(bo.(*reference.Contact).ID))
	case "LastName":
		return bo.(*reference.Contact).LastName
	default:
		return "unknown property: " + propertyName
	}
}

// setting a property's value with a given string value, without using reflection
func (thisClass *ContactClass) SetValueAsString(bo goald.IBusinessObject, propertyName string, valueAsString string) error {
	switch propertyName {
	case "FirstName":
		bo.(*reference.Contact).FirstName = valueAsString
	case "ID":
		bo.(*reference.Contact).ID = goald.BObjID(core.StringToInt64(valueAsString, "(*reference.Contact).ID"))
	case "LastName":
		bo.(*reference.Contact).LastName = valueAsString
	}

	return goald.Error("Unknown property: %T.%s", bo, propertyName)
}
