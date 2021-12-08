package main

type LanternFish int

func (l *LanternFish) Age() []LanternFish {
	if *l == 0 {
		(*l) = 6
		return []LanternFish{8}
	} else {
		(*l)--
		return []LanternFish{}
	}
}
