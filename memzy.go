package memzy

// Memzy provides an interface between go objects and persistence layers
//
// Current backends that implement this:
// - dynamodb
// - memory
type Memzy interface {
	GetItem(interface{}, map[string]interface{}) error
	PutItem(interface{}) error
}
