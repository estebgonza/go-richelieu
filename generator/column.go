package generator

type Column struct {
	// Last generated value
	value string
	// Value generator
	valueGenerator Value
	// Number of value rotation
	rotation_base int
	rotation_mod  int
	count         int
}

func (c *Column) nextValue() string {
	/** The cardinality magic should be here. */
	if c.count == c.rotation_base {
		newValue, _ := c.valueGenerator.GenerateValue()
		c.value = newValue
	} else if c.count == 0 && c.rotation_mod > 0 {
		c.rotation_mod--
	}
	c.count--
	if c.count < 0 || (c.count == 0 && c.rotation_mod == 0) {
		c.count = c.rotation_base
	}
	return c.value
}
