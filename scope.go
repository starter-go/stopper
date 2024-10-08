package stopper

import (
	"fmt"
	"strings"

	"github.com/starter-go/application"
)

// Scope 表示 stop 的作用域
type Scope int

// 定义 stop 的作用域
const (
	ScopeNone  Scope = 0
	ScopeThis  Scope = 1
	ScopeNewer Scope = 2
	ScopeOlder Scope = 3
	ScopeAll   Scope = 4
)

// ParseScope ...
func ParseScope(text string) (Scope, error) {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	switch text {
	case "all":
		return ScopeAll, nil
	case "newer":
		return ScopeNewer, nil
	case "older":
		return ScopeOlder, nil
	case "this":
		return ScopeThis, nil
	case "none":
		return ScopeNone, nil
	default:
		return ScopeNone, fmt.Errorf("bad scope name: '%s'", text)
	}
}

// GetScope  从上下文中取操作作用域
func GetScope(ac application.Context) Scope {
	const (
		name = "starter.stopper.scope"
	)
	props := ac.GetProperties()
	value := props.GetProperty(name)
	scope, _ := ParseScope(value)
	return scope
}
