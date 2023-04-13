package concurrency

type Connection struct {
	name string
}

func readFromDB(c Connection) string {
	return c.name
}

func Query(connections []Connection, query string) string {
	ch := make(chan string, 1)
	for _, c := range connections {
		go func(conn Connection) {
			select {
			case ch <- readFromDB(conn):
			default:
			}
		}(c)
	}
	return <-ch
}
