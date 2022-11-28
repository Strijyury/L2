package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test2() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test2()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

/*

Выведется слово "error" т.к. мы сравниваем переменную и nil и они не равны

*/
