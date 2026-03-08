// You can edit this code!
// Click here and start typing.
package playground

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var xxx int32
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&xxx, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(xxx)
}
