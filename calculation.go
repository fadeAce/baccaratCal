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
	var chargerWin float64
	var chargerdie float64
	var chargerlose float64

	total = 13.0 * 32.0

	step = 1.0

	idles := GenerateCardsetsIdle()
	chargers := GenerateCardsetsIdle()
	for _, i := range idles {
		fmt.Println("============ ", i.a.val, " ", i.b.val)
		for _, j := range chargers {
			_, w, d, l := generateDataSetBySet(i.a.val, i.b.val, j.a.val, j.b.val)
			chargerWin += w
			chargerdie += d
			chargerlose += l
		}

	}

	// calculate probability
	// 8281 is correct
	var sum int
	var f float64

	fmt.Println("< calculation start >")
	for _, i := range idles {
		for _, j := range chargers {
			sum++
			f += calculatePP(i, j)
		}
	}
	fmt.Println("all types sum to ", f)
	fmt.Println("the root types of base card range is ", sum)

	fmt.Println("the result win is ", chargerWin, " chargerlose ", chargerlose, " chargerdie ", chargerdie)
	fmt.Println("sum is ", chargerWin+chargerlose+chargerdie)
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

func generateDataSetBySet(a, b int, c, d int) (all, win, die, lose float64) {
	// todo : mark each loop a unique dataset
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

	dep := calculatePP(charger, idle)

	var count int

	idleset := idle.andCalculate(charger)

	for _, v := range idleset {
		endset := charger.andCalculate(v)
		for _, end := range endset {
			m := calculatePPSuffix(end, v)
			all = all + m
			res := v.compare(end)
			if res == "le" {
				win += dep * m
			}
			if res == "eq" {
				die += dep * m
			}
			if res == "la" {
				lose += dep * m
			}
			fmt.Println("闲家：", v.String(), " 庄家：", end.String())
			if end.and != nil {
				count++
			}
		}
	}

	fmt.Println("all is ", all, " idle and sum is ", len(idleset), " charger and sum is ", count, " dep is ", dep)
	return
}

func calculatePP(charger, idle *Cardset) float64 {
	var final float64
	var mid []int
	mid = append(mid, idle.a.val)
	mid = append(mid, idle.b.val)
	// a - b - c - d - idle_and - charger_and
	var f1 float64
	var f2 float64
	f1 = 1
	f2 = 1
	if idle.a.val == idle.b.val {
		f1 = f1 * ((32 / total) * (31 / totalMinus(total, 1)))
	} else {
		f1 = f1 * ((32 / total) * (32 / totalMinus(total, 1))) * 2
	}

	atimes := calculateDuplicate(mid, charger.a.val)
	mid = append(mid, charger.a.val)
	btimes := calculateDuplicate(mid, charger.b.val)
	f2 = f2 * (((32 - float64(atimes)) / totalMinus(total, 2)) * ((32 - float64(btimes)) / totalMinus(total, 3)))
	if charger.a.val != charger.b.val {
		f2 = f2 * 2
	}
	fmt.Println(idle.a.val, " ", idle.b.val, " ", charger.a.val, " ", charger.b.val, " ", f1*f2, " atimes ", atimes, "btimes ", btimes)
	final = f1 * f2

	var ctimes int
	var dtimes int
	if idle.and != nil {
		ctimes = calculateDuplicate(mid, idle.and.val)
		mid = append(mid, idle.and.val)
		final = final * ((32 - float64(ctimes)) / totalMinus(total, 4))
		if charger.and != nil {
			dtimes = calculateDuplicate(mid, charger.and.val)
			final = final * ((32 - float64(dtimes)) / totalMinus(total, 5))
		}
	} else {
		if charger.and != nil {
			dtimes = calculateDuplicate(mid, charger.and.val)
			final = final * ((32 - float64(dtimes)) / totalMinus(total, 4))
		}
	}
	return final
}
func calculatePPSuffix(charger, idle *Cardset) float64 {
	var final float64
	final = 1.0
	var mid []int
	mid = append(mid, idle.a.val)
	mid = append(mid, idle.b.val)
	mid = append(mid, charger.a.val)
	mid = append(mid, charger.b.val)

	var ctimes int
	var dtimes int
	if idle.and != nil {
		ctimes = calculateDuplicate(mid, idle.and.val)
		mid = append(mid, idle.and.val)
		final = final * ((32 - float64(ctimes)) / totalMinus(total, 4))
		if charger.and != nil {
			dtimes = calculateDuplicate(mid, charger.and.val)
			final = final * ((32 - float64(dtimes)) / totalMinus(total, 5))
		}
	} else {
		if charger.and != nil {
			dtimes = calculateDuplicate(mid, charger.and.val)
			final = final * ((32 - float64(dtimes)) / totalMinus(total, 4))
		}
	}
	return final
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
