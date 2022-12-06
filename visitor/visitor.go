package visitor

import (
	"errors"
	"fmt"
)

type ActivityType string

const (
	AddActivity      = ActivityType("join")
	SubtractActivity = ActivityType("left")
)

var negativeNumberMsg = errors.New("negative number is not accepted")

type Visitor struct {
	total, active int
}

func New() Visitor {
	return Visitor{
		total:  0,
		active: 0,
	}
}

// value based receiver
func (v *Visitor) Join() {
	v.active += 1
	v.total += 1
}

func (v *Visitor) Left() {
	v.active -= 1
}

func (v *Visitor) ActiveString() string {
	return fmt.Sprintf("Currently there are %d active user(s)", v.active)
}

func (v *Visitor) TotalString() string {
	return fmt.Sprintf("Total user until now: %d", v.total)
}

func (v *Visitor) join(value int) {
	defer func() {
		fmt.Println("closing the connection...")
		if p := recover(); p == negativeNumberMsg {
			fmt.Printf("panic happened but it's OK, panic message: %s\n", p)
		} else if p != nil {
			panic(fmt.Sprintf("panic happened but I can't recover this time, %s\n", p))
		}
	}()

	fmt.Println("new joiner...")

	if value < 0 {
		panic(negativeNumberMsg)
	}

	if value == 0 {
		panic("zero is not accepted")
	}
	// open connection to DB
	// write something to DB
	// close the connection
	v.active += value
	v.total += value
}

func (v *Visitor) left(value int) {
	if value < 0 {
		panic(negativeNumberMsg)
	}
	v.active -= value
}

func (v *Visitor) BulkActivity(ac string) func(int) {
	switch ac {
	case string(AddActivity):
		return v.join
	case string(SubtractActivity):
		return v.left
	default:
		return func(int) {
			fmt.Println("activity not supported")
		}

	}

}
