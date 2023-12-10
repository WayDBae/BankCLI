package main

import (
	"fmt"
)

var database = make(map[string]float64)

func CreateClient() {
	fmt.Println("\033[2J")
	var name string
	var amount float64 = 0.0

	fmt.Println("Введите ваше имя: ")
	fmt.Scan(&name)
	if _, exists := database[name]; exists {
		fmt.Println("Ошибка: клиент с таким именем уже существует!")
		return
	}
	database[name] = amount
	fmt.Println("Клиент успешно добавлен.")
	fmt.Println("________________________")
}

func AddMoney(name string, amount float64) {
	fmt.Println("\033[2J")
	var money float64
	fmt.Println("На сколько сомони вы хотите пополнить ваш кошелёк?")
	fmt.Scan(&money)
	if money < 0 {
		fmt.Println("Пополнение всегда больше нуля!")
		return
	}
	currentAmount := database[name]
	database[name] = currentAmount + money
	for ind, _ := range database {
		fmt.Printf("Привет, %v! Ваш кошелек успешно пополнен на сумму %v сомони\n", ind, money)
	}
}

func WithdrawMoney(name string, amount float64) {
	fmt.Println("\033[2J")
	var money float64
	fmt.Println("Какую сумму вы хотите снять с вашего кошелька?")
	fmt.Scan(&money)

	if currentAmount, exists := database[name]; exists {
		if money > currentAmount {
			fmt.Println("Ошибка: недостаточно средств на кошельке.")
		} else {
			database[name] = currentAmount - money
			fmt.Printf("Привет, %v! Вы сняли сумму: %v. Остаток на кошельке: %v\n", name, money, database[name])
		}
	} else {
		fmt.Println("Ошибка: такого клиента не существует!")
	}
}

func ShowAllClients() {
	fmt.Println("\033[2J")
	fmt.Println("Ваше имя, Sir | Ваш cash :D")
	for key, val := range database {
		fmt.Println(key, "|", val, "сомони")
	}
	fmt.Println("______________")
}
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
				CreateClient()
			}
		case 2:
			{
				var clientName string
				var amount float64
				fmt.Println("Введите имя клиента:")
				fmt.Scan(&clientName)

				if _, exists := database[clientName]; !exists {
					fmt.Println("Ошибка: клиент с таким именем не существует!")
					continue
				}
				AddMoney(clientName, amount)
			}
		case 3:
			{
				var clientName string
				var amount float64
				fmt.Println("Введите имя клиента:")
				fmt.Scan(&clientName)

				if _, exists := database[clientName]; !exists {
					fmt.Println("Ошибка: клиент с таким именем не существует!")
					continue
				}

				WithdrawMoney(clientName, amount)
			}
		case 4: {
			ShowAllClients()
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

