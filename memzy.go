package memzy

// Memzy provides an interface between go objects and persistence layers
//
// Current backends that implement this:
// - dynamodb
// - memory
type Memzy interface {
	// GetItem retrieves and unmarshals the object described by the key into the provided object
	GetItem(interface{}, map[string]interface{}) error
	// PutItem stores the object
	PutItem(interface{}) error
	// NewItem returns an Iter containing all of the objects that satisfy the conditions provided
	NewIter(...interface{}) Iter
}

// Iter provides an interface to loop through a subset of objects
type Iter interface {
	// Current unmarshals the object at the current pointer into the provided object
	Current(interface{}) error
	// Err returns the error that cause Next() to fail
	Err() error
	// Next moves the pointer along if possible, else returns false
	Next() bool
}
