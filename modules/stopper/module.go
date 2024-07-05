package stopper

import (
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/stopper"
	"github.com/starter-go/stopper/gen/main4stopper"
	"github.com/starter-go/stopper/gen/test4stopper"
)

// Module  ...
func Module() application.Module {
	mb := stopper.NewMainModule()
	mb.Components(main4stopper.ExportComponents)
	mb.Depend(starter.Module())
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := stopper.NewTestModule()
	mb.Components(test4stopper.ExportComponents)
	mb.Depend(Module())
	return mb.Create()
}
