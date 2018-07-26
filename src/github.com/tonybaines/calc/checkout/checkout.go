package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func trimNewline(s string) string {
	return strings.TrimSuffix(s, "\n")
}

func safeReadFloat(reader *bufio.Reader, prompt string) (float64, error) {
	fmt.Printf("%s: ", prompt)
	value, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(trimNewline(value), 64)

	if err == nil {
		return number, nil
	} else {
		return float64(-1), err
	}
}

func safeReadInt(reader *bufio.Reader, prompt string) (int64, error) {
	fmt.Printf("%s: ", prompt)
	value, _ := reader.ReadString('\n')
	number, err := strconv.ParseInt(trimNewline(value), 10, 64)

	if err == nil {
		return number, nil
	} else {
		return int64(-1), err
	}
}

type item struct {
	price float64
	qty   int64
}

func main() {

	const tax = 0.55
	items := []item{}
	reader := bufio.NewReader(os.Stdin)

	for i := 1; true; i++ {
		var itemPrice float64
		var itemQty int64

		fmt.Print("Add an item (Y/N)? ")
		ans, _ := reader.ReadString('\n')
		if strings.ToUpper(trimNewline(ans)) != "Y" {
			break
		}

		for true {
			val, err := safeReadFloat(reader, "Please enter the item price")
			if err == nil {
				itemPrice = val
				break
			} else {
				fmt.Printf("Please try again: %s\n", err)
			}
		}
		for true {
			val, err := safeReadInt(reader, "Please enter the number of items")
			if err == nil {
				itemQty = val
				break
			} else {
				fmt.Printf("Please try again: %s\n", err)
			}
		}

		items = append(items, item{itemPrice, itemQty})
	}

	subtotal := 0.0
	for _, item := range items {
		subtotal += item.price * float64(item.qty)
	}

	taxCost := subtotal * tax

	fmt.Printf("Subtotal: %.2f\n", subtotal)
	fmt.Printf("Tax: %.2f\n", taxCost)
	fmt.Printf("Total: %.2f\n", subtotal+taxCost)
}
