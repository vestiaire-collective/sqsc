package main

import (
	"fmt"
	"github.com/chaseisabelle/sqsc"
	"os"
)

func main() {
	bod := os.Args[1]

	cli, err := sqsc.New(&sqsc.Config{
		Region:   "us-east-1",
		URL:      "http://localhost:4100/queue/job",
		Endpoint: "http://127.0.0.1:4100",
	})

	res, err := cli.Produce(bod, 0)

	fmt.Printf("produce response: %+v\n", res)
	fmt.Printf("produce error:    %+v\n", err)

	res, rh, err := cli.Consume()

	fmt.Printf("consume response:       %+v\n", res)
	fmt.Printf("consume receipt handle: %+v\n", rh)
	fmt.Printf("consume error:          %+v\n", err)

	res, err = cli.Delete(rh)

	fmt.Printf("delete response: %+v\n", res)
	fmt.Printf("delete error:    %+v\n", err)
}
