package unit

import "github.com/starter-go/units"

// DemoUnit ... 单元测试示例
type DemoUnit struct {

	//starter:component
	_as func(units.Units) //starter:as(".")

}

func (inst *DemoUnit) _impl() units.Units { return inst }

// Units ...
func (inst *DemoUnit) Units(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Name:     "test1",
		Enabled:  true,
		Priority: 0,
		Test:     inst.test1,
	})

	return list
}

// Units ...
func (inst *DemoUnit) test1() error {
	return nil
}
