package lab2

import (
	"fmt"
	"testing"
)

func TestPostfixResult(t *testing.T) {
	for _, tc := range []struct {
		mathOp string
		resOp  int
	}{
		{mathOp: "4 3 - 2 * 3 2 ^ + 10 1 * - 3 +", resOp: 4},
		{mathOp: "4 3 - 2 * 3 2 ^ + 10 1 * -", resOp: 1},
		{mathOp: "4 3 - 2 * 3 2 ^ +", resOp: 11},
		{mathOp: "4 3 - 2 *", resOp: 2},
		{mathOp: "4 3 -", resOp: 1},
		//{mathOp: "4 3 - 2 * 3 2 ^ + 10 1 * - 3 ", resOp: 4}, // невірний запис виразу. Тест покаже помилку
		//{mathOp: "4 3 - 2 * 3 ^ +", resOp: 11},              // невірний запис виразу. Тест покаже помилку
		//{mathOp: "", resOp: 0},                              // відсутність виразу виразу. Тест покаже помилку
	} {
		t.Run(tc.mathOp, func(t *testing.T) {
			if got, textErr := PostfixResult(tc.mathOp); textErr != "" {
				t.Errorf("Have error: %s. ResOp: %d, got: %d", textErr, tc.resOp, got)
			} else {
				fmt.Println("ResOp:", tc.resOp, "got:", got)
			}
		})
	}
}

func BenchmarkPostfixResult(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			if _, textErr := PostfixResult("4 3 -"); textErr != "" {
				b.Errorf("Have error: %s", textErr)
			}
		}
	})
	b.Run("medium", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			if _, textErr := PostfixResult("4 3 - 2 * 3 2 ^ +"); textErr != "" {
				b.Errorf("Have error: %s", textErr)
			}
		}
	})
	b.Run("large", func(b *testing.B) {
		for k := 0; k < b.N; k++ {
			if _, textErr := PostfixResult("4 3 - 2 * 3 2 ^ + 10 1 * - 3 +"); textErr != "" {
				b.Errorf("Have error: %s", textErr)
			}
		}
	})
}

func ExamplePostfixResult() {
	res1, _ := PostfixResult("4 3 - 2 * 3 2 ^ + 10 1 * - 3 +")
	fmt.Println(res1)
	res2, _ := PostfixResult("4 3 - 2 * 3 2 ^ +")
	fmt.Println(res2)
	res3, _ := PostfixResult("4 3 -")
	fmt.Println(res3)
	// Output:
	// 4
	// 11
	// 1
}
