package gopkgtest

import (
	"fmt"

	"github.com/sandipbera35/gopkgtest/subpkg"
)

func PrintName(name string) {

	fmt.Println("Name : ", name)
}

func PrintTimeInUTC() {
	subpkg.PrintTimeInUTC()
}
