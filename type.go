package argpas

import "reflect"

type argPair struct {
	field reflect.Value
	name  string
}

// Trigger IsTrigger will be true if argument exist
type Trigger struct {
	IsTrigger bool
}
