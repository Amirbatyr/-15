
package main

import "fmt"

func main() {
	// Задание 1
	var numbers []int
	var n int

	fmt.Println("Задание 1. Введите числа (0 для остановки):")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		numbers = append(numbers, n)
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Сумма:", sum)

	// Задание 2
	var values []int

	fmt.Println("Задание 2. Введите числа (0 для остановки):")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		values = append(values, n)
	}

	var evenValues []int
	for _, num := range values {
		if num%2 == 0 {
			evenValues = append(evenValues, num)
		}
	}

	fmt.Println("Исходный слайс:", values)
	fmt.Println("Чётные числа:", evenValues)

	// Задание 3
	var data []int

	fmt.Println("Задание 3. Введите числа (0 для остановки):")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		data = append(data, n)
	}

	if len(data) > 2 {
		data = append(data[:2], data[3:]...)
	}
	fmt.Println("Результат:", data)

	// Задание 4
	var temps []int

	fmt.Println("Задание 4. Введите температуры (0 для остановки):")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		temps = append(temps, n)
	}

	if len(temps) > 0 {
		max := temps[0]
		min := temps[0]

		for _, temp := range temps {
			if temp > max {
				max = temp
			}
			if temp < min {
				min = temp
			}
		}

		fmt.Println("Максимум:", max)
		fmt.Println("Минимум:", min)
	} else {
		fmt.Println("Список пуст")
	}

	// Задание 5
	var words []string
	var word string

	fmt.Println("Задание 5. Введите слова (stop для остановки):")
	for {
		fmt.Scan(&word)
		if word == "stop" {
			break
		}
		words = append(words, word)
	}

	var reversed []string
	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i])
	}

	fmt.Println("Перевёрнутый список:", reversed)

	// Задание 6
	var nums []int

	fmt.Println("Задание 6. Введите числа (0 для остановки):")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		nums = append(nums, n)
	}

	sorted := true
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			sorted = false
			break
		}
	}

	fmt.Println(sorted)

	// Задание 7
	var myWishList []string
	var friendsWishList []string

	fmt.Println("Задание 7. Введите свои подарки (stop для остановки):")
	for {
		fmt.Scan(&word)
		if word == "stop" {
			break
		}
		myWishList = append(myWishList, word)
	}

	fmt.Println("Введите подарки друга (stop для остановки):")
	for {
		fmt.Scan(&word)
		if word == "stop" {
			break
		}
		friendsWishList = append(friendsWishList, word)
	}

	var resultList []string

	for _, gift := range myWishList {
		resultList = append(resultList, gift)
	}
	for _, gift := range friendsWishList {
		resultList = append(resultList, gift)
	}

	fmt.Println("Общий список:", resultList)

	// Задание 8
	toyList := []string{"Car", "Doll", "Ball"}
	testToyList := make([]string, len(toyList))
	copy(testToyList, toyList)

	testToyList[1] = "Boat"

	fmt.Println("Исходный список:", toyList)
	fmt.Println("Копия списка:", testToyList)
}