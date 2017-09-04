package systemd

import (
	"fmt"

	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
)

// Daemon reloads/reexecutes systemd manager
type Daemon struct {
	gopack.BaseTask
}

// Run initializes default property values and delegates to BaseTask RunActions method
func (d Daemon) Run(runActions ...action.Enum) gopack.ActionRunStatus {
	d.setDefaults()
	return d.RunActions(&d, d.registerActions(), runActions)
}

func (d Daemon) registerActions() action.Methods {
	return action.Methods{
		action.Reload: d.reload,
	}
}

func (s *Daemon) setDefaults() {
}

// String returns a string which identifies the task with it's property values
func (s Daemon) String() string {
	return fmt.Sprintf("daemon")
}

func (s Daemon) reload() (bool, error) {
	return true, nil
}
