package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/stopper/modules/stopper"
)

func main() {
	m := stopper.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
