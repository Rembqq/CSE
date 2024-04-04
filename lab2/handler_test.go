package lab2

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCompute(t *testing.T) {
	ch := ComputeHandler{YourPath: t.TempDir()}
	pr := filepath.Join(ch.YourPath, "input.txt")
	os.WriteFile(pr, []byte("4 3 -"), 0666) // створення файлу для читання у тимчасовій директорії
	for _, tc := range []struct {
		name  string
		input string
		pkey  string
		wkey  string
	}{
		{name: "1", input: "4 3 -", pkey: "", wkey: "result.txt"},
		{name: "2", input: "4 3 -", pkey: "", wkey: ""},
		{name: "3", input: "", pkey: "input.txt", wkey: ""},
		//{name: "4", input: "", pkey: "inp.txt", wkey: ""}, // неправильно вказаний файл для читання
		//{name: "5", input: "", pkey: "", wkey: ""},        // невказано виразу для обрахунку
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := ch.Compute(tc.input, tc.pkey, tc.wkey)
			if err != "" {
				t.Errorf("Помилка: %s", err)
			} else {
				fmt.Println("Виконання успішне")
			}
		})
	}
}
