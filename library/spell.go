package library

// Witchboard ...
type Witchboard interface {
	KillAverage() float32
}

// RawIngredients ...
type RawIngredients struct {
	Yoda, Aoda, Yodb, Aodb int
}

// KillAverage ...
func (ri RawIngredients) KillAverage() float32 {
	var r float32

	if ri.Yoda|ri.Aoda|ri.Yodb|ri.Aodb < 0 {
		return -1
	} else {
		ra := KillSpell(ri.Yoda - ri.Aoda)
		rb := KillSpell(ri.Yodb - ri.Aodb)
		r = float32(ra+rb) / 2
		return r
	}

}

// KillSpell ...
func KillSpell(year int) int {

	var bowl []int
	var onion int

	for i := 0; i < year; i++ {
		if i == 0 {
			bowl = append(bowl, 1)
		} else {
			onion = bowl[i-1] + Mantra(i)
			bowl = append(bowl, onion)
		}

	}
	return bowl[len(bowl)-1]

}

// Mantra ...
func Mantra(y int) int {

	var x []int

	for i := 0; i <= y; i++ {

		if i > 2 {
			r := x[i-1] + x[i-2]
			x = append(x, r)
		} else {

			x = append(x, i)
		}

	}
	return x[len(x)-1]
}
