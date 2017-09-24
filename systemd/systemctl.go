package systemd

import (
	"fmt"
	"time"

	"github.com/mschenk42/gopack"
	"github.com/mschenk42/gopack/action"
	"github.com/mschenk42/gopack/task"
)

// SystemCtl wraps basic systemctl manager commands
type SystemCtl struct {
	Service string

	gopack.BaseTask
}

// Run initializes default property values and delegates to BaseTask RunActions method
func (s SystemCtl) Run(runActions ...action.Name) gopack.ActionRunStatus {
	s.setDefaults()
	return s.RunActions(&s, s.registerActions(), runActions)
}

func (s SystemCtl) registerActions() action.Funcs {
	return action.Funcs{
		action.Start:   s.start,
		action.Restart: s.restart,
		action.Stop:    s.stop,
		action.Enable:  s.enable,
		action.Disable: s.disable,
		action.Reload:  s.reload,
	}
}

func (s *SystemCtl) setDefaults() {
}

// String returns a string which identifies the task with it's property values
func (s SystemCtl) String() string {
	return fmt.Sprintf("systemctl %s", s.Service)
}

func (s SystemCtl) start() (bool, error) {
	t := task.Command{
		Name:   "systemctl",
		Args:   []string{"start", s.Service},
		Stream: true,
	}

	t.SetNotIf(func() (bool, error) {
		notif := task.Command{
			Name:     "systemctl",
			Args:     []string{"is-active", s.Service},
			Timeout:  time.Second * 10,
			BaseTask: gopack.BaseTask{ContOnError: true},
		}
		return notif.Run(action.Run)[action.Run], nil
	})

	return t.Run(action.Run)[action.Run], nil
}

func (s SystemCtl) restart() (bool, error) {
	t := task.Command{
		Name:   "systemctl",
		Args:   []string{"restart", s.Service},
		Stream: true,
	}
	return t.Run(action.Run)[action.Run], nil
}

func (s SystemCtl) stop() (bool, error) {
	t := task.Command{
		Name:   "systemctl",
		Args:   []string{"stop", s.Service},
		Stream: true,
	}

	t.SetOnlyIf(func() (bool, error) {
		onlyif := task.Command{
			Name:     "systemctl",
			Args:     []string{"is-active", s.Service},
			Timeout:  time.Second * 10,
			BaseTask: gopack.BaseTask{ContOnError: true},
		}
		return onlyif.Run(action.Run)[action.Run], nil
	})

	return t.Run(action.Run)[action.Run], nil
}

func (s SystemCtl) enable() (bool, error) {
	t := task.Command{
		Name:   "systemctl",
		Args:   []string{"enable", s.Service},
		Stream: true,
	}

	t.SetNotIf(func() (bool, error) {
		notif := task.Command{
			Name:     "systemctl",
			Args:     []string{"is-enabled", s.Service},
			Timeout:  time.Second * 10,
			BaseTask: gopack.BaseTask{ContOnError: true},
		}
		return notif.Run(action.Run)[action.Run], nil
	})

	return t.Run(action.Run)[action.Run], nil
}

func (s SystemCtl) disable() (bool, error) {
	t := task.Command{
		Name:   "systemctl",
		Args:   []string{"disable", s.Service},
		Stream: true,
	}

	t.SetOnlyIf(func() (bool, error) {
		onlyif := task.Command{
			Name:     "systemctl",
			Args:     []string{"is-enabled", s.Service},
			Timeout:  time.Second * 10,
			BaseTask: gopack.BaseTask{ContOnError: true},
		}
		return onlyif.Run(action.Run)[action.Run], nil
	})

	return t.Run(action.Run)[action.Run], nil
}

func (s SystemCtl) reload() (bool, error) {
	t := task.Command{
		Name:   "systemctl",
		Args:   []string{"reload", s.Service},
		Stream: true,
	}
	return t.Run(action.Run)[action.Run], nil
}
