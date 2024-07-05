package test4stopper
import (
    p96f153e3d "github.com/starter-go/stopper/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type p96f153e3d.DemoUnit in package:github.com/starter-go/stopper/src/test/golang/unit
//
// id:com-96f153e3d8fbff2b-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p96f153e3d8_unit_DemoUnit struct {
}

func (inst* p96f153e3d8_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-96f153e3d8fbff2b-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p96f153e3d8_unit_DemoUnit) new() any {
    return &p96f153e3d.DemoUnit{}
}

func (inst* p96f153e3d8_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p96f153e3d.DemoUnit)
	nop(ie, com)

	


    return nil
}


