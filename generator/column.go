package generator

// Column Store useful values for rotation, as well as the current value struct of the column
type Column struct {
	colName string
	// Last generated value
	value string
	// Value generator
	valueGenerator value
	// Number of value rotation
	rotationBase int
	rotationMod  int
	count        int
	totCount     uint64
}

func (c *Column) nextValue() string {
	switch c.valueGenerator.(type) {
	case *idIntValue:
		c.valueGenerator.generateValue()
	default:
		if c.totCount == 0 {
			// Do nothing. We already have init value from plan.json or default one
		} else if c.count == c.rotationBase {
			c.valueGenerator.generateValue()
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
