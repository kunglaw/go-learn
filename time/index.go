package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	elapsed := time.Since(start)
	log.Printf("Binomial took %dms", elapsed.Nanoseconds()/1000)

}
