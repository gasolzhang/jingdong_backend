package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func demo() {
	fmt.Println("初始化项目")
	//` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
	fmt.Printf(`sds`)
	fmt.Printf("切片长度：%d 最后一个元素%d", len(slice), slice[len(slice)-1])
	m2 := map[string]string{

	}

	var m3 = make(map[string]string)
	m3["k1"] = "v1"
	fmt.Println(m2)
	fmt.Println(m3)

	for k, v := range m3 {
		fmt.Printf("k=%s v=%s", k, v)
	}

	s := []int8{1, 2, 3}
	fmt.Println(s)
	s[3] = 4

	var p = new(Person)
	p.name = "张三"
	p.age = 20
}

type Person struct {
	name string
	age  int
}

func (p *Person) say() {
	fmt.Println(p.name)
}

const (
	a = iota
	b
	c
)

var arr = [3]int8{1, 2, 3}

//自定判断长度
var arr2 = [...]int8{1, 2, 3}

var slice = []int{1, 2, 3}

var m = map[string]string{

}

var m1 = make(map[string]string)

type Human interface {
	Say(word string)
	Run()
}

type Man struct {
}

//实现接口的全部方法就是实现了这个接口
func (m *Man) Say(word string) {
}
func (m *Man) Run() {}

func fileDemo1() {
	f, err := os.Open("文件名")
	if err != nil {
		log.Fatalln("打开文件错误：", err.Error())
		return
	}
	fmt.Printf("文件内容%v", f)

	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		fmt.Println(str)
	}

	defer f.Close()
}

func fileDemo2() {
	//一次性将文件内容全部读取出来，仅适用于小文件
	content, err := ioutil.ReadFile("文件路径")
	if err != nil {
		fmt.Printf("读取错误:%v", err)
		return
	}
	fmt.Printf("读取的全部内容:%s", string(content))
}

func fileDemo3() {
	path := "E:/a.txt";
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	writter := bufio.NewWriter(f)
	writter.WriteString("写入的内容")
	//writter带缓存，所以要flush
	writter.Flush()

}

func CopyDemo() {
	srcFile, err := os.OpenFile("E:/a.txt", os.O_RDONLY, 666)
	if err != nil {
		log.Fatalln("打开源文件失败")
		return
	}

	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	destFile, err2 := os.OpenFile("E:/b.txt", os.O_WRONLY, 666)
	if err2 != nil {
		log.Fatalln("打开目标文件失败")
		return
	}

	defer destFile.Close()

	wrriter := bufio.NewWriter(destFile)

	io.Copy(wrriter, reader)
}