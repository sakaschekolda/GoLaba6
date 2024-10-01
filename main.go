package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Структура для задания, которая включает строку и её индекс для сохранения порядка
type Task struct {
	Index int
	Line  string
}

// Функция, выполняющая задачу (реверсирование строки)
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Воркеры, которые получают задачи и отправляют результаты
func worker(id int, tasks <-chan Task, results chan<- Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		// Обрабатываем задачу (реверсируем строку)
		fmt.Printf("Worker %d processing task %d\n", id, task.Index)
		task.Line = reverseString(task.Line)
		results <- task // Отправляем результат в канал результатов
	}
}

func main() {
	// Открытие исходного файла для чтения
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error has occured while oppening the file:", err)
		return
	}
	defer inputFile.Close()

	// Считываем строки из файла
	var tasks []Task
	scanner := bufio.NewScanner(inputFile)
	index := 0
	for scanner.Scan() {
		tasks = append(tasks, Task{Index: index, Line: scanner.Text()})
		index++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error has occured while reading the file:", err)
		return
	}

	// Пользователь задаёт количество воркеров
	var workerCount int
	fmt.Print("Enter amount of workers: ")
	fmt.Scan(&workerCount)

	// Канал для задач
	taskChan := make(chan Task, len(tasks))

	// Канал для результатов
	resultChan := make(chan Task, len(tasks))

	// WaitGroup для ожидания завершения работы всех воркеров
	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, taskChan, resultChan, &wg)
	}

	// Отправляем задачи в канал
	for _, task := range tasks {
		taskChan <- task
	}

	// Закрываем канал задач, т.к. все задачи отправлены
	close(taskChan)

	// Ожидаем завершения всех воркеров
	go func() {
		wg.Wait()
		close(resultChan) // Закрываем канал результатов после завершения воркеров
	}()

	// Собираем результаты
	results := make([]Task, len(tasks))
	for result := range resultChan {
		results[result.Index] = result
	}

	// Вывод результатов в консоль
	for _, result := range results {
		fmt.Printf("Reversed string (task %d): %s\n", result.Index, result.Line)
	}

	// Запись результатов в файл
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error has occured while creating the file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, result := range results {
		writer.WriteString(result.Line + "\n")
	}
	writer.Flush()

	fmt.Println("All tasks are done, all the results are in the output.txt")
}
