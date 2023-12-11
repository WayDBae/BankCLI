package pkg

import "fmt"

func AddMoney() {
	var name string
	var amount float64
	fmt.Println("Введите имя клиента:")
	fmt.Scan(&name)

	if _, exists := Database[name]; !exists {
		fmt.Println("Ошибка: клиент с таким именем не существует!")
		return
	}
	fmt.Println("\033[2J")
	fmt.Println("На сколько сомони вы хотите пополнить ваш кошелёк?")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("Пополнение всегда больше нуля!")
		return
	}
	balance := Database[name]
	Database[name] = balance + amount
	for ind, v := range Database {
		fmt.Printf("Привет, %v! Ваш кошелек успешно пополнен на сумму %v сомони\n", ind, v)
	}
}

func WithdrawMoney() {
	fmt.Println("\033[2J")
	var name string
	var amount float64

	fmt.Println("Введите имя клиента: ")
	fmt.Scan(&name)
	balance, ok := Database[name]
	if !ok {
		fmt.Println("Такого пользователя не существует!")
		return
	}
	fmt.Println("Какую сумму вы хотите снять с вашего кошелька?")
	fmt.Scan(&amount)

	if ok {
		if amount > balance {
			fmt.Println("Ошибка: недостаточно средств на кошельке.")
		} else {
			Database[name] = balance - amount
			fmt.Printf("Привет, %v! Вы сняли сумму: %v. Остаток на кошельке: %v\n", name, balance, Database[name])
		}
	} else {
		fmt.Println("Ошибка: такого клиента не существует!")
	}
}
