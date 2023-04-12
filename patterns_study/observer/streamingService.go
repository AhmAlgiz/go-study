package observer

type StreamingService struct {
	Subject
	observers []Observer
}

func (s *StreamingService) AddObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *StreamingService) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *StreamingService) Notify(msg string) {
	for _, obs := range s.observers {
		obs.Update(msg)
	}
}
