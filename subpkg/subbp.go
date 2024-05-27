package subpkg

import (
	"fmt"
	"time"
)

func PrintTimeInUTC() {

	fmt.Printf("time.Now().UTC(): %v\n", time.Now().UTC())
}
