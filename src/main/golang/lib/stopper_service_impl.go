package lib

import (
	"context"
	"time"

	"github.com/starter-go/afs"
	"github.com/starter-go/application"
	"github.com/starter-go/stopper"
	"github.com/starter-go/vlog"
)

// StopperServiceImpl ...
type StopperServiceImpl struct {

	//starter:component

	_as func(stopper.Service) //starter:as("#")

	AppContext application.Context //starter:inject("context")
	FS         afs.FS              //starter:inject("#")

	Enabled bool //starter:inject("${starter.stopper.enabled}")

	working *myWorking // 需要执行的任务
}

func (inst *StopperServiceImpl) _impl() stopper.Service {
	return inst
}

func (inst *StopperServiceImpl) getWorking() *myWorking {
	wk := inst.working
	if wk == nil {
		wk = &myWorking{}
		wk.init(inst)
		inst.working = wk
	}
	return wk
}

// Life ...
func (inst *StopperServiceImpl) Life() *application.Life {

	wk := inst.getWorking()

	if inst.Enabled {

		if wk.action == stopper.ActionStart {
			x := &myStarter{service: inst}
			return x.life()
		}

		if wk.action == stopper.ActionRestart {
			wk.stopAllOthers()
			time.Sleep(time.Second * 5)
			x := &myStarter{service: inst}
			return x.life()
		}

		if wk.action == stopper.ActionStop {
			x := &myStopper{service: inst}
			return x.life()
		}
	}

	return &application.Life{}
}

// // GetAction ...
// func (inst *StopperServiceImpl) GetAction() stopper.Action {
// 	wk := inst.getWorking()
// 	return wk.action
// }

// // GetScope ...
// func (inst *StopperServiceImpl) GetScope() stopper.Scope {
// 	wk := inst.getWorking()
// 	return wk.scope
// }

// Stop ...
func (inst *StopperServiceImpl) Stop(c context.Context, scope stopper.Scope) error {

	wk := inst.getWorking()
	sfile := wk.sfile
	man := sfile.manager
	todolist := make([]*stopperFile, 0)

	switch scope {
	case stopper.ScopeThis:
		todolist = append(todolist, sfile)
		break
	case stopper.ScopeOlder:
		older := man.getOlder()
		todolist = append(todolist, older)
		break
	case stopper.ScopeNewer:
		newer := man.getNewer()
		todolist = append(todolist, newer)
		break
	case stopper.ScopeAll:
		todolist = man.listAll()
		break
	}

	for _, item := range todolist {
		if item == nil {
			continue
		}
		if item.exists() {
			err := item.remove()
			vlog.Warn(err.Error())
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type myWorking struct {
	scope  stopper.Scope  // 请求的作用域
	action stopper.Action // 请求的 动作
	sfile  *stopperFile   // 本进程的停止控制文件
}

func (inst *myWorking) init(service *StopperServiceImpl) {
	if inst.sfile != nil {
		// 已经初始化了， 跳过这个步骤
		return
	}

	ac := service.AppContext
	inst.action = stopper.GetAction(ac)
	inst.scope = stopper.GetScope(ac)

	ctx := &stopperContext{
		ac: service.AppContext,
		fs: service.FS,
	}
	man := getStopperFileManager(ctx)
	sfile := man.new()
	inst.sfile = sfile
}

func (inst *myWorking) parseScope(text string) {
	scope, err := stopper.ParseScope(text)
	if err != nil {
		vlog.Warn(err.Error())
	}
	inst.scope = scope
}

func (inst *myWorking) parseAction(text string) {
	action, err := stopper.ParseAction(text)
	if err != nil {
		vlog.Warn(err.Error())
	}
	inst.action = action
}

func (inst *myWorking) stopAllOthers() {
	ignore := inst.sfile.file.GetName()
	all := inst.sfile.manager.listAll()
	for _, item := range all {
		name := item.file.GetName()
		if name == ignore {
			continue // skip
		}
		item.remove()
	}
}

////////////////////////////////////////////////////////////////////////////////

type myStarter struct {
	service *StopperServiceImpl
}

func (inst *myStarter) life() *application.Life {
	return &application.Life{
		OnCreate:   inst.onInit,
		OnStart:    inst.onStart,
		OnStop:     inst.onStop,
		OnLoop:     inst.onLoop,
		OnStopPost: inst.onStopped,
	}
}

func (inst *myStarter) onInit() error {
	return nil
}

func (inst *myStarter) onStart() error {
	wk := inst.service.getWorking()
	sfile := wk.sfile
	err := sfile.create()
	if err == nil {
		path := sfile.file
		vlog.Info("make stopper listener file @(%s)", path.GetPath())
	} else {
		vlog.Warn(err.Error())
	}
	return nil
}

func (inst *myStarter) onLoop() error {
	wk := inst.service.getWorking()
	sfile := wk.sfile
	for {
		if !sfile.exists() {
			break
		}
		time.Sleep(time.Second * 2)
	}
	vlog.Info("stopping ...")
	return nil
}

func (inst *myStarter) onStop() error {
	wk := inst.service.getWorking()
	sfile := wk.sfile
	if sfile.exists() {
		err := sfile.remove()
		vlog.Warn(err.Error())
	}
	return nil
}

func (inst *myStarter) onStopped() error {
	vlog.Info("stopped")
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type myStopper struct {
	service *StopperServiceImpl
}

func (inst *myStopper) life() *application.Life {
	return &application.Life{
		OnStart: inst.doShutdown,

		// OnCreate: inst.onInit,
		// OnStop:   inst.onStop,
		// OnLoop:   inst.onLoop,
	}
}

func (inst *myStopper) doShutdown() error {
	wk := inst.service.working
	scope := wk.scope
	ctx := context.Background()
	return inst.service.Stop(ctx, scope)
}

////////////////////////////////////////////////////////////////////////////////
