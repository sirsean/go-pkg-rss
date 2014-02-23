package feeder

import (
    "testing"
)

func Test_Empty(t *testing.T) {
    db := NewDatabase()

    if db.Request() <- "key"; <-db.Response() {
        t.Errorf("Should not have found the key")
    }
    if db.Request() <- "key"; !<-db.Response() {
        t.Errorf("Should have found the key")
    }
}
