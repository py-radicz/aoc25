/* Package dayone */
package dayone

import (
	"bytes"
	"log"

	"github.com/py-radicz/aoc25/utils"
)

type Rotation struct {
	Dir byte
	Num int
}

type Lock struct {
	dial        int
	transitions int
}

func (l *Lock) Right(x int) {
	for range x {
		if l.dial == 99 {
			l.dial = 0
		} else {
			l.dial++
		}

		if l.DialedZero() {
			l.transitions++
		}

	}
}

func (l *Lock) Left(x int) {
	for range x {
		if l.dial == 0 {
			l.dial = 99
		} else {
			l.dial--
		}

		if l.DialedZero() {
			l.transitions++
		}
	}
}

func (l *Lock) DialedZero() bool {
	return l.dial == 0
}

func NewLock(start int) *Lock {
	return &Lock{dial: start}
}

func Rotations(in []byte) (rots []Rotation) {
	//input := [][]byte{
	//	[]byte("L68"),
	//	[]byte("L30"),
	//	[]byte("R48"),
	//	[]byte("L5"),
	//	[]byte("R60"),
	//	[]byte("L55"),
	//	[]byte("L1"),
	//	[]byte("L99"),
	//	[]byte("R14"),
	//	[]byte("L82"),
	//}
	input := bytes.Fields(in)

	for _, rot := range input {
		num := utils.AtoiBytes(rot[1:])
		rots = append(rots, Rotation{Dir: rot[0], Num: num})
	}
	return
}

func DayOne(day int) (partOne, partTwo int) {
	lock := NewLock(50)
	in, err := utils.GetInput(day)
	if err != nil {
		log.Fatal(err)
	}
	for _, rot := range Rotations(in) {
		switch rot.Dir {
		case 'R':
			lock.Right(rot.Num)
		case 'L':
			lock.Left(rot.Num)
		}

		if lock.DialedZero() {
			partOne++
		}
	}

	partTwo = lock.transitions
	return
}
