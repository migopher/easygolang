package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "hello world!"
	str := "中华人民共和国"
	fmt.Println(strings.Contains(s, "world"))         //使用strings包中都Contains方法判断字符串中是否包含某给字符串
	fmt.Println(strings.ContainsAny(str, "国"))        //匹配更广泛都内容Unicode字符串
	fmt.Println(strings.HasPrefix(s, "hello"))        //字符串前缀是否包含
	fmt.Println(strings.HasSuffix(s, "!"))            //字符串后缀是否包含
	fmt.Println(strings.Index(s, "l"))                //返回字符串索引值 如果值不存在则返回-1
	fmt.Println(strings.LastIndex(s, "!"))            //返回字符串最后一个字符出现都位置，如果不存在则返回-1
	fmt.Println(strings.IndexRune(str, '和'))          //字符串占3个字符 15
	fmt.Println(strings.Replace(s, "o", "golang", 1)) //将world替换为golang,第四个参数表示替换到第几个，-1表示替换所有相符合到内容
	fmt.Println(strings.Count(s, "o"))                // 统计字符串出现到内容次数，第一个参数为原字符串，第二参数为要统计到子字符串

	fmt.Println(len([]rune(str)))            //计算字符数量
	fmt.Println(utf8.RuneCountInString(str)) //更推荐这种方式计算字符数量，但只是统计使用utf8，可以不用这种方式统计,减少包到引入

	fmt.Println(strings.ToUpper(s))                   //字符串转大写
	fmt.Println(strings.ToLower(s))                   //字符串转小写
	fmt.Println(strings.Trim(s, "!"))                 //修剪字符串前后
	fmt.Println(strings.TrimLeft(s, "hello "))        //修剪字符串左边
	fmt.Println(strings.TrimRight(s, "!"))            //修剪字符串右边
	fmt.Println(strings.TrimSpace("    " + s + "  ")) //修剪字符串左右空格
	fmt.Println(strings.Split(s, " "))                //将字符串切割
	fmt.Println(strings.SplitN(s,"",2))    //对字符串进行切割，第三个参数用来控制总切割
	var strSli = strings.Fields(s)                    //将字符串按照空白就行切割成字符串切片
	fmt.Println(strSli)
	fmt.Println(strings.Join(strSli, " "))         //将字符串切片拼接成一个字符串
	fmt.Println(strings.FieldsFunc(s, fieldsFunc)) //使用FieldsFunc 函数来确定分割符
	fmt.Println(strings.SplitAfter(s,"o"))   //要切割字符都后一个字符开始切割
	fmt.Println(strings.SplitAfterN(s,"l",-1)) //要切割字符都后一个字符开始切割总数为第三个参数都值-1表示回所有的子字符串组成的切
	fmt.Println(strings.IndexFunc(s,indexFunc))		//传入满足函数
	fmt.Println(strings.IndexByte(s,'h'))
}
func fieldsFunc(r rune) bool {
	return r == rune('o')
}
func indexFunc(r rune)bool  {
	return r == rune(' ')
}
