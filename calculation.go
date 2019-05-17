package main

import (
	"fmt"
	"strconv"
)

const (
	CHARGE = iota
	IDLE
)

var total float64

var step float64

func main() {

	total = 13.0 * 32.0

	step = 1.0

	idles := GenerateCardsetsIdle()
	for _, i := range idles {
		fmt.Println("============ ", i.a.val, " ", i.b.val)
		generateDataSetBySet(i.a.val, i.b.val, 3, 13)
	}

	chargers := GenerateCardsetsIdle()
	// calculate probability
	var f float64
	fmt.Println("< calculation start >")
	for _, i := range idles {
		for _, j := range chargers {

			f += calculatePP(i, j)
		}
	}
	fmt.Println(f)
}

func totalMinus(f float64, i int) float64 {
	f = f - (1.0)*float64(i)
	return f
}

type card struct {
	// val : 1 ~ 13
	val int
}

type Cardset struct {
	typ int
	a   *card
	b   *card

	and *card
}

func (cs *Cardset) compare(in *Cardset) string {
	volumeA := 0
	if cs.and != nil {
		volumeA = (cs.a.val + cs.b.val + cs.and.val) % 10
	} else {
		volumeA = (cs.a.val + cs.b.val) % 10
	}
	volumeB := 0
	if in.and != nil {
		volumeB = (in.a.val + in.b.val + in.and.val) % 10
	} else {
		volumeB = (in.a.val + in.b.val) % 10
	}
	if volumeA < volumeB {
		return "le"
	}
	if volumeA == volumeB {
		return "eq"
	}
	if volumeA > volumeB {
		return "la"
	}
	return ""
}

func checkIdle(e *Cardset) {
	if e.typ != IDLE {
		panic("wrong idle")
	}
}

func checkCharge(e *Cardset) {
	if e.typ != CHARGE {
		panic("wrong idle")
	}
}

func (cs *Cardset) andCalculate(e *Cardset) (res []*Cardset) {
	if e == nil {
		panic("fail input e")
	}
	if cs.and != nil {
		res = append(res, cs)
		return
	}
	// main
	if cs.typ == CHARGE {
		if sum(e.a.getVal(), e.b.getVal()) == 8 || sum(e.a.getVal(), e.b.getVal()) == 9 {
			res = append(res, cs)
			return
		}
		// <3
		if sum(cs.a.getVal(), cs.b.getVal()) < 3 {
			for i := 1; i <= 13; i++ {
				res = append(res, &Cardset{
					typ: cs.typ,
					a:   cs.a,
					b:   cs.b,
					and: &card{
						val: i,
					},
				})
			}
			return
		}
		// 3
		if sum(cs.a.getVal(), cs.b.getVal()) == 3 {
			checkIdle(e)
			if e.and == nil || e.and.getVal() != 8 {
				for i := 1; i <= 13; i++ {
					res = append(res, &Cardset{
						typ: cs.typ,
						a:   cs.a,
						b:   cs.b,
						and: &card{
							val: i,
						},
					})
				}
				return
			}
		}
		// 4
		if sum(cs.a.getVal(), cs.b.getVal()) == 4 {
			checkIdle(e)
			if e.and == nil ||
				!(e.and.getVal() == 8 || e.and.getVal() == 0 || e.and.getVal() == 1 || e.and.getVal() == 9) {
				for i := 1; i <= 13; i++ {
					res = append(res, &Cardset{
						typ: cs.typ,
						a:   cs.a,
						b:   cs.b,
						and: &card{
							val: i,
						},
					})
				}
				return
			}
		}
		// 5
		if sum(cs.a.getVal(), cs.b.getVal()) == 5 {
			checkIdle(e)
			if e.and == nil ||
				!(e.and.getVal() == 0 || e.and.getVal() == 1 || e.and.getVal() == 2 || e.and.getVal() == 3 || e.and.getVal() == 8 || e.and.getVal() == 9) {
				for i := 1; i <= 13; i++ {
					res = append(res, &Cardset{
						typ: cs.typ,
						a:   cs.a,
						b:   cs.b,
						and: &card{
							val: i,
						},
					})
				}
				return
			}
		}
		// 6
		if sum(cs.a.getVal(), cs.b.getVal()) == 6 {
			checkIdle(e)
			if e.and != nil &&
				(e.and.getVal() == 6 || e.and.getVal() == 7) {
				for i := 1; i <= 13; i++ {
					res = append(res, &Cardset{
						typ: cs.typ,
						a:   cs.a,
						b:   cs.b,
						and: &card{
							val: i,
						},
					})
				}
				return
			}
		}
		res = append(res, cs)
		return
	}
	// idle
	if cs.typ == IDLE {
		checkCharge(e)
		if sum(e.a.getVal(), e.b.getVal()) == 8 || sum(e.a.getVal(), e.b.getVal()) == 9 {
			res = append(res, cs)
			return
		}
		// <6
		if sum(cs.a.getVal(), cs.b.getVal()) < 6 {
			for i := 1; i <= 13; i++ {
				res = append(res, &Cardset{
					typ: cs.typ,
					a:   cs.a,
					b:   cs.b,
					and: &card{
						val: i,
					},
				})
			}
			return
		}
		res = append(res, cs)
		return
	}
	return
}

