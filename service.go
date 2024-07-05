package stopper

import (
	"context"
)

// Service ...
type Service interface {

	// 返回当前进程的 收到的 指令动作
	GetAction() Action

	Stop(c context.Context, scope Scope) error
}
