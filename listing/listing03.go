package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

/*

В первом случае выведет nil, т.к. мы создаем экземпляр функции в которой лежит переменная равная nil.
Во втором случае выведется false, т.к. переменная не равна nil.

*/
