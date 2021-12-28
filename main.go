package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type uselessFunc func()

var vtp = []uselessFunc{
	download,
	mvs,
	bootVax,
	c64,
	adminports,
	fsck,
	bootMsDos,
	stateLinux,
	qradarAutoupdate,
	pfSenseMenu,
	vm370,
	ciscoboot,
	curlDownload,
}

var clo = []uselessFunc{
	standard,
}

func main() {
	args := os.Args[1:]
	rand.Seed(time.Now().UnixNano())
	calls := rand.Intn(len(vtp))

	if len(args) > 0 {
		var err error
		calls = 0
		calls, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Print("ok, you tried to pass the number of how many program execute...\n")
			fmt.Printf("but what you type is not a number [%v]\n", args[0])
			fmt.Printf("so you will have the default for now :)\n")
		}
	}
	calls = int(math.Min(float64(calls), float64(len(vtp))))
	if calls == 0 {
		calls = 1
	}

	fmt.Printf("here we go!  [%v]\n\n\n\n\n", calls)

	rand.Shuffle(len(vtp), func(i, j int) { vtp[i], vtp[j] = vtp[j], vtp[i] })

	vtp = vtp[:calls]

	for _, call := range vtp {
		call()
	}

	call := clo[rand.Intn(len(clo))]
	call()
}
