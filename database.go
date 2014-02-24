/*
Credits go to github.com/SlyMarbo/rss for inspiring this solution.
*/
package feeder

type Database interface {
    HasChannel(key string) bool
    HasItem(key string) bool
}

// This is an in-memory database, which exists for a single
// feed, but does not persist.
//
// It simply maintains a map that knows whether a particular
// channel or item key has been seen before.
type database struct {
	known    map[string]struct{}
}

func (d *database) HasChannel(key string) bool {
    return d.hasKey(key)
}

func (d *database) HasItem(key string) bool {
    return d.hasKey(key)
}

func (d *database) hasKey(key string) bool {
    if _, ok := d.known[key]; ok {
        return true
    } else {
        d.known[key] = struct{}{}
        return false
    }
}

func NewDatabase() *database {
	database := new(database)
	database.known = make(map[string]struct{})
	return database
}
