package main

import "fmt"

func main() {
	const course = "Бэкэнд разработка"
	var name string = "Denis"
	group := "Учебная-04"
	var count int = 1
	interest := true
	fmt.Println("Участник:", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Println("Завершено работ:", count)
	fmt.Printf("Интререс к бэкэнду %t\n", interest)

}
