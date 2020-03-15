package generator

type Column struct {
	// Last generated value
	value string
	// Value generator
	valueGenerator Value
	// Number of value rotation
	rotation int
	count    int
}

func (c *Column) resetOne() {
	c.count = c.rotation
}

func (c *Column) subOne() {
	c.count--
}

func (c *Column) nextValue() string {
	/** The cardinality magic should be here. */
	if c.count == c.rotation {
		newValue, _ := c.valueGenerator.GenerateValue()
		c.value = newValue
	}
	c.count = c.count - 1
	if c.count <= 0 {
		c.resetOne()
	}
	return c.value
}
