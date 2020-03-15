package generator

// import (
// 	"fmt"
// )

type Column struct {
	// Last generated value
	value string
	// Value generator
	valueGenerator Value
	// Number of value rotation
	rotation int
	count    int
}

func (c *Column) countOne() {
	c.count = 1
}

func (c *Column) countAdd() {
	c.count += 1
}

func (c *Column) nextValue() string {
	/** The cardinality magic should be here. */
	// fmt.Println(c.count)
	if c.count == 1 {
		// fmt.Println(c.rotation)
		// fmt.Println(c.count)
		newValue, _ := c.valueGenerator.GenerateValue()
		c.value = newValue
	}
	c.count = c.count + 1
	// fmt.Println(c.count)
	// if c.count > c.rotation {
	// 	c.countOne()
	// }
	return c.value
}
