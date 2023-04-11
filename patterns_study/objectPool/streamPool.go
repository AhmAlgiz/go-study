package streamPool

type StreamPool chan *Stream

type Stream struct {
	Data *StreamData
}

type StreamData struct {
	Packages []string
}

func (sd *StreamData) SendPackage(data string) {
	sd.Packages = append(sd.Packages, data)
}

func (sd *StreamData) GetPackages(size int) []string {
	data := make([]string, 0, size)
	sdSize := len(sd.Packages)
	if sdSize <= size {
		data = sd.Packages
	} else {
		data = sd.Packages[sdSize-size-1 : sdSize-1]
	}
	return data
}

func (s *Stream) GetLastPackage() string {
	return s.Data.GetPackages(1)[0]
}

// Создание пула(канала) с объектами Stream, у которых один и тот же источник данных
func NewPool(poolCount int, sd *StreamData) *StreamPool {
	sp := make(StreamPool, poolCount)
	for i := 0; i < poolCount; i++ {
		s := new(Stream)
		s.Data = sd
		sp <- s
	}
	return &sp
}

func NewData() *StreamData {
	return new(StreamData)
}
