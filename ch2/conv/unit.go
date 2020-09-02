package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch2/conv/unitconv"
	"os"
	"strconv"
)

//运行  go run  ch2/conv/unit.go  42
func main() {
	flag := 0 //定义一个标志位
	for _, v := range os.Args[1:] {
		if v != "" {
			flag = 1 //标志位
		}
		conv(v)
		//fmt.Println(v)
	}
	//判断是否有命令行传参
	if flag != 0 {
		return
	}
	//从标准输入读
	fmt.Print("请输入数字:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	unitStr := input.Text()
	conv(unitStr)
}

/*
转换函数
*/
func conv(pa string) {
	p, _ := strconv.ParseFloat(pa, 64)
	//显式类型转换
	pMeter := unitconv.Meter(p)
	pFoot := unitconv.Foot(p)

	a := unitconv.MtoF(pMeter)
	b := unitconv.FtoM(pFoot)
	fmt.Printf("%s = %s ; %s = %s \n", pMeter, a, pFoot, b)

}

// %!s(unitconv.Meter=42) = %!s(unitconv.Foot=137.79527559055117) ; %!s(unitconv.Foot=42) = %!s(unitconv.Meter=12.8016)
// 42m = 137.79527559055117ft ; 42ft = 12.8016m
