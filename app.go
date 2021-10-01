package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var value string
	fmt.Println("Введите пример в формате a+b, доступны операторы: +, -, *, / ")
	fmt.Scan(&value)
	operator, err := checkOperator(value)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	arrNumbers, err := operands(value)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err := calc(arrNumbers, operator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Результат: %s=%.2f", value, result)
}

//калькулятор
func calc(num []float64, t string) (result float64, err error) {
	if len(num) != 2 {
		err = fmt.Errorf("Должно быть два операнда")
		return
	}

	ch := make(chan float64)
	switch t {
	case "+":
		go addition(num, ch)
	case "-":
		go substraction(num, ch)
	case "*":
		go multiplication(num, ch)
	case "/":
		go division(num, ch)
	}

	result = <-ch

	return
}

//Возврат операндов
func operands(str string) (result []float64, err error) {
	var num float64
	words := regexp.MustCompile("[+,*,-,/]{1}").Split(str, -1)
	for _, v := range words {
		num, err = strconv.ParseFloat(v, 64)
		if err != nil {
			err = fmt.Errorf("Не является числом, ошибка: %s\n", err)
			return
		}
		result = append(result, num)
	}

	return
}

//Сложение
func addition(num []float64, ch chan float64) {
	ch <- num[0] + num[1]
}

//Вычитание
func substraction(num []float64, ch chan float64) {
	ch <- num[0] - num[1]
}

//Умножение
func multiplication(num []float64, ch chan float64) {
	ch <- num[0] * num[1]
}

//Деление
func division(num []float64, ch chan float64) {
	if num[1] == 0 {
		fmt.Println("Нельзя делить на ноль")
		os.Exit(1)
	}

	ch <- num[0] / num[1]
}

//проверка количества операторов, возврат текущего оператора
func checkOperator(str string) (operator string, err error) {
	countOperator := 0
	operators := [...]string{"+", "-", "*", "/"}
	for _, v := range operators {
		countOperator += strings.Count(str, v)
		if countOperator == 1 && operator == "" {
			operator = v
		}
	}

	if countOperator != 1 {
		err = fmt.Errorf("Недопустимое количество операторов:%d. Должен быть один оператор", countOperator)
		return
	}

	return
}
