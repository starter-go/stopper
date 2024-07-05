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

	sfile *stopperFile
}

func (inst *StopperServiceImpl) _impl() stopper.Service {
	return inst
}

// Life ...
func (inst *StopperServiceImpl) Life() *application.Life {

	if !inst.Enabled {
		return &application.Life{}
	}

	return &application.Life{
		OnCreate: inst.onInit,
		OnStart:  inst.onStart,
		OnStop:   inst.onStop,
		OnLoop:   inst.onLoop,
	}
}

func (inst *StopperServiceImpl) onInit() error {
	ctx := &stopperContext{
		ac: inst.AppContext,
		fs: inst.FS,
	}
	man := getStopperFileManager(ctx)
	sfile := man.new()
	inst.sfile = sfile
	return nil
}

func (inst *StopperServiceImpl) onStart() error {
	err := inst.sfile.create()
	if err == nil {
		path := inst.sfile.file
		vlog.Info("make stopper listener file @(%s)", path.GetPath())
	} else {
		vlog.Warn(err.Error())
	}
	return nil
}

func (inst *StopperServiceImpl) onLoop() error {
	sfile := inst.sfile
	for {
		if !sfile.exists() {
			break
		}
		time.Sleep(time.Second * 2)
	}
	vlog.Info("StopperServiceImpl: stopping ...")
	return nil
}

func (inst *StopperServiceImpl) onStop() error {
	sfile := inst.sfile
	if sfile.exists() {
		err := sfile.remove()
		vlog.Warn(err.Error())
	}
	return nil
}

// Stop ...
func (inst *StopperServiceImpl) Stop(c context.Context, scope stopper.Scope) error {

	sfile := inst.sfile
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
	case stopper.ScopeLatest:
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
