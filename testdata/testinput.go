package main

import "fmt"

func main() {

	var (
		userName string
		email    string
	)
	fmt.Print("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Print("请输入邮箱：")
	fmt.Scan(&email)
	fmt.Println(userName, email)

}
