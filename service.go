package stopper

import (
	"context"
)

// Service ...
type Service interface {

	// 返回当前进程的收到的指令动作
	GetAction() Action

	// 返回当前进程的收到的指令作用域
	GetScope() Scope

	Stop(c context.Context, scope Scope) error
}
