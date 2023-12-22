``` go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Below is an example of using our PrintMemUsage() function
	// Print our starting memory usage (should be around 0mb)
	PrintMemUsage()
	go func() {}()
	fmt.Println("<<declare a goroutine>>")
	PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v KiB", bTokb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v kiB", bTokb(m.TotalAlloc))
	fmt.Printf("\tSys = %v KiB", bTokb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bTokb(b uint64) uint64 {
	return b / 1024
}

```