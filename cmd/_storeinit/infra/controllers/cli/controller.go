package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"storeinit/config"
	"storeinit/domain"

	"github.com/jedib0t/go-pretty/table"
)

/*
	A handler is a component that "Command & Query" the application.
	This handler permit to control the application
	from the command line interfaces (CLI) to perform the desired action.
*/

var (
	ErrDirty  = errors.New("dirty storage")
	ErrFailed = errors.New("failed storage")
)

type Controller struct {
	app         ApplicationDeps
	lockTimeout time.Duration
}

// Run handles user command line and execute command.
func (ctrl *Controller) Run(gitTag string, cman config.ManagerContract) error {
	const (
		MinimalLockTimeout = 1 * time.Second
		DefaultLockTimeout = 15 * time.Second // same as go-migrate
	)

	var (
		flagVersion = flag.Bool("v", false, "show version and exit")
		flagUp      = flag.Bool("up", false, "Apply all schemas")
		flagDown    = flag.Bool("down", false, "Remove all schemas")
		flagForce   = flag.Bool("force", false, "Force migration to configured version and clear dirty flag")
		flagStatus  = flag.Bool("status", false, "Report migration status")
		flagTimeout = flag.Float64("lock-timeout", DefaultLockTimeout.Seconds(),
			fmt.Sprintf("lock timeout in seconds (default: %.2fs)", DefaultLockTimeout.Seconds()))
	)

	flag.Parse()

	if *flagVersion {
		slog.Info(os.Args[0], "version", gitTag)
		return nil
	}

	ctrl.lockTimeout = max(MinimalLockTimeout, time.Duration(*flagTimeout*float64(time.Second)))

	switch {
	case *flagForce:
		return ctrl.migrateForce()

	case *flagUp:
		return ctrl.migrateUp()

	case *flagDown:
		return ctrl.migrateDown()

	case *flagStatus:
		return ctrl.status()
	}

	return nil
}

func (ctrl *Controller) Close(_ context.Context) error {
	return nil
}

func (ctrl *Controller) migrateUp() error {
	req := domain.MigrateRequest{Action: domain.ActionMigrateUp}

	ctx, cancel := context.WithTimeout(context.Background(), ctrl.lockTimeout)
	defer cancel()
	results := ctrl.app.Migrate(ctx, req)

	tbl := [][]any{{"#", "Repo", "Driver", "Updated", "Error"}}

	withError := false
	for i, result := range results {
		repo := result.Repo
		updated := ""
		if !result.NoChange {
			updated = "Yes"
		}
		if err := result.Err; err != nil {
			tbl = append(tbl, []any{i, repo.Name(), repo.Driver(), "", err.Error()})
			withError = true
		} else {
			tbl = append(tbl, []any{i, repo.Name(), repo.Driver(), updated, ""})
		}
	}

	ctrl.tableReport(tbl)

	if withError {
		return fmt.Errorf("%w: migration UP with error", ErrController)
	}

	return nil
}

func (ctrl *Controller) status() error {
	ctx, cancel := context.WithTimeout(context.Background(), ctrl.lockTimeout)
	defer cancel()
	results := ctrl.app.MigrationStatus(ctx)

	tbl := [][]any{{"#", "Repo", "Driver", "Dirty", "Version", "Error"}}

	oneDirty, oneErr := false, false
	for i, result := range results {
		repo := result.Repo
		if err := result.Err; err != nil {
			tbl = append(tbl, []any{i, repo.Name(), repo.Driver(), "", "", err.Error()})
			oneErr = true
		} else {
			dirty := ""
			if result.Dirty {
				dirty = "  X"
				oneDirty = true
			}
			if result.SchemaVersion < 0 {
				tbl = append(tbl, []any{i, repo.Name(), repo.Driver(), "", "NoInit", ""})
			} else {
				tbl = append(tbl, []any{i, repo.Name(), repo.Driver(), dirty, result.SchemaVersion, ""})
			}
		}
	}

	ctrl.tableReport(tbl)

	var err error
	if oneDirty {
		err = errors.Join(err, ErrDirty)
	}
	if oneErr {
		err = errors.Join(err, ErrFailed)
	}

	return err
}

func (ctrl *Controller) migrateDown() error {
	req := domain.MigrateRequest{Action: domain.ActionMigrateDown}

	ctx, cancel := context.WithTimeout(context.Background(), ctrl.lockTimeout)
	defer cancel()
	results := ctrl.app.Migrate(ctx, req)

	tbl := [][]any{{"#", "Repo", "Driver", "Updated", "Error"}}

	withError := false
	for i, result := range results {
		repo := result.Repo
		updated := ""
		if !result.NoChange {
			updated = "Yes"
		}

		if err := result.Err; err != nil {
			tbl = append(tbl, []any{i, repo.Name(), repo.Driver(), "", err.Error()})
			withError = true
		} else {
			tbl = append(tbl, []any{i, repo.Name(), repo.Driver(), updated, ""})
		}
	}

	ctrl.tableReport(tbl)

	if withError {
		return fmt.Errorf("%w: migration DOWN with error", ErrController)
	}

	return nil
}

func (ctrl *Controller) migrateForce() error {
	req := domain.MigrateRequest{Action: domain.ActionMigrateForce}

	ctx, cancel := context.WithTimeout(context.Background(), ctrl.lockTimeout)
	defer cancel()
	results := ctrl.app.Migrate(ctx, req)

	withError := false
	tbl := [][]any{{"#", "Repo", "Driver", "Updated", "Error"}}
	for i, result := range results {
		repo := result.Repo
		updated := ""
		if !result.NoChange {
			updated = "Yes"
		}
		if err := result.Err; err != nil {
			tbl = append(tbl, []any{i, repo.Name(), repo.Driver(), "", err.Error()})
			withError = true
		} else {
			tbl = append(tbl, []any{i, repo.Name(), repo.Driver(), updated, ""})
		}
	}

	ctrl.tableReport(tbl)

	if withError {
		return fmt.Errorf("%w: migration FORCE with error", ErrController)
	}

	return nil
}

func (ctrl *Controller) tableReport(rows [][]any) {
	tblReport := table.NewWriter()
	tblReport.SetOutputMirror(os.Stdout)

	for i := range rows {
		if i == 0 {
			tblReport.AppendHeader(rows[i])
		} else {
			tblReport.AppendRow(rows[i])
		}
	}

	tblReport.Render()
}
