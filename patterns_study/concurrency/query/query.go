package query

type Connection struct {
	name string
}

func readFromDB(c Connection, s string) string {
	return c.name + s
}

func Query(connections []Connection, query string) string {
	ch := make(chan string, 1)
	for _, c := range connections {
		go func(conn Connection) {
			select {
			case ch <- readFromDB(conn, query):
			default:
			}
		}(c)
	}
	return <-ch
}
