package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {
	osName := runtime.GOOS
	if osName == "darwin" {
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
	} else if osName == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
}