package lab2

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type ComputeHandler struct {
	YourPath string
}

func (ch *ComputeHandler) Compute(input, pkey, wkey string) string {
	err := ""
	res := 0

	if pkey != "" {
		pr := filepath.Join(ch.YourPath, pkey)
		data, merr1 := os.ReadFile(pr)
		if merr1 != nil {
			err = "Помилка читання файлу"
		} else {
			res, err = PostfixFunc(string(data))
			if err == "" {
				if wkey != "" {
					pw := filepath.Join(ch.YourPath, wkey)
					merr2 := os.WriteFile(pw, []byte(strconv.Itoa(res)), 0666)
					if merr2 != nil {
						err = "Помилка запису"
					}
				} else {
					fmt.Println("Відповідь: ", res)
				}
			}
		}
	} else {
		if input != "" {
			res, err = PostfixFunc(input)
			if err == "" {
				if wkey != "" {
					pw := filepath.Join(ch.YourPath, wkey)
					merr2 := os.WriteFile(pw, []byte(strconv.Itoa(res)), 0666)
					if merr2 != nil {
						err = "Помилка запису"
					}
				} else {
					fmt.Println("Відповідь: ", res)
				}
			}
		} else {
			err = "Неправильно ведені данні"
		}
	}
	return err
}
