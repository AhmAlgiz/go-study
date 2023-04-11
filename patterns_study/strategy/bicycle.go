package delivery

type Bicycle struct{}

func (Bicycle) Calculate(dist, expTime, weight int) int {
	return weight * dist * expTime / 5
}
