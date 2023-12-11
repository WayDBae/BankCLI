package pkg

import "fmt"

func TransferMoney() {
	var sender string
	fmt.Println("Введите ваше имя: ")
	fmt.Scan(&sender)
	balanceSender, ok := Database[sender]
	if !ok {
		fmt.Println("Отправителя не существует в базе!")
		return
	}
	var receipent string
	fmt.Println("Введите имя того, кому желаете перевести: ")
	fmt.Scan(&receipent)
	balanceReceiper, ok := Database[receipent]
	if !ok {
		fmt.Println("Получателя не существует в базе!")
		return
	}
	var amount, amountToTransfer float64
	fmt.Println("Какую сумму вы хотите перевести с вашего кошелька?")
	fmt.Scan(&amountToTransfer)
	var profit float64
	if ok {
		if amountToTransfer > balanceSender {
			fmt.Println("Ошибка: недостаточно средств на кошельке.")
		} else {
			amount = amountToTransfer
			profit = amount / 100 * Percent
			amountWithoutProfit := amount - profit
			Database[sender] = balanceSender - amount
			Database[receipent] = balanceReceiper + amountWithoutProfit
			balanceProfit, ok := Database["Profit"]
			if !ok{
				Database["Profit"] = 0
			}
			Database["Profit"] = balanceProfit + profit
			fmt.Printf("Привет, %v! Вы сняли сумму: %v. \n", sender, amountToTransfer)
			fmt.Printf("На кошельке у %v теперь %v сомони. \n", receipent, Database[receipent])
			fmt.Printf("Баланс Khayrakoff bank: %v\n", Database["Profit"])
		}
	} else {
		fmt.Println("Ошибка: такого клиента не существует!")
	}
}
