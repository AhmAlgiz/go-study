package delivery

type Walker struct{}

func (Walker) Calculate(dist, expTime, weight int) int {
	return weight * dist * expTime / 10
}
