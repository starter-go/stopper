package main4stopper
import (
    p0d2a11d16 "github.com/starter-go/afs"
    p0ef6f2938 "github.com/starter-go/application"
    p938006534 "github.com/starter-go/stopper/src/main/golang/lib"
     "github.com/starter-go/application"
)

// type p938006534.StopperServiceImpl in package:github.com/starter-go/stopper/src/main/golang/lib
//
// id:com-93800653451eab6a-lib-StopperServiceImpl
// class:
// alias:alias-fee945b1a371c0b5131cb6da550039d6-Service
// scope:singleton
//
type p9380065345_lib_StopperServiceImpl struct {
}

func (inst* p9380065345_lib_StopperServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-93800653451eab6a-lib-StopperServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-fee945b1a371c0b5131cb6da550039d6-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9380065345_lib_StopperServiceImpl) new() any {
    return &p938006534.StopperServiceImpl{}
}

func (inst* p9380065345_lib_StopperServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p938006534.StopperServiceImpl)
	nop(ie, com)

	
    com.AppContext = inst.getAppContext(ie)
    com.FS = inst.getFS(ie)
    com.Enabled = inst.getEnabled(ie)
    com.FlagScope = inst.getFlagScope(ie)
    com.FlagAction = inst.getFlagAction(ie)


    return nil
}


func (inst*p9380065345_lib_StopperServiceImpl) getAppContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p9380065345_lib_StopperServiceImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*p9380065345_lib_StopperServiceImpl) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${starter.stopper.enabled}")
}


func (inst*p9380065345_lib_StopperServiceImpl) getFlagScope(ie application.InjectionExt)string{
    return ie.GetString("${starter.stopper.scope}")
}


func (inst*p9380065345_lib_StopperServiceImpl) getFlagAction(ie application.InjectionExt)string{
    return ie.GetString("${starter.stopper.action}")
}


