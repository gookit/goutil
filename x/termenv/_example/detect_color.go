package main

import (
	"fmt"
	"os"

	"github.com/gookit/goutil/x/termenv"
)

func main() {
	termenv.SetDebugMode(true)

	cl := termenv.DetectColorLevel()
	fmt.Printf("DetectColorLevel: %s\n", cl)

	fmt.Printf("TermColorLevel: %s\n", termenv.TermColorLevel())
	fmt.Printf("IsSupportColor: %v\n", termenv.IsSupportColor())
	fmt.Printf("IsSupport256Color: %v\n", termenv.IsSupport256Color())
	fmt.Printf("IsSupportTrueColor: %v\n", termenv.IsSupportTrueColor())

	fmt.Println("Examples")
	fmt.Println("\x1b[34mHello \x1b[35mWorld\x1b[0m!")
	_, _ = fmt.Fprintf(os.Stdout, "\x1b[34mHello \x1b[35mWorld\x1b[0m!\n")
}
