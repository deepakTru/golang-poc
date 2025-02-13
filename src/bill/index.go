package bill

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

func (b *Bill) addItem(itemName string, itemPrice float64) {
	b.items[itemName] = itemPrice
}
func (b *Bill) addTip(tip float64) {
	b.tip = tip
}

func (b *Bill) formatBill() string {
	var formatted string = ""
	total := b.tip
	formatted += fmt.Sprintf("\n%-25v %v\n", "Bill Name:", b.name)
	for key, value := range b.items {
		total += value
		formatted += fmt.Sprintf("%-25v ..$%.2f\n", key+":", value)
	}
	formatted += fmt.Sprintf("%-25v ..$%.2f\n", "Tip:", b.tip)
	formatted += fmt.Sprintf("%-25v ..$%.2f\n", "Total:", total)

	return formatted
}

func promptReader(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')
	return strings.TrimSpace(input), error
}

func addItemsInBill(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := promptReader(reader, "Select Options: (Add Item: a),  (Add Tip: t),  (Save: s) :")

	switch option {
	case "a":
		fmt.Println("You choose a")
		itemName, _ := promptReader(reader, "Enter Item Name: ")
		itemPrice, _ := promptReader(reader, "Enter Item Price: ")
		floatPrice, err := strconv.ParseFloat(itemPrice, 64)
		if err != nil {
			fmt.Println("Input Item Price!")
			addItemsInBill(b)
		}

		b.addItem(itemName, floatPrice)
		addItemsInBill(b)

	case "t":
		fmt.Println("You choose t")
		tip, _ := promptReader(reader, "Enter Tip Amount: ")
		floatTip, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Input Tip Amount!")
			addItemsInBill(b)
		}
		b.addTip(floatTip)
		addItemsInBill(b)
	case "s":
		err := os.WriteFile(fmt.Sprintf("src/bill/%v.txt", b.name), []byte(b.formatBill()), 0644)
		if err != nil {
			fmt.Print("Error writing bill", err)
		}
		fmt.Println("Bill has been saved!")
	default:
		fmt.Println("Invalid Option!")
		addItemsInBill(b)
	}

}

func CreateBill() {
	reader := bufio.NewReader(os.Stdin)
	billName, _ := promptReader(reader, "Enter Bill Name: ")
	bill := Bill{name: billName, items: map[string]float64{}, tip: 0}
	addItemsInBill(bill)
}
