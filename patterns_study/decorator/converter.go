package converter

func CalculateCostRubles(base, days int) int {
	return base * days
}

type Fn func(int, int) int

func ConvertRubblesToDollars(fn Fn, rate int) Fn {
	return func(base int, days int) int {
		res := fn(base, days) * rate
		return res
	}
}
