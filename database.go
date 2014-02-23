/*
Credits go to github.com/SlyMarbo/rss for inspiring this solution.
*/
package feeder

type Database interface {
    Request() chan string
    Response() chan bool
    Run()
}

type BaseDatabase struct {
	request  chan string
	response chan bool
}

func (d *BaseDatabase) Request() chan string {
    if d.request == nil {
        d.request = make(chan string)
    }
    return d.request
}

func (d *BaseDatabase) Response() chan bool {
    if d.response == nil {
        d.response = make(chan bool)
    }
    return d.response
}

// This is an in-memory database, which exists for a single
// feed, but does not persist.
//
// It simply maintains a map that knows whether a particular
// channel or item key has been seen before.
type database struct {
    BaseDatabase
	known    map[string]struct{}
}

func (d *database) Run() {
	d.known = make(map[string]struct{})
	var s string

	for {
		s = <-d.Request()
		if _, ok := d.known[s]; ok {
			d.Response() <- true
		} else {
			d.Response() <- false
			d.known[s] = struct{}{}
		}
	}
}

func NewDatabase() *database {
	database := new(database)
	go database.Run()
	return database
}
