package pagination

// Filter is a common set of filters for List calls.
type Filter struct {
	Slugs []string
	Names []string
}

// IsEmpty returns whether any filters are set.
func (f Filter) IsEmpty() bool {
	return len(f.Names)+len(f.Slugs) == 0
}
