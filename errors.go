package memzy


import "fmt"

// ErrNotFound is returned if you try to get an item that doesn't exist
var ErrNotFound = fmt.Errorf("Item not found")
