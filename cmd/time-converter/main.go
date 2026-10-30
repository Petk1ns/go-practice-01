package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")

	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	const minutes = 60
	hours := totalMinutes / minutes
	remain := totalMinutes % minutes
	fmt.Printf("Получено минут: %d \n", totalMinutes)
	fmt.Printf("Полных часов: %d \nОставшихся минут: %d", hours, remain)
}