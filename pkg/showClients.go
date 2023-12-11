package pkg

import "fmt"

func ShowAllClients() {
	fmt.Println("\033[2J")
	fmt.Println("Ваше имя, Sir | Ваш cash :D")
	for key, val := range Database {
		fmt.Println(key, "|", val, "сомони")
	}
	fmt.Println("______________")
}
