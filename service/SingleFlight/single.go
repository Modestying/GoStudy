package SingleFlight

import (
	"sync"
)

type (
	SingleFlight interface {
		Do(key string, fn func() (interface{}, error)) (interface{}, error)
		DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error)
	}

	call struct {
		wg  sync.WaitGroup
		val interface{}
		err error
	}

	flightGroup struct {
		calls map[string]*call
		sync.Mutex
	}
)

func NewSingleFlight() SingleFlight {
	return &flightGroup{
		calls: make(map[string]*call),
	}
}
func (g *flightGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	c, done := g.createCall(key)
	if done {
		return c.val, c.err
	}
	g.makeCall(c, key, fn)
	return c.val, c.err
}

func (g *flightGroup) DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error) {
	c, done := g.createCall(key)
	if done {
		return c.val, false, c.err
	}
	g.makeCall(c, key, fn)
	return c.val, true, c.err
}
func (g *flightGroup) createCall(key string) (c *call, done bool) {
	g.Lock()
	if c, ok := g.calls[key]; ok {
		g.Unlock()
		c.wg.Wait()
		return c, true
	}
	c = new(call)
	c.wg.Add(1)
	g.calls[key] = c
	g.Unlock()
	return c, false
}

func (g *flightGroup) makeCall(c *call, key string, fn func() (interface{}, error)) {
	defer func() {
		g.Lock()
		delete(g.calls, key)
		g.Unlock()
		c.wg.Done()
	}()
	c.val, c.err = fn()
}
