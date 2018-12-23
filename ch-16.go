package main

/**
自定义类型
 */
type stu struct {
	Name string //首字母大写表示其他包可以直接使用 表示公有
	age  int    //首字母小写不允许外面都包调用 表示私有
}

func main() {

}
