package main

import "fmt"

// 声明浮点type的 类型
type ErrNegativeSqrt float64

// 类型ErrNegativeSqrt 实现了error接口
func (err ErrNegativeSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number:%f", err)
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		var num ErrNegativeSqrt = ErrNegativeSqrt(x)
		return float64(num), num
	}

	return x, nil
}

func init() {

	fmt.Println("===== exercise-errors.go =====")

	f, err := Sqrt(-1)

	fmt.Println(f, err)
}
