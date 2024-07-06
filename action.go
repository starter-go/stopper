package stopper

import (
	"fmt"
	"strings"

	"github.com/starter-go/application"
)

// Action ...  表示传入的动作指令
type Action string

// 定义几个传入的动作指令
const (
	ActionNone    Action = "none"
	ActionStart   Action = "start"
	ActionStop    Action = "stop"
	ActionRestart Action = "restart"
)

// ParseAction ...
func ParseAction(text string) (Action, error) {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	a1 := Action(text)
	switch a1 {
	case ActionStart:
		return a1, nil
	case ActionStop:
		return a1, nil
	case ActionRestart:
		return a1, nil
	case ActionNone:
		return a1, nil
	default:
		return ActionNone, fmt.Errorf("bad action name: '%s'", text)
	}
}

// GetAction 从上下文中取操作码
func GetAction(ac application.Context) Action {

	action := ActionNone

	// from args
	available := map[Action]bool{
		ActionStart:   true,
		ActionStop:    true,
		ActionRestart: true,
	}
	args := ac.GetArguments().Raw()
	for _, a1 := range args {
		a2 := Action(a1)
		if available[a2] {
			action = a2
			break
		}
	}

	if action != ActionNone {
		return action
	}

	// from properties
	const (
		name = "starter.stopper.action"
	)
	props := ac.GetProperties()
	value := props.GetProperty(name)
	action, _ = ParseAction(value)
	return action
}
