package chronosphere

// OptionalValue is used for fields that are marked as optional, where Terraform differentiates
// between null and a default value, but the server does not.
type OptionalValue interface {
	Value() any
	IsEmpty() bool
}

// OptionalString implements OptionalValue for strings.
type OptionalString string

func (s OptionalString) Value() any    { return s }
func (s OptionalString) IsEmpty() bool { return s == "" }

// OptionalMap implements OptionalValue for maps.
type OptionalMap[T comparable, V any] map[T]V

func (m OptionalMap[T, V]) Value() any    { return m }
func (m OptionalMap[T, V]) IsEmpty() bool { return len(m) == 0 }
