package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordsCounter int

func (w *WordsCounter) Write(p []byte) (int, error) {
	str := strings.NewReader(string(p)) //strings.NewReader函数实现了io.Reader接口，所以可以使用
	bs := bufio.NewScanner(str)         //NewScanner需要传入一个实现了Io.reader接口的参数
	bs.Split(bufio.ScanWords)           //bufio.ScanWords 是用于 Scanner类型的分隔函数,根据文档中提到的bufio.ScanWords顺着接口一层层往上找
	sum := 0
	for bs.Scan() {
		sum++
	}
	*w += WordsCounter(sum)
	return sum, nil
}

type LinesCouter int

func (l *LinesCouter) Write(p []byte) (int, error) {
	lines := strings.NewReader(string(p))
	bs := bufio.NewScanner(lines)
	bs.Split(bufio.ScanLines)
	sum := 0
	for bs.Scan() {
		sum++
	}
	*l += LinesCouter(sum)
	return sum, nil
}

//使用来自ByteCounter的思路， 实现一个针对对单词和行数的计数器。 你会发现 bufio.ScanWords非常的有用
func main() {
	var w WordsCounter
	/*strs := map[string]string{
		"a": "hello world you nice",
		"b": "hello world you",
	}
	for _, v := range strs {
		res, _ := w.Write([]byte(v))
		fmt.Println(res)
	}*/
	fmt.Fprintf(&w, "hello world you nice")
	fmt.Println(w)

	/*lines := "a\nb\nc"
	line1 := "aa\nbb\ncc"
	var l LinesCouter
	l.Write([]byte(lines))
	l.Write([]byte(line1))
	fmt.Println(l)*/
	var l LinesCouter
	fmt.Fprintf(&l, "aa\nbb\ncc")
	fmt.Println(l)
}
