package main

import "fmt"

func main(){
	var apple, pear, money float64 = 5.99, 7, 23

	nineAppleCoast := 9 * apple
	eightPearCoast := 8 * pear
	summ := nineAppleCoast+eightPearCoast

	resBuyApple := money/apple
	resBuyPear := money/pear

	resultBool := (pear*2 + apple*2)<money
	var res string
	if resultBool == true {
		res = "так"
	}else if resultBool == false {
		res = "ні"
	}

	fmt.Printf("1) Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\n9 яблук будуть коштувати: %.2f грн\n8 груш будуть коштувати: %.2f грн\nРазом: %.2f грн\n", nineAppleCoast, eightPearCoast, summ)
	fmt.Printf("2) Скільки груш ми можемо купити?\nЗа наші кошти ми можемо купити: %d груш\n", int(resBuyPear))
	fmt.Printf("3) Скільки яблук ми зможемо купити?\nЗа наші кошти ми можемо купити: %d яблук\n", int(resBuyApple))
	fmt.Printf("4) Чи ми можемо купити 2 груші та 2 яблука?\nЗвісно %s", res)
}
