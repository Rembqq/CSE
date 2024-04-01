package lab2

import (
	"strconv"
)

func opNum(i int, a [15]int) (int, int, int, [15]int) {
	i--
	n2 := a[i]
	a[i] = 0
	i--
	n1 := a[i]
	return n1, n2, i, a
}

func opNumRes(i int, r int, a [15]int) (int, [15]int) {
	a[i] = r
	i++
	return i, a
}

func convText(i int, a [15]int, tN string) (int, [15]int, string, string) {
	conNum, err1 := strconv.Atoi(tN)
	e := ""
	if err1 != nil {
		e = "Помилка конвертації тексту"
	} else {
		a[i] = conNum
		i++
		tN = ""
	}
	return i, a, tN, e
}

func PostfixFunc(input string) (int, string) {

	textNum := ""
	var numArray [15]int
	indexArray := 1
	var num1 int
	var num2 int
	var res int
	s1 := '+'
	s2 := '-'
	s3 := '*'
	s4 := '/'
	s5 := '^'
	err := ""

	if input != "" {
		numArray[0] = 0
		for c := range input {
			if err == "" {
				if input[c] != ' ' {
					if input[c] != byte(s1) && input[c] != byte(s2) && input[c] != byte(s3) && input[c] != byte(s4) && input[c] != byte(s5) {
						switch input[c] {
						case '0':
							textNum += "0"
						case '1':
							textNum += "1"
						case '2':
							textNum += "2"
						case '3':
							textNum += "3"
						case '4':
							textNum += "4"
						case '5':
							textNum += "5"
						case '6':
							textNum += "6"
						case '7':
							textNum += "7"
						case '8':
							textNum += "8"
						case '9':
							textNum += "9"
						default:
							err = "У виразі присутні недопустимі символи"
						}
					} else {
						if textNum != "" {
							indexArray, numArray, textNum, err = convText(indexArray, numArray, textNum)
						}
						switch input[c] {
						case byte(s1):
							num1, num2, indexArray, numArray = opNum(indexArray, numArray)
							res = num1 + num2
							indexArray, numArray = opNumRes(indexArray, res, numArray)
						case byte(s2):
							num1, num2, indexArray, numArray = opNum(indexArray, numArray)
							res = num1 - num2
							indexArray, numArray = opNumRes(indexArray, res, numArray)
						case byte(s3):
							num1, num2, indexArray, numArray = opNum(indexArray, numArray)
							res = num1 * num2
							indexArray, numArray = opNumRes(indexArray, res, numArray)
						case byte(s4):
							num1, num2, indexArray, numArray = opNum(indexArray, numArray)
							res = num1 / num2
							indexArray, numArray = opNumRes(indexArray, res, numArray)
						case byte(s5):
							num1, num2, indexArray, numArray = opNum(indexArray, numArray)
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
							indexArray, numArray = opNumRes(indexArray, res, numArray)
						default:
							err = "У виразі присутні недопустимі оператори"

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
