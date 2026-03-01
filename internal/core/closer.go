package core

import "context"

type Closer interface {
	Close(context.Context) error
}
