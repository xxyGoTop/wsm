package util

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
)

func PrintEnv()  {
	envs := os.Environ()

	fmt.Println(color.GreenString("=== Runtime ==="))

	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("Go Os: %s\n", runtime.GOOS)
	fmt.Printf("Go Path: %s\n", runtime.GOARCH)

	fmt.Println(color.GreenString("=== Environmental Variable ==="))

	for _, e := range envs {
		fmt.Println(e)
	}
}