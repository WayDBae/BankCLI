package main

import (
	"fmt"
	"bankCLI/pkg"
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
				pkg.CreateClient()
			}
		case 2:
			{
				var clientName string
				var amount float64
				fmt.Println("Введите имя клиента:")
				fmt.Scan(&clientName)

				if _, exists := pkg.Database[clientName]; !exists {
					fmt.Println("Ошибка: клиент с таким именем не существует!")
					continue
				}
				pkg.AddMoney(clientName, amount)
			}
		case 3:
			{
				var clientName string
				var amount float64
				fmt.Println("Введите имя клиента:")
				fmt.Scan(&clientName)

				if _, exists := pkg.Database[clientName]; !exists {
					fmt.Println("Ошибка: клиент с таким именем не существует!")
					continue
				}

				pkg.WithdrawMoney(clientName, amount)
			}
		case 4: {
			pkg.ShowAllClients()
		}
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

