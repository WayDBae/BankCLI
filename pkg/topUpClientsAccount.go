package pkg

import "fmt"

func AddMoney(name string, amount float64) {
	fmt.Println("\033[2J")
	var money float64
	fmt.Println("На сколько сомони вы хотите пополнить ваш кошелёк?")
	fmt.Scan(&money)
	if money < 0 {
		fmt.Println("Пополнение всегда больше нуля!")
		return
	}
	currentAmount := Database[name]
	Database[name] = currentAmount + money
	for ind := range Database {
		fmt.Printf("Привет, %v! Ваш кошелек успешно пополнен на сумму %v сомони\n", ind, money)
	}
}

func WithdrawMoney(name string, amount float64) {
	fmt.Println("\033[2J")
	var money float64
	fmt.Println("Какую сумму вы хотите снять с вашего кошелька?")
	fmt.Scan(&money)

	if currentAmount, ok := Database[name]; ok {
		if money > currentAmount {
			fmt.Println("Ошибка: недостаточно средств на кошельке.")
		} else {
			Database[name] = currentAmount - money
			fmt.Printf("Привет, %v! Вы сняли сумму: %v. Остаток на кошельке: %v\n", name, money, Database[name])
		}
	} else {
		fmt.Println("Ошибка: такого клиента не существует!")
	}
}