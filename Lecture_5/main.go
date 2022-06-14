package main

import "fmt"

func main(){
	var apple = 5.99
	var pear, money = 7, 23

	nineAppleCoast := 9 * apple
	eightPearCoast := 8 * pear

	resBuyApple := float64(money)/apple
	resBuyPear := money/pear

	resultBool := float64(pear)*2+apple*2<float64(money)
	var res string
	switch resultBool {
	case true:	res = "так"
	case false: res = "ні"
	}

	fmt.Printf("1) Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\n9 яблук будуть коштувати: %.2f грн\n8 груш будуть коштувати: %d грн\n", nineAppleCoast, eightPearCoast)
	fmt.Printf("2) Скільки груш ми можемо купити?\nЗа наші кошти ми можемо купити: %d груш\n", resBuyPear)
	fmt.Printf("3) Скільки яблук ми зможемо купити?\nЗа наші кошти ми можемо купити: %d яблук\n", int(resBuyApple))
	fmt.Printf("4) Чи ми можемо купити 2 груші та 2 яблука?\nЗвісно %s", res)
}
