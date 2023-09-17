package common

import (
	"fmt"
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func ExecShell(name string, args ...string) (output string, status_code int) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("cmd.Run() failed with %s\n", err)
		status_code = 1
		output = fmt.Sprint(err)
	} else {
		output = string(out)
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		status_code = ws.ExitStatus()
	}
	return
}
