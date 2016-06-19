package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("started the program")
	defer fmt.Println("ended the program")
	tim := NewTimer(250, 3000)
	tim.Tic()
}

type timer struct {
	tic     uint
	timeout uint
	t0      uint
	counts  uint
}

func NewTimer(tic, timeout uint) *timer {
	deftic := uint(500)
	defboom := uint(3000)
	return &timer{tic: max(tic, deftic), timeout: max(timeout, defboom)}

}
func max(a, b uint) uint {
	if a >= b {
		return a
	}
	return b
}

func (t *timer) Tic() {
	boom := time.After(time.Duration(t.timeout) * time.Millisecond)
	tick := time.Tick(time.Duration(t.tic) * time.Millisecond)
	for {
		select {
		case <-boom:
			fmt.Println("BOOM!")
			fmt.Printf("%s", t.TimeString())
			return
		case <-tick:
			fmt.Println("tick.")
			t.counts++
			//default:
			//fmt.Println("    .")
			//time.Sleep(50 * time.Millisecond)
		}
	}
}
func (t *timer) Time() float64{
	return float64((t.counts*t.tic)/uint(1000))
}
func (t *timer) TimeString() string{
	return fmt.Sprintf("%.1fs passed by\n",float64((t.counts*t.tic)/uint(1000)))
}
