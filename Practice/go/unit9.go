package main

import "fmt"

/*
문자열 사용하기
> 일반 문자열부터 유니코드까지 사용가능.
> \n으로
*/
func main() {

	var str1 string = "a"
	var str2 string = `test
  test`
	fmt.Println(str1)
	fmt.Println(str2)
}
