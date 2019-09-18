package main

import (
	"fmt"
	"geometry/rectangle" //自定义包导入,如果想在包外访问一个函数，它应该首字母大写
	"log"
)

//如果导入包不使用，程序会抛出异常
//但是也是会用空白标识符代替

var rectLen, rectWide float64 = 5, 6

//初始化判断
func init() {
	fmt.Println("The main function initializied!")
	if rectLen < 0 {
		log.Fatal("Length is less than 0")
	}
	if rectWide < 0 {
		log.Fatal("Wide is less than 0")
	}
}

func main() {
	fmt.Println("Geometrical shape properties!")
	fmt.Printf("Area of rectangle is %.2f!", rectangle.Area(rectLen, rectWide))
	fmt.Printf("Diagonal of rectangle is %.2f!", rectangle.Diagonal(rectLen, rectWide))
}

//首先初始化被导入的包。因此，首先初始化了 rectangle 包。
//接着初始化了包级别的变量 rectLen 和 rectWidth。
//调用 init 函数。
//最后调用 main 函数。

//调试阶段，导入包调试，开始完了之后，打开错误屏蔽器
//package main
//
//import (
//"geometry/rectangle"
//)
//
//var _ = rectangle.Area // 错误屏蔽器
//
//func main() {
//
//}

//package main
//
//import (
//    _ "geometry/rectangle"  //有时候我们导入一个包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量
//)
//func main() {
//
//}

//
