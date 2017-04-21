package controller

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
		exec.Command("xdg-open", url)
	case "darwin":
		exec.Command("open", url)
	}
}
