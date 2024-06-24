// randomXY, LitFill <marrazzy54 at gmail dot com>
// program for a cli returning a random combination of 2 ints for some purposes.
package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var namaHari = []string{"Sen", "Sel", "Rab", "Kam", "Jum", "Sab"}

type Address struct {
	X, Y int
}

func (a Address) Get() (int, int) { return a.X, a.Y }

// func (a Address) String() string  { return fmt.Sprintf("addr: %d ⊗ %d", a.X, a.Y) }
func (a Address) String() string {
	hari := namaHari[a.X%6]
	jam := a.X/6 + 1
	kelas := rune('A' + a.Y)
	return fmt.Sprintf("%s %d %c", hari, jam, kelas)
}

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
	a := *as
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
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