func cal(a, b int) int {
	return (a + b) % 10
}

func sum(a, b int) int {
	return cal(a, b)
}

func (c *card) getVal() int {
	if c.val >= 10 {
		return 0
	}
	return c.val
}

func (cs *Cardset) String() string {
	pend := ""
	if cs.and != nil {
		pend = " " + strconv.Itoa(cs.and.val)
	}
	return "" + strconv.Itoa(cs.a.val) + " " + strconv.Itoa(cs.b.val) + pend
}

func GenerateCardsetsIdle() (res []*Cardset) {
	var cardIA *card
	var cardIB *card
	for i := 1; i <= 13; i++ {
		cardIA = &card{
			val: i,
		}
		for j := i; j <= 13; j++ {
			cardIB = &card{
				val: j,
			}
			res = append(res, &Cardset{
				typ: IDLE,
				a:   cardIA,
				b:   cardIB,
				and: nil,
			})
		}
	}
	return
}

func generateDataSetBySet(a, b int, c, d int) {
	// a>b
	// c>d
	idle := &Cardset{
		typ: IDLE,
		a: &card{
			// A
			val: a,
		},
		b: &card{
			// K
			val: b,
		},
		and: nil,
	}

	charger := &Cardset{
		typ: CHARGE,
		a: &card{
			// 8
			val: c,
		},
		b: &card{
			// K
			val: d,
		},
		and: nil,
	}

	idleset := idle.andCalculate(charger)

	for _, v := range idleset {
		endset := charger.andCalculate(v)
		for _, end := range endset {
			fmt.Println("闲家：", v.String(), " 庄家：", end.String())
		}
	}
}

func calculatePP(charger, idle *Cardset) float64 {

	var mid []int
	mid = append(mid, idle.a.val)
	mid = append(mid, idle.b.val)
	mid = append(mid, charger.a.val)
	mid = append(mid, charger.b.val)
	// a - b - c - d - idle_and - charger_and
	var f1 float64
	var f2 float64
	f1 = 1
	f2 = 1
	if idle.a.val == idle.b.val {
		f1 = f1 * ((32 / total) * (31 / totalMinus(total, 1)))
		if charger.a.val == charger.b.val {

			if charger.a.val == idle.a.val {
				f2 = f2 * ((30 / totalMinus(total, 2)) * (29 / totalMinus(total, 3)))
			} else {
				f2 = f2 * ((32 / totalMinus(total, 2)) * (32 / totalMinus(total, 3)))
			}

		} else {

			if charger.a.val == idle.a.val {
				f2 = f2 * ((30 / totalMinus(total, 2)) * (32 / totalMinus(total, 3)))
			} else {
				f2 = f2 * ((32 / totalMinus(total, 2)) * (32 / totalMinus(total, 3)))
			}

		}
	} else {
		f1 = f1 * ((32 / total) * (32 / totalMinus(total, 1)))
		if charger.a.val == charger.b.val {

			if charger.a.val == idle.a.val {
				f2 = f2 * ((31 / totalMinus(total, 2)) * (30 / totalMinus(total, 3)))
			} else {
				f2 = f2 * ((32 / totalMinus(total, 2)) * (31 / totalMinus(total, 3)))
			}

		} else {

			if charger.a.val == idle.a.val {
				f2 = f2 * ((31 / totalMinus(total, 2)) * (30 / totalMinus(total, 3)))
			} else {
				ca := float64(calculateDuplicate(mid, charger.a.val) - 1)
				cb := float64(calculateDuplicate(mid, charger.b.val) - 1)
				f2 = f2 * ((32 - ca/totalMinus(total, 2)) * ((32 - cb) / totalMinus(total, 3)))
			}

		}
	}
	fmt.Println(idle.a.val, " ", idle.b.val, " ", charger.a.val, " ", charger.b.val, " ", f1*f2)
	return f1 * f2
}

func calculateDuplicate(source []int, input int) int {
	var res int
	for _, v := range source {
		if v == input {
			res++
		}
	}
	return res
}
