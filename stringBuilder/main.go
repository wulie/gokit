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
	for _, id := range ids {
		sb.WriteString(id)
	}
	fmt.Println(time.Since(t), len(sb.String()))
	//fmt.Println(sb.String())

}

//output
//3.082893072s 488890
//756.227µs 488890

//字符串拼接用strings.Builder
