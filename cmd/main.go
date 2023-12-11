package main

import (
	"bankCLI/pkg"
	"fmt"
)

func main() {
	var choice int
	fmt.Println("\033[2J")
	fmt.Printf("Вас приветствует Khayrakoff Bank!\nВыберите, что вам нужно:\n\n")
	for {
		fmt.Println("1. Добавить клиента")
		fmt.Println("2. Пополнить баланс клиента")
		fmt.Println("3. Снять деньги клиента")
		fmt.Println("4. Показать баланс клиента")
		fmt.Println("5. Перевести деньги другому клиенту")
		fmt.Println("0. Выйти")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			{
				fmt.Println("Программа завершена!")
				return
			}
		case 1:
			{
				pkg.RegisterClient()
			}
		case 2:
			{
				pkg.AddMoney()
			}
		case 3:
			{
			pkg.WithdrawMoney()
			}
		case 4:
			{
				pkg.ShowClientsAccount()
			}
		case 5:
				pkg.TransferMoney()
		default:
			{
				fmt.Println("Ошибка. Вы ввели неверную цифру!")
				break
			}
		}
	}
}

// Дз
/*
1. Создать клиента
2. Пополнить баланс клиента
3. Показать баланс клиента
4. Снять деньги клиента
*/
