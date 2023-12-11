package pkg

import "fmt"

func CreateClient() {
	fmt.Println("\033[2J")
	var name string
	var amount float64 = 0.0

	fmt.Println("Введите ваше имя: ")
	fmt.Scan(&name)
	if _, ok := Database[name]; !ok {
		fmt.Println("Ошибка: клиент с таким именем уже существует!")
		return
	}
	Database[name] = amount
	fmt.Println("Клиент успешно добавлен.")
	fmt.Println("________________________")
}
