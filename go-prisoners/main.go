package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Prisoners\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("./prisoners <boxes> <prisoners> <runs> <random|sequence>\n")
	fmt.Printf("e.g.\n")
	fmt.Printf("For 10 boxes, 10 prisoners with random box choice repeated 1000 times\n")
	fmt.Printf("./prisoners 10 10 1000 random\n")
	fmt.Printf("For 10 boxes, 10 prisoners with sequence box choice repeated 1000 times\n")
	fmt.Printf("./prisoners 10 10 1000 sequence\n")
	fmt.Printf("\n")
	if len(os.Args) < 5 {
		return
	}

	experimentFn := RandomChoice
	if os.Args[4] == "random" {
		experimentFn = RandomChoice
	} else {
		experimentFn = SequenceChoice
	}

	rand.Seed(time.Now().UnixNano())
	prisoners, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Print(err)
		return
	}
	boxes, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Print(err)
		return
	}
	runs, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		fmt.Print(err)
		return
	}

	live := 0
	die := 0
	for r := 0; r < int(runs); r++ {
		if RunExperiment(int(boxes), int(prisoners), experimentFn) {
			live++
		} else {
			die++
		}
	}
	log("%v%% escapes", (100*float64(live))/(float64(live)+float64(die)))
}
