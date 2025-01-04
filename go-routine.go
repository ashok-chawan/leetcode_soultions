package main

import (
	"time"
)

// func cpuIntensive(p *int64) {
// 	for i := int64(1); i <= 10000000; i++ {
// 		*p = i
// 	}
// 	fmt.Println("Done intensive thing")
// }

// func printVar(p *int64) {
// 	fmt.Printf("print x = %d.\n", *p)
// }

// // func main() {
// // 	fmt.Println("Main function")
// // 	runtime.GOMAXPROCS(1)

// // 	x := int64(0)
// // 	go cpuIntensive(&x) // This should go into background
// // 	go printVar(&x)     // This won't get scheduled until everything has finished.

// // 	time.Sleep(1 * time.Nanosecond) // Wait for goroutines to finish (has Gosched in code)
// // }
var global int

func testfunc() {
	global++
}
func main() {
	for i := 0; i < 100000; i++ {
		go testfunc()
	}
	time.Sleep(1 * time.Second)
}
