package cart

type Cart struct {
	Items map[string]Item
}

type Item struct {
	Name  string
	ID    string
	Price float64
	Qty   int
}

func (c *Cart) AddItem(i Item) {
	c.Items = map[string]Item{}
	c.Items["first"] = i
}

func (c *Cart) RemoveItem(id string, n int) {}

func (c *Cart) Totalmount() int {
	return len(c.Items)
}

func (c *Cart) TotalUnits() int {
	return 0
}

func (c *Cart) TotalUniqueItems() int {
	return 0
}
