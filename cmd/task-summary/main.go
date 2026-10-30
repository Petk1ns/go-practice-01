package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		status = "Новая"
	case "in_progress":
		status = "В работе"
	case "done":
		status = "Завершена"
	default:
		status = "Неизвестный статус"
	}

	return status
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	var counter int = 0
	for _, status := range statuses {
		if status == target {
			counter += 1
		}
	}

	return counter
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d", counts["new"], counts["in_progress"], counts["done"])

}
