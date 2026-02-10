//go:build windows

package internal

import (
	"os/exec"
)

func setProcessAttrs(cmd *exec.Cmd) {
	// Windows 不支持 Setpgid，留空即可
}

func killProcess(p *Process) error {
	if p.proc == nil {
		return nil
	}
	return p.proc.Kill()
}
