package lib

import (
	"os/exec"
	"runtime"
)

// AutoStartBrowser a
func AutoStartBrowser(url string) {
	os := runtime.GOOS

	switch os {
	case "windows":
		exec.Command("cmd", "/C", "start", url).Run()
	case "linux":
		exec.Command("xdg-open", url).Run()
	case "darwin":
		exec.Command("open", url).Run()
	}
}
