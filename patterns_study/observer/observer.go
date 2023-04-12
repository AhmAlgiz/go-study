package observer

type Observer interface {
	Update(string)
}

type Subject interface {
	AddObserver(Observer)
	RemoveObserver(Observer)
	Notify()
}
