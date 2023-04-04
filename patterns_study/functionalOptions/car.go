package car

type Car struct {
	Model      string
	Brand      string
	HorsePower int
	Length     int
	Width      int
}

type Option func(*Car)

func NewCar(options ...Option) *Car {
	//default settings
	car := &Car{
		Model:      "",
		Brand:      "",
		HorsePower: 0,
		Length:     0,
		Width:      0,
	}
	//custom settings
	for _, option := range options {
		option(car)
	}
	return car
}

func Model(model string) Option {
	return func(c *Car) {
		c.Model = model
	}
}

func Brand(brand string) Option {
	return func(c *Car) {
		c.Brand = brand
	}
}

func HorsePower(horsePower int) Option {
	return func(c *Car) {
		c.HorsePower = horsePower
	}
}

func Length(length int) Option {
	return func(c *Car) {
		c.Length = length
	}
}

func Width(width int) Option {
	return func(c *Car) {
		c.Width = width
	}
}
