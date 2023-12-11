package pkg

import "fmt"

func ShowClientsAccount() {
	fmt.Println("\033[2J")
	var name string
	fmt.Println("Введите имя клиента:")
	fmt.Scan(&name)

	fmt.Println("Ваше имя, Sir | Ваш cash :D")

	balance, ok := Database[name]
	if !ok {
		fmt.Println("Ошибка! Данного пользователя не существует!")
		return
	}
	fmt.Printf("Баланс счёта %v является %v сомони\n", name, balance)
	fmt.Println("______________")
}
