package stopper

import (
	"fmt"
	"strings"
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
