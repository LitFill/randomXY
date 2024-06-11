// randomXY, LitFill <author at email dot com>
// program for ...
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Address struct {
	X, Y int
}

func (a Address) Get() (int, int) { return a.X, a.Y }
func (a Address) String() string  { return fmt.Sprintf("addr: %d ⊗ %d", a.X, a.Y) }

type Addrs []Address

func newAddress(rows, cols int) Addrs {
	addrs := make(Addrs, 0)
	for i := range rows {
		for j := range cols {
			addrs = append(addrs, Address{i, j})
		}
	}
	return addrs
}

func (as *Addrs) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(*as), func(i, j int) {
		(*as)[i], (*as)[j] = (*as)[j], (*as)[i]
	})
}

func (as *Addrs) Get() (Address, bool) {
	l := len(*as)
	if l == 0 {
		return Address{}, false
	}
	addr := (*as)[0]
	*as = (*as)[1:]
	return addr, true
}

func main() {
	rows := flag.Int("r", 18, "[int] for rows")
	cols := flag.Int("c", 8, "[int] for cols")
	flag.Parse()
	// rows, cols := 18, 8
	// addrses := newAddress(rows, cols)
	addrses := newAddress(*rows, *cols)
	addrses.Shuffle()
	for {
		a, got := addrses.Get()
		if !got {
			break
		}
		fmt.Printf("%s\n", a)
	}

}
