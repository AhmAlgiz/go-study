package starship

type Starship struct {
	Name        string
	MaxSpeed    int
	FuelStorage int
	Mass        int
	Capacity    int
}

type StarshipBuilder struct {
	name        string
	maxSpeed    int
	fuelStorage int
	mass        int
	capacity    int
}

func NewStarshipBuilder() *StarshipBuilder {
	return &StarshipBuilder{}
}

func (sb *StarshipBuilder) SetName(name string) *StarshipBuilder {
	sb.name = name
	return sb
}

func (sb *StarshipBuilder) SetMaxSpeed(maxSpeed int) *StarshipBuilder {
	sb.maxSpeed = maxSpeed
	return sb
}

func (sb *StarshipBuilder) SetFuelStorage(capacity int) *StarshipBuilder {
	sb.fuelStorage = capacity
	return sb
}

func (sb *StarshipBuilder) SetMass(mass int) *StarshipBuilder {
	sb.mass = mass
	return sb
}

func (sb *StarshipBuilder) SetCapacity(capacity int) *StarshipBuilder {
	sb.capacity = capacity
	return sb
}

func (sb *StarshipBuilder) Build() *Starship {
	return &Starship{
		Name:        sb.name,
		MaxSpeed:    sb.maxSpeed,
		FuelStorage: sb.fuelStorage,
		Mass:        sb.mass,
		Capacity:    sb.capacity,
	}
}
