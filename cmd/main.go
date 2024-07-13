package main

import (
	"github.com/th-yoo/vlcode"
	"fmt"
	"os"
	"strconv"
	"strings"
	"encoding/hex"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: $ vlcode <uint_value>\nex) $vlcode 0x10\n")
		os.Exit(1)
	}

	arg := os.Args[1]
	if !strings.HasPrefix(arg, "0x") {
		fmt.Fprintf(os.Stderr, "Error: Argument must start with '0x'\n")
		os.Exit(2)
	}

	value, err := strconv.ParseUint(arg[2:], 16, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(3)
	}

	b := vlcode.Encode(uint(value))
	fmt.Println(hex.Dump(b))

	val, _ := vlcode.Decode(b)
	fmt.Printf("0x%X\n", val)
}
