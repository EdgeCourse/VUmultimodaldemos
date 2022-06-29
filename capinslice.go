/*
In Go (golang) what happens to memory when you append to a slice?

If there’s enough space in the slice’s backing array, the element just gets added. If there’s not enough space, a new array is allocated, all the items are copied over, and the new item is added at the end. The interesting part is allocating that new array. And here’s the answer:

Go slices grow by doubling until size 1024, after which they grow by 25% each time

This is an implementation detail and may change. The above is correct for Go 1.1 and 1.2.

*/

package main

import "fmt"

func main() {
	var x []int // Same as x := make([]int, 0)
	for i := 0; i < 100; i++ {
		fmt.Printf("%d: %p cap %d\n", i, x, cap(x))
		x = append(x, i)
	}
}

/*
When you print %p of a slice, it prints the address of it’s backing array (see here in reflect.Pointer and SliceHeader type). That allows us to see what’s happening with that array:

The slice starts out without a backing array (pointer is 0x0).
When you append the very first element, an array of size 1 gets allocated.
When you append a second element, an array of size 2 gets allocated, the first element copied over, and the second element added. Note how the pointer changes.
When you add the third element, another allocation and copy happen, this time to an array of size 4.
The fourth element is the first one that doesn’t need to grow the backing array, so the pointer doesn’t change there.
To see how this doubling changes at 1024, let’s only print when an allocation happens (i.e. cap changes):
*/
/*
package main

import (
    "fmt"
)

func main() {
    var oldcap int
    var x []int
    fmt.Printf("   0: %p cap %d\n", x, cap(x))
    x = append(x, 0)
    for i := 1; i < 2048; i++ {
        if oldcap != cap(x) {
            fmt.Printf("%4d: %p cap %d->%d\n", i, x, oldcap, cap(x))
            oldcap = cap(x)
        }
        x = append(x, i)
    }
}
*/
//At 1024 we grow to 1280 (1024 + 1024/4), not 2048. This all happens in pkg/runtime/slice.c, with similar code in reflect/value.go.

//What all this means to me is that if you know even vaguely how big your slice will get, always allocate with a capacity (i.e. make([]int, 0, 256)). That will save a lot of allocation, copying, and subsequent garbage collection.
