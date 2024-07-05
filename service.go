package stopper

import "context"

// Scope 表示 stop 的作用域
type Scope int

// 定义 stop 的作用域
const (
	ScopeThis   Scope = 1
	ScopeLatest Scope = 2
	ScopeOlder  Scope = 3
	ScopeAll    Scope = 4
)

// Service ...
type Service interface {
	Stop(c context.Context, scope Scope) error
}
