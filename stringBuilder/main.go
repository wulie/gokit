package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	ids := make([]string, 0)
	for i := 0; i < 100000; i++ {
		ids = append(ids, strconv.Itoa(i))
	}

	idString := ""
	t := time.Now()
	for _, id := range ids {
		idString += id
	}
	fmt.Println(time.Since(t), len(idString))
	//fmt.Println(idString)
	t = time.Now()
	sb := strings.Builder{}
	sb.Grow(100000)
	for _, id := range ids {
		sb.WriteString(id)
	}
	fmt.Println(time.Since(t), len(sb.String()))
	//fmt.Println(sb.String())

	var v interface{}

	//v = make([]int32,0)

	v = append((v).([]int32), 34)
	v = append((v).([]int32), 34)
	v = append((v).([]int32), 34)
	v = append((v).([]int32), 34)
	v = append((v).([]int32), 34)
	v = append((v).([]int32), 34)
	v = append((v).([]int32), 34)

	fmt.Println(v)

}

//output
//3.082893072s 488890
//756.227µs 488890

//字符串拼接用strings.Builder strings.Builder 的使用场景一般是在循环中对字符串进行拼接，
//如果只是拼接两个或少数几个字符串的话，推荐使用 "+"操作符，例如: s := s1 + s2 + s3，
//该操作并非每个 + 操作符都计算一次长度，而是会首先计算三个字符串的总长度，然后分配对应的内存，再将三个字符串都拷贝到新申请的内存中去。
// strings.trim
