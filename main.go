package main

import (
	"fmt"
)

func main() {
	fmt.Println("Банкомат.")

	//переменная создана для удобной возможности изменить доступный номинал купюр.
	cashPar := 100
	fmt.Println("Доступный номинал купюр:", cashPar)

	//переменная создана для удобной возможности изменить лимит суммы снятия.
	cashLimit := 100000
	fmt.Println("Доступный лимит для снятия:", cashLimit)

	fmt.Println("Введите желаемую сумму снятия:")
	var cashRequired int
	fmt.Scan(&cashRequired)

	if cashRequired > cashLimit {
		fmt.Println("Введённая сумма недоступна для снятия.")
		fmt.Println("Укажите сумму с учётом доступного лимита.")
	} else if cashRequired%cashPar != 0 {
		fmt.Println("Отсутствует подходящий номинал для выдачи введённой суммы.")
		fmt.Println("Укажите сумму с учётом доступного номинала купюр.")
	} else if cashRequired > 0 {
		fmt.Println("Операция выполнена успешно.")
		fmt.Print("Вы сняли: ", cashRequired, "руб.\n")
	} else {
		fmt.Println("Введённая сумма некорректна.")
		fmt.Println("Укажите другую сумму.")
	}
}
