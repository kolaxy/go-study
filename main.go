package main

import "fmt"

type Char struct {
	On    bool
	Ammo  int
	Power int
}

func (c *Char) Shoot() bool {
	if c.On == false || c.Ammo <= 0 {
		return false
	} else {
		c.Ammo = c.Ammo - 1
		return true
	}
}

func (c *Char) RideBike() bool {
	if c.On == false || c.Power <= 0 {
		return false
	} else {
		c.Power = c.Power - 1
		return true
	}
}

func main() {
	testStruct := &Char{}
}
