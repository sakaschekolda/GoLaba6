1. Создание и запуск горутин: • Напишите программу, которая параллельно выполняет три функции (например, расчёт факториала, генерация случайных чисел и вычисление суммы числового ряда). • Каждая функция должна выполняться в своей горутине. • Добавьте использование time.Sleep() для имитации задержек и продемонстрируйте параллельное выполнение.
```
  package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для расчёта факториала числа n
func factorial(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		time.Sleep(100 * time.Millisecond) // Имитируем задержку
	}
	fmt.Printf("factorial %d = %d\n", n, result)
}

// Функция для генерации случайных чисел
func generateRandomNumbers(count int) {
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		fmt.Printf("Random number: %d\n", num)
		time.Sleep(150 * time.Millisecond) // Имитируем задержку
	}
}

// Функция для вычисления суммы ряда чисел от 1 до n
func sumSeries(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		time.Sleep(200 * time.Millisecond) // Имитируем задержку
	}
	fmt.Printf("Sum of a sequence from 1 to %d = %d\n", n, sum)
}

func main() {
	// Запуск трёх функций в горутинах
	go factorial(5)
	go generateRandomNumbers(5)
	go sumSeries(5)

	// Даем время горутинам завершить выполнение
	time.Sleep(3 * time.Second)
	fmt.Println("All GoRoutines done")
}
```
![{EAB1758A-5166-4416-B8B1-8B138691FD15}](https://github.com/user-attachments/assets/7f817e88-d5c5-480f-bd95-ed9db9e9351e)

2. Использование каналов для передачи данных: • Реализуйте приложение, в котором одна горутина генерирует последовательность чисел (например, первые 10 чисел Фибоначчи) и отправляет их в канал. • Другая горутина должна считывать данные из канала и выводить их на экран. • Добавьте блокировку чтения из канала с помощью close() и объясните её роль.
```
  package main

import (
	"fmt"
	"time"
)

// Функция для генерации первых n чисел Фибоначчи и отправки их в канал
func generateFibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(100 * time.Millisecond) // Имитируем задержку
	}
	close(ch) // Закрываем канал, когда все данные отправлены
}

// Функция для чтения данных из канала и вывода их на экран
func printFibonacci(ch chan int) {
	// Используем цикл для чтения данных из канала
	// Канал автоматически завершает цикл, когда он закрыт
	for num := range ch {
		fmt.Println(num)
		time.Sleep(150 * time.Millisecond) // Имитируем задержку
	}
}

func main() {
	// Создаём канал для передачи целых чисел
	fibChannel := make(chan int)

	// Запускаем горутины для генерации и вывода чисел
	go generateFibonacci(10, fibChannel)
	go printFibonacci(fibChannel)

	// Даем время горутинам завершить выполнение
	time.Sleep(2 * time.Second)
	fmt.Println("All GoRoutines done")
}
```
![{CD798485-0AD6-4799-872B-8A560F938C8B}](https://github.com/user-attachments/assets/cc46dffe-1731-470d-8c18-aacc6b877b46)

3. Применение select для управления каналами: • Создайте две горутины, одна из которых будет генерировать случайные числа, а другая — отправлять сообщения об их чётности/нечётности. • Используйте конструкцию select для приёма данных из обоих каналов и вывода результатов в консоль. • Продемонстрируйте, как select управляет многоканальными операциями.
```
  package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Горутина для генерации случайных чисел и отправки их в канал
func generateRandomNumbers(numbers chan int) {
	for {
		num := rand.Intn(100) // Генерируем случайное число от 0 до 99
		numbers <- num
		time.Sleep(200 * time.Millisecond) // Задержка для имитации работы
	}
}

// Горутина для проверки чётности/нечётности числа и отправки сообщения в другой канал
func checkEvenOdd(numbers chan int, messages chan string) {
	for {
		num := <-numbers // Получаем число из канала
		if num%2 == 0 {
			messages <- fmt.Sprintf("Number %d — is even", num)
		} else {
			messages <- fmt.Sprintf("Number %d — is odd", num)
		}
		time.Sleep(100 * time.Millisecond) // Задержка для имитации работы
	}
}

func main() {
	// Создаём два канала
	numbers := make(chan int)
	messages := make(chan string)

	// Запускаем горутины
	go generateRandomNumbers(numbers)
	go checkEvenOdd(numbers, messages)

	// Используем select для одновременной обработки данных из двух каналов
	for {
		select {
		case num := <-numbers:
			fmt.Printf("Generated number: %d\n", num)
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
```
![{1CD0972F-EF82-4884-B9DA-4484F6AEF5DE}](https://github.com/user-attachments/assets/5e757802-3c38-45c2-bf50-f5113eb34354)

4. Синхронизация с помощью мьютексов: • Реализуйте программу, в которой несколько горутин увеличивают общую переменную-счётчик. • Используйте мьютексы (sync.Mutex) для предотвращения гонки данных. • Включите и выключите мьютексы, чтобы увидеть разницу в работе программы.
```
  package main

import (
	"fmt"
	"sync"
)

// Общая переменная-счётчик
var counter = 0

// Мьютекс для синхронизации доступа к счётчику
var mutex sync.Mutex

// Функция, которая увеличивает счётчик
func increment(wg *sync.WaitGroup, useMutex bool) {
	defer wg.Done() // Сообщаем WaitGroup, что горутина завершена

	for i := 0; i < 1000; i++ {
		if useMutex {
			// Блокируем мьютекс перед изменением счётчика
			mutex.Lock()
		}

		// Увеличиваем счётчик
		counter++

		if useMutex {
			// Разблокируем мьютекс после изменения счётчика
			mutex.Unlock()
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Запуск горутин с использованием мьютекса
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment(&wg, true) // Включён мьютекс
	}
	wg.Wait()

	fmt.Printf("Counter with Mutex: %d\n", counter)

	// Обнуляем счётчик
	counter = 0

	// Запуск горутин без использования мьютекса
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment(&wg, false) // Мьютекс выключен
	}
	wg.Wait()

	fmt.Printf("Counter without Mutex: %d\n", counter)
}
```
![{3487C208-585B-4855-A07B-4C5520D297D0}](https://github.com/user-attachments/assets/0d0e4d93-4eb2-4ac2-bf75-c6a6229de731)

5. Разработка многопоточного калькулятора: • Напишите многопоточный калькулятор, который одновременно может обрабатывать запросы на выполнение простых операций (+, -, *, /). • Используйте каналы для отправки запросов и возврата результатов. • Организуйте взаимодействие между клиентскими запросами и серверной частью калькулятора с помощью горутин.
```
  package main

import (
	"fmt"
	"log"
	"sync"
)

// Структура для запроса на выполнение операции
type CalculationRequest struct {
	A, B   float64      // Операнды
	Op     string       // Операция: "+", "-", "*", "/"
	Result chan float64 // Канал для возврата результата
	Error  chan error   // Канал для возврата ошибки (например, деление на 0)
}

// Серверная часть калькулятора, которая выполняет операции
func calculator(requests chan CalculationRequest, wg *sync.WaitGroup) {
	defer wg.Done() // Уведомляем о завершении горутины

	for req := range requests {
		var res float64
		var err error

		// Выполнение операции на основе типа
		switch req.Op {
		case "+":
			res = req.A + req.B
		case "-":
			res = req.A - req.B
		case "*":
			res = req.A * req.B
		case "/":
			if req.B == 0 {
				err = fmt.Errorf("dividing by zero")
			} else {
				res = req.A / req.B
			}
		default:
			err = fmt.Errorf("unknown operation: %s", req.Op)
		}

		// Отправляем результат или ошибку в соответствующий канал
		if err != nil {
			req.Error <- err
		} else {
			req.Result <- res
		}
	}
}

func main() {
	// Создаём канал для отправки запросов
	requests := make(chan CalculationRequest)

	// WaitGroup для синхронизации завершения работы калькулятора и клиентских горутин
	var wg sync.WaitGroup

	// Запускаем серверную часть калькулятора в отдельной горутине
	wg.Add(1)
	go calculator(requests, &wg)

	// Создаём клиентские запросы
	operations := []CalculationRequest{
		{A: 10, B: 20, Op: "+", Result: make(chan float64), Error: make(chan error)},
		{A: 15, B: 5, Op: "-", Result: make(chan float64), Error: make(chan error)},
		{A: 7, B: 3, Op: "*", Result: make(chan float64), Error: make(chan error)},
		{A: 40, B: 0, Op: "/", Result: make(chan float64), Error: make(chan error)}, // Ошибка деления на 0
		{A: 16, B: 4, Op: "/", Result: make(chan float64), Error: make(chan error)},
	}

	// Используем WaitGroup для клиентских горутин
	var clientWg sync.WaitGroup

	// Запуск клиентских запросов в отдельных горутинах
	for _, op := range operations {
		clientWg.Add(1)
		go func(op CalculationRequest) {
			defer clientWg.Done()

			// Отправляем запрос на выполнение
			requests <- op

			// Ожидаем либо результат, либо ошибку
			select {
			case res := <-op.Result:
				fmt.Printf("Result: %f %s %f = %f\n", op.A, op.Op, op.B, res)
			case err := <-op.Error:
				log.Printf("Error %f %s %f: %s\n", op.A, op.Op, op.B, err)
			}
		}(op)
	}

	// Ожидаем завершения всех клиентских горутин
	clientWg.Wait()

	// Закрываем канал с запросами после того, как все клиентские горутины завершились
	close(requests)

	// Ожидаем завершения работы калькулятора
	wg.Wait()

	fmt.Println("All operations are done")
}
```
![{934C7DE6-3F01-4CAB-B913-9268E30B81A3}](https://github.com/user-attachments/assets/5b7d758a-411f-45aa-8b91-74c806a5c070)
6. Создание пула воркеров: • Реализуйте пул воркеров, обрабатывающих задачи (например, чтение строк из файла и их реверсирование). • Количество воркеров задаётся пользователем. • Распределение задач и сбор результатов осуществляется через каналы. • Выведите результаты работы воркеров в итоговый файл или в консоль.
Go
```
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
```
![{AC72F897-EB54-4EBB-A8B6-EA56B9046930}](https://github.com/user-attachments/assets/a40d5ee9-8c8a-4a84-918a-6930df690945)
input.txt
```
  Are we near the end, love?
  Was this just pretend love?
  I guess it's all downhill from here
' Cause every word I speak, girl, you don't hear
  Truth between the lies, gotta face our fears
```
output.txt
```
  ?evol ,dne eht raen ew erA
  ?evol dneterp tsuj siht saW
  ereh morf llihnwod lla s'ti sseug I
  raeh t'nod uoy ,lrig ,kaeps I drow yreve esuaC'
  sraef ruo ecaf attog ,seil eht neewteb hturT
```
