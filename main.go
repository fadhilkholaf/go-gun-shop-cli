package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(r *bufio.Reader, m string) (string, error) {
	fmt.Print(m)
	i, e := r.ReadString('\n')
	return strings.TrimSpace(i), e
}

func (o *Order) orderOptions(r *bufio.Reader) {
	s, _ := input(r, "---------------------\na - Add new gun order\nt - Give some tips\ns - Save and pay\ne - Unsave and exit shop\nType the alphabet to continue: ")
	switch s {
	case "a":
		gN, _ := input(r, "---------------------\nInput gun name: ")
		gP, _ := input(r, "Input gun price: ")
		cGP, e := strconv.ParseUint(gP, 10, 32)
		if e != nil {
			fmt.Println("---------------------\nError: Invalid price!")
			o.orderOptions(r)
		} else {
			o.AddGun(gN, uint(cGP))
			o.orderOptions(r)
		}
	case "t":
		t, _ := input(r, "---------------------\nInput tip: ")
		cT, e := strconv.ParseUint(t, 10, 32)
		if e != nil {
			fmt.Println("---------------------\nError: Invalid price!")
			o.orderOptions(r)
		} else {
			o.addTip(uint(cT))
			o.orderOptions(r)
		}
	case "s":
		if len(o.order) == 0 {
			fmt.Println("---------------------\nError: You must have at least one ordered gun to save!")
			o.orderOptions(r)
		} else {
			e := o.Save()
			if e {
				fmt.Println(o.Format())
				fmt.Println("Order saved, thanks for shoping!")
			} else {
				fmt.Println("Error: Cannot saving order!")
				o.orderOptions(r)
			}
		}
	case "e":
		fmt.Println("---------------------\nExiting shop, thanks for comming!")
	default:
		fmt.Println("---------------------\nError: Command not found!")
		o.orderOptions(r)
	}
}

func main() {
	fmt.Print(`
 ______   __  __   __   __       ______   __  __   ______   ______     ______   __       __
/\  ___\ /\ \/\ \ /\ "-.\ \     /\  ___\ /\ \_\ \ /\  __ \ /\  == \   /\  ___\ /\ \     /\ \
\ \ \__ \\ \ \_\ \\ \ \-.  \    \ \___  \\ \  __ \\ \ \/\ \\ \  _-/   \ \ \____\ \ \____\ \ \
 \ \_____\\ \_____\\ \_\\"\_\    \/\_____\\ \_\ \_\\ \_____\\ \_\      \ \_____\\ \_____\\ \_\
  \/_____/ \/_____/ \/_/ \/_/     \/_____/ \/_/\/_/ \/_____/ \/_/       \/_____/ \/_____/ \/_/
` + "\n")
	r := bufio.NewReader(os.Stdin)
	n, _ := input(r, "---------------------\nCreate new order\nYour name: ")
	o := createOrder(n)
	o.orderOptions(r)
}
