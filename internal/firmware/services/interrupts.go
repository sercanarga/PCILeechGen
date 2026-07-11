package services

type Delivery struct {
	Valid  bool
	Vector int
}

type InterruptController struct {
	enabled      bool
	functionMask bool
	vectorMask   []bool
	pending      []bool
}

func NewInterruptController(vectors int) *InterruptController {
	if vectors < 0 {
		vectors = 0
	}
	return &InterruptController{
		vectorMask: make([]bool, vectors),
		pending:    make([]bool, vectors),
	}
}

func (c *InterruptController) Request(vector int) Delivery {
	if !c.valid(vector) {
		return Delivery{}
	}
	if c.deliverable(vector) {
		return Delivery{Valid: true, Vector: vector}
	}
	c.pending[vector] = true
	return Delivery{}
}

func (c *InterruptController) SetEnabled(enabled bool) []Delivery {
	if c == nil {
		return nil
	}
	c.enabled = enabled
	return c.drainOne()
}

func (c *InterruptController) SetFunctionMask(masked bool) []Delivery {
	if c == nil {
		return nil
	}
	c.functionMask = masked
	return c.drainOne()
}

func (c *InterruptController) SetVectorMask(vector int, masked bool) []Delivery {
	if !c.valid(vector) {
		return nil
	}
	c.vectorMask[vector] = masked
	if !masked && c.pending[vector] && c.deliverable(vector) {
		c.pending[vector] = false
		return []Delivery{{Valid: true, Vector: vector}}
	}
	return nil
}

func (c *InterruptController) Drain() Delivery {
	if deliveries := c.drainOne(); len(deliveries) != 0 {
		return deliveries[0]
	}
	return Delivery{}
}

func (c *InterruptController) Pending(vector int) bool {
	return c.valid(vector) && c.pending[vector]
}

func (c *InterruptController) Reset() {
	if c == nil {
		return
	}
	c.enabled = false
	c.functionMask = false
	for i := range c.vectorMask {
		c.vectorMask[i] = false
		c.pending[i] = false
	}
}

func (c *InterruptController) valid(vector int) bool {
	return c != nil && vector >= 0 && vector < len(c.pending)
}

func (c *InterruptController) deliverable(vector int) bool {
	return c.enabled && !c.functionMask && !c.vectorMask[vector]
}

func (c *InterruptController) drainOne() []Delivery {
	if c == nil || !c.enabled || c.functionMask {
		return nil
	}
	for vector := range c.pending {
		if c.pending[vector] && !c.vectorMask[vector] {
			c.pending[vector] = false
			return []Delivery{{Valid: true, Vector: vector}}
		}
	}
	return nil
}
