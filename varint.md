详解varint编码原理

Kevin
发布于 2019-09-26
什么是Varint编码
Varint是一种使用一个或多个字节序列化整数的方法，会把整数编码为变长字节。对于32位整型数据经过Varint编码后需要1~5个字节，小的数字使用1个byte，大的数字使用5个bytes。64位整型数据编码后占用1~10个字节。在实际场景中小数字的使用率远远多于大数字，因此通过Varint编码对于大部分场景都可以起到很好的压缩效果。

编码原理
除了最后一个字节外，varint编码中的每个字节都设置了最高有效位（most significant bit - msb）–msb为1则表明后面的字节还是属于当前数据的,如果是0那么这是当前数据的最后一个字节数据。每个字节的低7位用于以7位为一组存储数字的二进制补码表示，最低有效组在前，或者叫最低有效字节在前。这表明varint编码后数据的字节是按照小端序排列的。

关于字节排列的方式引用一下维基百科上的词条

字节的排列方式有两个通用规则。例如，一个多位的整数，按照存储地址从低到高排序的字节中，如果该整数的最低有效字节（类似于最低有效位）在最高有效字节的前面，则称小端序；反之则称大端序。在网络应用中，字节序是一个必须被考虑的因素，因为不同机器类型可能采用不同标准的字节序，所以均按照网络标准转化。
通俗一点说就是：大端序是按照数字的书写顺序排列的，而小端序是颠倒书写顺序进行排列的。

看下面的图示会更好理解一些

图片描述

图中对数字123456进行varint编码，123456用二进制表示为1 11100010 01000000，每次低从向高取7位再加上最高有效位变成1100 0000 11000100 00000111 所以经过varint编码后123456占用三个字节分别为192 196 7。

解码的过程就是将字节依次取出，去掉最高有效位，因为是小端排序所以先解码的字节要放在低位，之后解码出来的二进制位继续放在之前已经解码出来的二进制的高位最后转换为10进制数完成varint编码的解码过程。

编码实现
由于protocol buffer中大量使用了varint编码，我从github.com/golang/protobuf/proto库中找到了对数据进行varint编解码的Go语言实现方法，实现代码中用位运算完成了上面说的varint编码过程。

const maxVarintBytes = 10 // maximum length of a varint

// 返回Varint类型编码后的字节流
func EncodeVarint(x uint64) []byte {
var buf [maxVarintBytes]byte
var n int
// 下面的编码规则需要详细理解:
// 1.每个字节的最高位是保留位, 如果是1说明后面的字节还是属于当前数据的,如果是0,那么这是当前数据的最后一个字节数据
//  看下面代码,因为一个字节最高位是保留位,那么这个字节中只有下面7bits可以保存数据
//  所以,如果x>127,那么说明这个数据还需大于一个字节保存,所以当前字节最高位是1,看下面的buf[n] = 0x80 | ...
//  0x80说明将这个字节最高位置为1, 后面的x&0x7F是取得x的低7位数据, 那么0x80 | uint8(x&0x7F)整体的意思就是
//  这个字节最高位是1表示这不是最后一个字节,后面7为是正式数据! 注意操作下一个字节之前需要将x>>=7
// 2.看如果x<=127那么说明x现在使用7bits可以表示了,那么最高位没有必要是1,直接是0就ok!所以最后直接是buf[n] = uint8(x)
//
// 如果数据大于一个字节(127是一个字节最大数据), 那么继续, 即: 需要在最高位加上1
for n = 0; x > 127; n++ {
// x&0x7F表示取出下7bit数据, 0x80表示在最高位加上1
buf[n] = 0x80 | uint8(x&0x7F)
// 右移7位, 继续后面的数据处理
x >>= 7
}
// 最后一个字节数据
buf[n] = uint8(x)
n++
return buf[0:n]
}
0x7F的二进制表示是0111 1111 ，所以x & 0x7F 与操作时，得到x二进制表示的最后7个bit位（前面的bit位通过与0做位与运算都被舍弃了）
0x80 的二进制表示是 1000 0000 ，所以 0x80 | uint8(x&0x7F)是在取出的x的后7个bit位前加上1（msb）
解码实现
解码就是编码的逆过程，同样是用位运算就能快速有效的完成解码，结合下面的代码注释再在纸上推演一遍理解起来就不难了。

func DecodeVarint(buf []byte) (x uint64, n int) {
for shift := uint(0); shift < 64; shift += 7 {
if n >= len(buf) {
return 0, 0
}
b := uint64(buf[n])
n++
// 下面这个分成三步走:
// 1: b & 0x7F 获取下7bits有效数据
// 2: (b & 0x7F) << shift 由于是小端序, 所以每次处理一个Byte数据, 都需要向高位移动7bits
// 3: 将数据x和当前的这个字节数据 | 在一起
x |= (b & 0x7F) << shift
if (b & 0x80) == 0 {
return x, n
}
}

    // The number is too large to represent in a 64-bit value.
    return 0, 0
}
Playground

到这里varint的编解码过程就都搞懂了，理解了varint编码原理后再看protocol buffer的编码原理就会容易很多。