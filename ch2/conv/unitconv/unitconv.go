package unitconv

import (
	"fmt"
	"strconv"
)

//定义类型
type Foot float64
type Meter float64

/*
米转换成英尺
*/
func MtoF(m Meter) Foot {
	//四舍五入保留两位小数
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", m/0.3048), 64)
	return Foot(v)
}

/*
英尺转换成米
*/
func FtoM(f Foot) Meter {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f*0.3048), 64)
	return Meter(v)
}

//  许多类型都会定义一个String方法，因为当使用fmt包的打印方法时，将会优先使用该类型对应的String 方法返回的结果打印
/*
类型的String方法
*/
func (f Foot) String() string {
	return fmt.Sprintf("%gft", f) //%g  以%e或者%f表示的浮点数或者复数，任何一个都以最为紧凑的方式输出
}

/*
类型的String方法
*/
func (m Meter) String() string {
	return fmt.Sprintf("%gm", m) //%g  以%e或者%f表示的浮点数或者复数，任何一个都以最为紧凑的方式输出
}
