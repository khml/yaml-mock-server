package routing

import (
	"os/exec"
	"runtime"
)

func openBrowser(setting *Setting) error {
	var url = "http://localhost:" + setting.Config.Port + setting.Config.Browser.OpenPath

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	case "linux":
		cmd = "xdg-open"
	default:
		return nil
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
