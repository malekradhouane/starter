package cli

import (
	"fmt"

	"storeinit/infra/controllers"
)

// Sentinel error. Do not use alone.
var ErrController = fmt.Errorf("%w: CLI", controllers.ErrController)
