package termenv

import (
	"os/exec"
	"runtime"

	"github.com/gookit/goutil/internal/checkfn"
)

var (
	cmdList  = []string{"cmd", "cmd.exe"}
	pwshList = []string{"powershell", "powershell.exe", "pwsh", "pwsh.exe"}
)

func shellExec(expr, shell string) (string, error) {
	// "-c" for bash,sh,zsh shell
	mark := "-c"

	// special for Windows shell
	if runtime.GOOS == "windows" {
		// use cmd.exe, mark is "/c"
		if checkfn.StringsContains(cmdList, shell) {
			mark = "/c"
		} else if checkfn.StringsContains(pwshList, shell) {
			// "-Command" for powershell
			mark = "-Command"
		}
	}
	
	cmd := exec.Command(shell, mark, expr)
	bs, err := cmd.CombinedOutput()
	return string(bs), err
}
