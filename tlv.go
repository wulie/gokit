package gokit

import (
	"bytes"
	"encoding/binary"
)

/*
	TLV是一种可变格式，意思就是：
	Type类型， Lenght长度，Value值；
	Type和Length的长度固定，一般那是2、4个字节（这里统一采用4个字节）；
	Value的长度有Length指定；

	编码方法：
	1.       将类型type用htonl转换为网络字节顺序，指针偏移+4
	2.       将长度length用htonl转换为网络字节顺序，指针偏移+4
	3.       若值value数据类型为int、char、short，则将其转换为网络字节顺序，指针偏移+4；
             若值为字符串类型，写进后，指针偏移+length
	……继续处理后面的tlv；

	解码方法：
	1.       读取type 用ntohl转换为主机字节序得到类型，指针偏移+4
	2.       读取lengh用ntohl转换为主机字节序得到长度；指针偏移+4
	3.       根据得到的长度读取value，若value数据类型为int、char、short，用ntohl转换为主机字节序，指针偏移+4；
             若value数据类型为字符串类型，指针偏移+length
	……继续处理后面的tlv；

	类型(Type)字段是关于标签和编码格式的信息；
	长度 (Length)字段定义数值的长度；
	内容(Value)字段表示实际的数值。

	因此，一个编码值又称TLV(Type,Length,Value)三元组。
    编码可以是基本型或结构型，如果它表示一个简单类型的、完整的显式值，那么编码就是基本型 (primitive)；
    如果它表示的值具有嵌套结构，那么编码就是结构型 (constructed)。


*/
func TLVEncode(tag int32, data string) ([]byte, error) {
	buf := new(bytes.Buffer)
	// 写入TAG
	if err := binary.Write(buf, binary.BigEndian, tag); err != nil {
		return nil, err
	}
	dataBuf := []byte(data)
	// 写入length
	if err := binary.Write(buf, binary.BigEndian, int32(len(dataBuf))); err != nil {
		return nil, err
	}
	// 写入数据
	if err := binary.Write(buf, binary.BigEndian, dataBuf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func TLVDecode(b []byte) (int32, string, error) {
	buf := bytes.NewBuffer(b)
	var tag, length int32
	// 读取tag
	if err := binary.Read(buf, binary.BigEndian, &tag); err != nil {
		return 0, "", err
	}
	// 读取length
	if err := binary.Read(buf, binary.BigEndian, &length); err != nil {
		return 0, "", err
	}
	// 读取数据
	dataBuf := make([]byte, length)
	if err := binary.Read(buf, binary.BigEndian, &dataBuf); err != nil {
		return 0, "", err
	}
	return tag, string(dataBuf), nil
}
