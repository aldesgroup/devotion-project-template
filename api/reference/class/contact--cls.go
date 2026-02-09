// Generated file, do not edit!
package class

import (
	"github.com/aldesgroup/goald"
"github.com/aldesgroup/devotion-project-template/reference"
)

type ContactClass struct {
	goald.IClassCore
}

func ClassForContact(srcPath, lastMod string) goald.IClass {
	return &ContactClass{IClassCore: goald.NewClassCore(srcPath, "Contact", lastMod)}
}

func (thisClass *ContactClass) NewObject() any {
	return &reference.Contact{}
}

func (thisClass *ContactClass) NewSlice() any {
	return []*reference.Contact{}
}
