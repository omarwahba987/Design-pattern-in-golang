/////  classs digram : https://www.geeksforgeeks.org/builder-design-pattern/
//Separate the construction of a complex object from its representation so that the same construction process
//can create different representations.
//It is used to construct a complex object step by step and the final step will return the object
package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	name  string
	price int
}

func (p *Product) String() string {

	return "Product Name=" + p.name + ", Price=" + strconv.Itoa(p.price) + "]"
}


/**************************************************/ //class director
type Director struct {
	builder Builder
}

func (d *Director) Construct() *Product {
	d.builder.SetName()
	d.builder.SetPrice()

	return d.builder.GetResult()
}

/*****************************************************/
type Builder interface {
	SetName()
	SetPrice()
	GetResult() *Product
}

/*********************************/
type AppleBuilder struct { //////concrete class
	App     string
	product *Product
}

func NewAppleBuilder(ss string) *AppleBuilder {
	return &AppleBuilder{ss, new(Product)}
}

func (b *AppleBuilder) SetName() {
	b.product.name = "apple"
}

func (b *AppleBuilder) SetPrice() {
	b.product.price = 10
}

func (b *AppleBuilder) GetResult() *Product {
	fmt.Println(b.App)
	return b.product
}

/****************************************************************************************/
type OrangeBuilder struct {
	product *Product
	orange  int
}

func NewOrangeBuilder(o int) *OrangeBuilder {
	return &OrangeBuilder{new(Product), o}
}

func (b *OrangeBuilder) SetName() {
	b.product.name = "orange"
}

func (b OrangeBuilder) SetPrice() {
	b.product.price = 20
}
func (b *OrangeBuilder) GetResult() *Product {
	fmt.Println(b.orange)
	return b.product
}


/******************************************************************/
func main() {
	director1 := &Director{NewAppleBuilder("Apple")} //--you can get a pointer to a new instance by doing  t := &Thing{}
	director2 := &Director{NewOrangeBuilder(11)}
	product1 := director1.Construct()
	product2 := director2.Construct()
	fmt.Println(product1)
	fmt.Println(product2)



}


