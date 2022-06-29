//more detail:
//GODEBUG=gctrace=1 go run memstats.go

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var ms runtime.MemStats
	printMemStat(ms)

	//----------------------------------
	// you can write any code here
	//----------------------------------
	intArr := make([]int, 900000)
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Int()
	}
	//------------------------------------
	time.Sleep(5 * time.Second)

	printMemStat(ms)

}

func printMemStat(ms runtime.MemStats) {
	runtime.ReadMemStats(&ms)
	fmt.Println("--------------------------------------")
	fmt.Println("Memory Statistics Reporting time: ", time.Now())
	fmt.Println("--------------------------------------")
	fmt.Println("Bytes of allocated heap objects: ", ms.Alloc)
	fmt.Println("Total bytes of Heap object: ", ms.TotalAlloc)
	fmt.Println("Bytes of memory obtained from OS: ", ms.Sys)
	fmt.Println("Count of heap objects: ", ms.Mallocs)
	fmt.Println("Count of heap objects freed: ", ms.Frees)
	fmt.Println("Count of live heap objects", ms.Mallocs-ms.Frees)
	fmt.Println("Number of completed GC cycles: ", ms.NumGC)
	fmt.Println("--------------------------------------")
}

/*
detailed one has output like:
gc 9 @0.126s 2%: 0.10+0.73+0.012 ms clock, 0.40+0.48/0.68/0.060+0.048 ms cpu, 4->4->0 MB, 5 MB goal, 4 P
…
Here, in the numbers 4->4->0, the first number represents heap size prior to garbage collection, the second number represents heap size after garbage collection, and the last number is the live heap object count.

*/
