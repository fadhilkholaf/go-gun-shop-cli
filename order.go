package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

type Order struct {
	name      string
	order     []Gun
	tip       uint
	orderedAt time.Time
}

func createOrder(name string) Order {
	r := Order{
		name: name, order: []Gun{}, tip: 0, orderedAt: time.Now(),
	}
	return r
}

func (o *Order) AddGun(name string, price uint) {
	o.order = append(o.order, Gun{name, price})
}

func (o *Order) addTip(tip uint) {
	o.tip = tip
}

func idr(n uint) string {
	c, _ := currency.FromTag(language.Indonesian)
	s, _ := currency.Cash.Rounding(c)
	f := number.Decimal(n, number.Scale(s))
	p := message.NewPrinter(language.Indonesian)
	fS := p.Sprintf("%v %v", currency.Symbol(c), f)
	return fS
}

func (o Order) Format() string {
	f := fmt.Sprintf("\n%v\n%-13v | %v\n", "=====================", "Order name", o.name+"'s order")
	var t uint = 0
	f += fmt.Sprintf("%-13v | Price\n---------------------\n", "Gun name")
	for _, g := range o.order {
		f += fmt.Sprintf("%-13v | %v\n", g.name, idr(g.price))
		t += g.price
	}
	if o.tip != 0 {
		f += fmt.Sprintf("%-13v | %v\n", "Order tip", idr(o.tip))
	}
	f += fmt.Sprintf("---------------------\n%-13v | %v\n", "Total price", idr(o.tip+t))
	f += fmt.Sprintf("%-13v | %v\n=====================", "Ordered at", o.orderedAt.Format("02 Jan 06 15:04 MST"))
	return f
}

func (o *Order) Save() bool {
	b := []byte(o.Format())
	e := os.WriteFile("orders/"+o.name+"'s.order"+"-"+strings.ReplaceAll(o.orderedAt.Format("2006-01-02T15:04:05.999999999Z07:00"), ":", ".")+".txt", b, 0644)
	if e != nil {
		return false
	} else {
		return true
	}
}
