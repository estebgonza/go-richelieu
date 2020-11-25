package generator

type Column struct {
	colName string
	// Last generated value
	value string
	// Value generator
	valueGenerator Value
	// Number of value rotation
	rotationBase int
	rotationMod  int
	count        int
	totCount     uint64
}

func (c *Column) nextValue() string {
	/** The cardinality magic should be here. */
	switch c.valueGenerator.(type) {
	case IdIntValue:
		newValue, _ := c.valueGenerator.GenerateValue()
		c.value = newValue
	default:
		if c.count == c.rotationBase {
			newValue, _ := c.valueGenerator.GenerateValue()
			c.value = newValue
		} else if c.count == 0 && c.rotationMod > 0 {
			c.rotationMod--
		}
		c.count--
		if c.count < 0 || (c.count == 0 && c.rotationMod == 0) {
			c.count = c.rotationBase
		}
		c.totCount++
	}
	return c.value
}
