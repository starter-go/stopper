package unit

import (
	"context"

	"github.com/starter-go/units"
)

// DemoUnit ... 单元测试示例
type DemoUnit struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

}

func (inst *DemoUnit) _impl() units.Unit { return inst }

// Units ...
func (inst *DemoUnit) ListRegistrations(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:     "test1",
		Enabled:  true,
		Priority: 0,
		Do:       inst.test1,
	})

	return list
}

// Units ...
func (inst *DemoUnit) test1(cc context.Context) error {
	return nil
}
