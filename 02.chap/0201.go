package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "Format String %d\n", 10)
	fmt.Fprintf(os.Stdout, "Format String %s\n", "test")
	fmt.Fprintf(os.Stdout, "Format String %f\n", 10.01)

}
