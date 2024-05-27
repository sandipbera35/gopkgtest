package gopkgtest

import (
	"fmt"
	"time"
)

func PrintDateInUtc() {
	fmt.Printf("time.Now().UTC(): %v\n", time.Now().UTC())
}
