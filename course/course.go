package course

import "fmt"

type course struct {
	name    string
	Price   float64
	isfree  bool
	userId  []uint
	classes map[uint]string
}

func (c course) imprimirclase() {
	text := "Las clases son: "
	for _, class := range c.classes {
		text += class + " ,"
	}
	fmt.Println(text[:len(text)-2])
}

func (c *course) Actualizarprecio(price float64) {
	c.Price = price

}
