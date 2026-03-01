package main

import "fmt"

//BAD: overusing getters and setters when they add zero value
//This is not idiomatic Go - it mirrors Java/C# habits
//Fields are unexported but getters/setters do nothing
type Customer struct {
	id      int
	name    string
	balance float64
}

//These getters/setters add NO value - no validation, no logic, no mutex
//They jsut wrap field access for no reason
func (c *Customer) GetId() int { // wrong naming too - should be Id() not GetId()
	return c.id
}

func (c *Customer) SetId(id int) {
	c.id = id
}

func (c *Customer) GetName() string { //wrong naming - should be Name()
	return c.name

}

func (c *Customer) SetName(name string) {
	c.name = name
}

func (c *Customer) GetBalance() float64 { //wrong naming should be Balance()
	return c.balance
}

func (c *Customer) SetBalance(balance float64) {
	c.balance = balance
}

func main() {
	c := Customer{}
	c.SetId(1)
	c.SetName("Coder Collo")
	c.SetBalance(1000.0)

	//Verbose for zero benefit - getters/setters add nothing here
	fmt.Println(c.GetId())
	fmt.Println(c.GetName())
	fmt.Println(c.GetBalance())
}
