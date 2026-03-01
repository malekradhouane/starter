package store

import (
	"fmt"
	"sync"

	"github.com/malekradhouane/trippy/store/postgres"
	"github.com/malekradhouane/trippy/store/types"
)

type Options struct{}

var (
	createPostgresStoresOnce sync.Once
	stores                   StoreSet
)

// StoreSet holds instances of the concrete type implementing the data stores
type StoreSet struct {
	User types.UserStore
}

func CreatePostgresStores(opts *Options) error {
	var err error

	createPostgresStoresOnce.Do(func() {
		if opts == nil {
			opts = &Options{}
		}
		stores.User, err = postgres.NewUserStore()
		if err != nil {
			err = fmt.Errorf("CreateStores : NewUserStore err: %w", err)
			return
		}
	})

	return err
}

func Users() types.UserStore {
	return stores.User
}
