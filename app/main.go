package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println("Hello, this package should be able to access other packages")
	fmt.Println("Name: ", stuff.Name)
	intArr := []int{1,2,3,4,5}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}