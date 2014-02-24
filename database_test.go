package feeder

import (
    "testing"
)

func Test_Item(t *testing.T) {
    db := NewDatabase()

    if db.HasItem("key") {
        t.Errorf("Should not have found the key")
    }
    if !db.HasItem("key") {
        t.Errorf("Should have found the key")
    }
}

func Test_Channel(t *testing.T) {
    db := NewDatabase()

    if db.HasChannel("key") {
        t.Errorf("Should not have found the key")
    }
    if !db.HasChannel("key") {
        t.Errorf("Should have found the key")
    }
}
