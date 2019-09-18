package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("Rectangle Package initializied!")
}

//初始化包，多次导入只会被执行一次
//包可以有多个 init 函数,但是只会被执行一次

func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
