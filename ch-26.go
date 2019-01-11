package main

import "fmt"

/**
标签
标签的主要作用是用于跳转到某给地方继续执行
标签建议别使用保留字（for等）
标签区分大小写
未使用的标签会引发错误
建议需要才使用
 */

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			goto GO
		}
		fmt.Println(i)
	}
GO:
	fmt.Println("stop")
}
