package base

// FilteredAttributes represent a mapping from the output
// generated by a filter to the original value.
type FilteredAttribute struct {
	Old Attribute
	New Attribute
}

// Filters transform the byte sequences stored in DataGrid
// implementations.
type Filter interface {
	// Adds an Attribute to the filter
	AddAttribute(Attribute) error
	// Allows mapping old to new Attributes
	GetAttributesAfterFiltering() []FilteredAttribute
	// Gets a string for printing
	String() string
	// Accepts an old Attribute, the new one and returns a sequence
	Transform(Attribute, Attribute, []byte) []byte
	// Builds the filter
	Train() error
}
