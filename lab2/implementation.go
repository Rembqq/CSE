package lab2

import (
	"fmt"
	"strconv"
)

var textNum string = ""
var numArray [15]int
var indexArray int = 0
var num1 int
var num2 int
var res int
var s1 rune = '+'
var s2 rune = '-'
var s3 rune = '*'
var s4 rune = '/'
var s5 rune = '^'
var err string = ""

func opNum() {
	indexArray--
	num2 = numArray[indexArray]
	numArray[indexArray] = 0
	indexArray--
	num1 = numArray[indexArray]
}

func opNumRes() {
	numArray[indexArray] = res
	indexArray++
}

func convText() {
	conNum, err1 := strconv.Atoi(textNum)
	if err1 != nil {
		err = "Помилка конвертації тексту"
	} else {
		numArray[indexArray] = conNum
		indexArray++
		textNum = ""
	}
}

func PostfixFunc(input string) (string, error) {
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
						convText()
						textNum = ""
					}
					switch input[c] {
					case byte(s1):
						opNum()
						res = num1 + num2
						opNumRes()
					case byte(s2):
						opNum()
						res = num1 - num2
						opNumRes()
					case byte(s3):
						opNum()
						res = num1 * num2
						opNumRes()
					case byte(s4):
						opNum()
						res = num1 / num2
						opNumRes()
					case byte(s5):
						opNum()
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
						opNumRes()
					default:
						err = "У виразі присутні недопустимі оператори"

					}
				}
			} else {
				if textNum != "" {
					convText()
				}
			}
		}
	}
	output := strconv.Itoa(numArray[0])
	return output, fmt.Errorf(err)
}
