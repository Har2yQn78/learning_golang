package main

import "fmt"

type bill struct {
	names string 
	items map[string]float64
	tip float64
}

func newBill(name string) bill {
	b := bill{
		names: name,
		items: map[string]float64{"iced coffee": 98, "crossant": 153},
		tip: 0,	  
	}

	return b 
}

func (b bill) format() string{
	fs := "Bill breakdown:\n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%v...%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v..%0.2f", "total:", total)

	return fs
}