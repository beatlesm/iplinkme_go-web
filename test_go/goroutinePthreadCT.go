package main
 
import (
	"fmt"
	"time"
	"math/rand"
)

func routine(name string, delay time.Duration) {
	t0 := time.Now()
	fmt.Println(name, " start at ", t0)

	time.Sleep(delay)

	t1 := time.Now()
	fmt.Println(name, " end at ", t1)

	fmt.Println(name, " lasted ", t1.Sub(t0))
}
func main() {
	//rand seed
	rand.Seed(time.Now().Unix())

	var name string
	for i := 0; i < 3; i++ {
		name = fmt.Sprintf("")
	}
}