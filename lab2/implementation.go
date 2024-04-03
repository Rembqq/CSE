package lab2

import (
	"strconv"
)

// takes last 2 nums from numArray
func opNum(index int, numArray [15]int) (int, int, int, [15]int) {
	index--
	n2 := numArray[index]
	numArray[index] = 0
	index--
	n1 := numArray[index]
	return n1, n2, index, numArray
}

// operator number result
func opNumRes(index int, result int, numArray [15]int) (int, [15]int) {
	numArray[index] = result
	index++
	return index, numArray
}

func convText(index int, numArray [15]int, textNum string) (int, [15]int, string, string) {
	conNum, err1 := strconv.Atoi(textNum)
	e := ""
	if err1 != nil {
		e = "Помилка конвертації тексту"
	} else {
		numArray[index] = conNum
		index++
		textNum = ""
	}
	return index, numArray, textNum, e
}

func PostfixResult(input string) (int, string) {

	textNum := ""
	var numArray [15]int
	indexArray := 1

	// Operands
	var num1 int
	var num2 int

	var res int
	err := ""

	// Operators
	opPlus := '+'
	opMinus := '-'
	opMultiplication := '*'
	opDivision := '/'
	opPowerOf := '^'

	if input != "" {
		// numArray[0] = 0
		for c := range input {
			if err == "" {
				if input[c] != ' ' {
					if input[c] != byte(opPlus) && input[c] != byte(opMinus) && input[c] != byte(opMultiplication) &&
						input[c] != byte(opDivision) && input[c] != byte(opPowerOf) {
						switch input[c] {
						case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
							textNum += string(input[c])
						default:
							err = "У виразі присутні недопустимі символи"
						}
					} else {
						if textNum != "" {
							indexArray, numArray, textNum, err = convText(indexArray, numArray, textNum)
						}

						num1, num2, indexArray, numArray = opNum(indexArray, numArray)

						switch input[c] {
						case byte(opPlus):
							res = num1 + num2
						case byte(opMinus):
							res = num1 - num2
						case byte(opMultiplication):
							res = num1 * num2
						case byte(opDivision):
							res = num1 / num2
						case byte(opPowerOf):
							// Math.pow is omitted because of its resource intensity
							if num2 == 0 {
								res = 1
							}
							if num2 >= 2 {
								res = num1
								for i := 2; i <= num2; i++ {
									res = res * num1
								}
							} else {
								res = num1
							}
						default:
							err = "У виразі присутні недопустимі оператори"
						}
						if err == "" {
							indexArray, numArray = opNumRes(indexArray, res, numArray)
						}
					}
				} else {
					if textNum != "" {
						indexArray, numArray, textNum, err = convText(indexArray, numArray, textNum)
					}
				}
			}
		}
		if err == "" && numArray[0] == 0 && numArray[2] == 0 && numArray[3] == 0 && numArray[4] == 0 {
			output := numArray[1]
			return output, err
		} else {
			if err == "" {
				err = "Невірне обчислення. Невірна кількість операторів або операндів або невірний запис виразу."
			}
			output := numArray[1]
			return output, err
		}
	} else {
		err = "Немає виразу для обчислення"
		return 0, err
	}
}
