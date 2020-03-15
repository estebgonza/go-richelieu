package generator

type Column struct {
	// Last generated value
	value string
	// Value generator
	valueGenerator Value
}

func (c Column) nextValue() string {
	/** The cardinality magic should be here. */
	newValue, _ := c.valueGenerator.GenerateValue()
	c.value = newValue
	return c.value
}
