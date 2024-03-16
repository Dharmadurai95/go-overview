package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(n string) Bill {
	bill := Bill{
		name:  n,
		items: map[string]float64{"water": 19.99},
		tip:   0,
	}
	return bill
}

func (b Bill) formatBill() string {
	billSummary := "Bill Summary: \n"
	var total float64 = 0
	for item, value := range b.items {

		str := fmt.Sprintf("%-25v ....%v \n", item+":", value)
		billSummary += str
		total += value
	}
	tip := fmt.Sprintf("%-25v ....%0.2f \n", "tip:", b.tip)
	totalVal := fmt.Sprintf("%-25v ....%0.2f \n", "total:", total+b.tip)
	billSummary += tip
	billSummary += totalVal
	return billSummary

}

func (b *Bill) addTip(amt float64) { //if we pass ref to function it will not create new value otherwise it will new (Bill) value
	b.tip += amt
}
func (b *Bill) addItem(item string, amt float64) {
	b.items[item] = amt
}
func (b Bill) saveFile() {
	bytes := []byte(b.formatBill())
	err := os.WriteFile(b.name+".txt", bytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved in Bills folder/directory")
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	str, err := r.ReadString('\n')
	return strings.TrimSpace(str), err
}

func promptOption(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item s -  save bill t - add tip)", reader)
	switch opt {
	case "a":
		itemName, _ := getInput("Item name:", reader)
		price, _ := getInput("Itme Price ($):", reader)
		fPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number:")
			promptOption(b)
		}
		b.addItem(itemName, fPrice)
		fmt.Printf("You added a item %v and this price %v\n", itemName, price)
		promptOption(b)
	case "t":
		tipAmt, _ := getInput("Enter tip amount ($):", reader)
		fTip, err := strconv.ParseFloat(tipAmt, 64)
		if err != nil {
			fmt.Println("Tip Must be a number:")
			promptOption(b)
		}
		b.addTip(fTip)
		fmt.Println("You added tip ($):", tipAmt)
		promptOption(b)
	case "s":
		b.saveFile()
		fmt.Println("you saved the file --", b.name)
	default:
		fmt.Println("That was not an valid option...")
		promptOption(b)

	}

}

func createBill() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name:", reader)
	newBill := newBill(name)
	fmt.Println(newBill)
	promptOption(newBill)

}
