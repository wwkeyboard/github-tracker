// You can edit this code!
// Click here and start typing.
package main

import (
	"flag"
	"fmt"
)

func main() {
	conf, _ := parseArgs()
	fmt.Printf("Hello, %v, %v, %v\n", conf.Str, conf.Num, conf.Flg)
}

type Config struct {
	Str string
	Num int
	Flg bool
}

func parseArgs() (Config, error) {
	// flag
	strPtr := flag.String("word", "foo", "a string")
	intPtr := flag.Int("number", 56, "an int")
	flgPtr := flag.Bool("work", false, "a bool")

	flag.Parse()

	return Config{
		Str: *strPtr,
		Num: *intPtr,
		Flg: *flgPtr,
	}, nil
}
