package main

type LanternFish byte

func (l *LanternFish) Age() []LanternFish {
	if *l == 0 {
		(*l) = 6
		return []LanternFish{8}
	} else {
		(*l)--
		return []LanternFish{}
	}
}

type FishTank [9]int

func (t *FishTank) Age() {
	t[0], t[1], t[2], t[3], t[4], t[5], t[6], t[7], t[8] = t[1], t[2], t[3], t[4], t[5], t[6], t[0]+t[7], t[8], t[0]
}

func (t *FishTank) Count() int {
	sum := 0
	for i := 0; i < len(t); i++ {
		sum += t[i]
	}
	return sum
}
