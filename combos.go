package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pdk/combos/comp"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprint(os.Stderr, "usage: combos n a b c ...\n")
		fmt.Fprint(os.Stderr, "prints all combos up to length n of a, b, c, ...\n")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("%s", err)
	}

	things := os.Args[2:]

	result := comp.Combos(n, things)
	for _, each := range result {
		fmt.Printf("%s\n", each)
	}
}
