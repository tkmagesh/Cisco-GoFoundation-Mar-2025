package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

/*
Create a "Sort" api (function) to sort the products by any attribute
	IMPORTANT: DO NOT write your own logic for sorting. instead use sort.Sort() function
*/

type Products []Product

func (products Products) String() string {
	var builder strings.Builder
	for _, product := range products {
		builder.WriteString(fmt.Sprintf("%s\n", product))
	}
	return builder.String()
}

// sort.Interface interface implementation
func (products Products) Len() int {
	return len(products)
}

// compare products by Id
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

// Comparison by Name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

// Comparison by Cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

// Comparison by Units
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

// Comparison by Category
type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

const (
	Id       = "Id"
	Name     = "Name"
	Cost     = "Cost"
	Units    = "Units"
	Category = "Category"
)

// utility method
func (products Products) Sort(attrName string) {
	switch attrName {
	case Id:
		sort.Sort(products)
	case Name:
		sort.Sort(ByName{products})
	case Cost:
		sort.Sort(ByCost{products})
	case Units:
		sort.Sort(ByUnits{products})
	case Category:
		sort.Sort(ByCategory{products})
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Sort by Id")
	// sort.Sort(products)
	products.Sort(Id)
	fmt.Println(products)

	fmt.Println("Sort by Name")
	// sort.Sort(ByName{products})
	products.Sort(Name)
	fmt.Println(products)

	fmt.Println("Sort by Cost")
	// sort.Sort(ByCost{products})
	products.Sort(Cost)
	fmt.Println(products)

	fmt.Println("Sort by Units")
	// sort.Sort(ByUnits{products})
	products.Sort(Units)
	fmt.Println(products)

	fmt.Println("Sort by Category")
	// sort.Sort(ByCategory{products})
	products.Sort(Category)
	fmt.Println(products)
}
