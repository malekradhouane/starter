package pg

import (
	"fmt"

	"storeinit/infra/repo"
)

var ErrRepo = fmt.Errorf("%w: PostgreSQL", repo.ErrRepo)
