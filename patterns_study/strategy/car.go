package delivery

type Car struct {
	DeliveryType
}

func (Car) Calculate(dist, expTime, weight int) int {
	return weight * dist * expTime / 2
}
