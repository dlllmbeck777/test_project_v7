package main

import (
	"fmt"
	"github.com/transferMVP/transfer.webapp/internal/transferauth"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println("main")
	transferauth.Run()

}
