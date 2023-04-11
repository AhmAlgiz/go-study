package delivery

type Calculator interface {
	Calculate(int, int, int) int
}

type DeliveryType struct {
	Calculator Calculator
}

func (d *DeliveryType) Calculate(dist, expTime, weight int) int {
	return d.Calculator.Calculate(dist, expTime, weight)
}
