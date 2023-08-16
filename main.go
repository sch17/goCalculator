package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("КАЛЬКУЛЯТОР")
	fmt.Println("Введите арифметическую задачу! (Примеры: 1 + 2, 4-3, 5 * 6, 8/7, X - IX)")

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		res, err := calc(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	}
}

func convertToDecimal(value string, valueType string) (int, error) {
	if valueType == "decimal" {
		return strconv.Atoi(value)
	} else if valueType == "roman" {
		return getDecFromRom(value), nil
	} else {
		return 0, errors.New("неподдерживаемый тип числа")
	}
}

func calc(s string) (string, error) {
	parts := strings.Fields(s)
	if len(parts) != 3 {
		return "", errors.New("неправильно введена арифметическая задача")
	}

	val1 := parts[0]
	act := parts[1]
	val2 := parts[2]

	// Проверяем цифры римские или арабские
	valueType := isDecimalOrRoman(val1)
	value2Type := isDecimalOrRoman(val2)
	if valueType == "decimal" && value2Type == "roman" || valueType == "roman" && value2Type == "decimal" {
		return "", errors.New("калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	}

	// Преобразуем строки в тип int и переводим в десятичную систему при необходимости
	one, err := convertToDecimal(val1, valueType)
	if err != nil {
		return "", err
	}
	two, err := convertToDecimal(val2, value2Type)
	if err != nil {
		return "", err
	}

	// Арифметические действия выполняет метод action. Передаем все параметры, метод возвращает результат
	result, err := action(one, two, act)
	if err != nil {
		return "", err
	}

	if valueType == "roman" && result < 1 {
		return "", errors.New("в римской системе нет отрицательных чисел")
	} else if valueType == "roman" {
		return getRomFromDec(result), nil
	} else {
		return strconv.Itoa(result), nil
	}
}

// Метод, выполняющий все арифметические действия
func action(one, two int, act string) (int, error) {
	switch act {
	case "+":
		return one + two, nil
	case "-":
		return one - two, nil
	case "*":
		return one * two, nil
	case "/":
		if two == 0 {
			return 0, errors.New("на ноль делить нельзя")
		}
		return int(one / two), nil
	default:
		return 0, errors.New("произошла ошибка во время выполнения арифметического действия")
	}
}

// Функция проверяет является число римским или арабским
func isDecimalOrRoman(value string) string {
	if _, err := strconv.Atoi(value); err == nil {
		return "decimal"
	} else {
		return "roman"
	}
}

// Метод преобразует римские цифры в арабские (от 1 до 10)
func getDecFromRom(value string) int {
	switch value {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		return 0
	}
}

// Метод преобразует арабские цифры в римские (от 1 до 100)
func getRomFromDec(number int) string {

	var result string
	for number != 0 {
		if number == 100 {
			number -= 100
			result = result + "C"
		} else if number >= 90 && number < 100 {
			number -= 90
			result = result + "XC"
		} else if number >= 50 && number < 90 {
			number -= 50
			result = result + "L"
		} else if number >= 40 && number < 50 {
			number -= 40
			result = result + "XL"
		} else if number >= 10 && number < 40 {
			number -= 10
			result = result + "X"
		} else if number == 9 {
			number -= 9
			result = result + "IX"
		} else if number >= 5 && number < 9 {
			number -= 5
			result = result + "V"
		} else if number == 4 {
			number -= 4
			result = result + "IV"
		} else {
			number -= 1
			result = result + "I"
		}
	}
	return result

}
