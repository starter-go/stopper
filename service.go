package stopper

import (
	"context"
)

// Service ...
type Service interface {
	Stop(c context.Context, scope Scope) error
}
