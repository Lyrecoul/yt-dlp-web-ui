//go:build !windows

package internal

import (
	"os/exec"
	"syscall"
)

func setProcessAttrs(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func killProcess(p *Process) error {
	if p.proc == nil {
		return nil
	}
	pgid, err := syscall.Getpgid(p.proc.Pid)
	if err != nil {
		return err
	}
	return syscall.Kill(-pgid, syscall.SIGTERM)
}
