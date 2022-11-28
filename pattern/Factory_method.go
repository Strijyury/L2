package pattern

import (
	"fmt"
	"log"
)

// Паттерн "Фабричный метод"

//iCar интерфейс машины
type iCar interface {
	setEnginePower(enginePower int)
	setName(name string)
	getEnginePower() int
	getName() string
}

//car стрктура машины
type car struct {
	enginePower int
	name        string
}

//mercedes структура мерседеса с встроенной структурой машины
type mercedes struct {
	car
}

//audi структура ауди с встроенной структурой машины
type audi struct {
	car
}

//Реализация интерфейса машины
func (c *car) setEnginePower(enginePower int) {
	c.enginePower = enginePower
}

func (c *car) setName(name string) {
	c.name = name
}

func (c *car) getEnginePower() int {
	return c.enginePower
}

func (c *car) getName() string {
	return c.name
}

//newMercedes конструктор мерседеса
func newMercedes() iCar {
	return &mercedes{
		car{
			enginePower: 500,
			name:        "Mercedes E500",
		},
	}
}

//newAudi конструктор ауди
func newAudi() iCar {
	return &audi{
		car{
			enginePower: 450,
			name:        "Audi Q8",
		},
	}
}

//getCar фабрика, выпускающая несколько моделей машин
func getCar(carType string) (iCar, error) {
	switch {
	case carType == "mercedes":
		return newMercedes(), nil
	case carType == "audi":
		return newAudi(), nil
	default:
		return nil, fmt.Errorf("wrong car type")
	}
}

//Реализация паттерна
func FactoryMethodPattern() {
	myCar, err := getCar("mercedes")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("My car name: %s, my car engine power: %d\n", myCar.getName(), myCar.getEnginePower())

	myCar, err = getCar("audi")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("My car name: %s, my car engine power: %d\n", myCar.getName(), myCar.getEnginePower())

	myCar, err = getCar("bmw")
	if err != nil {
		log.Fatalf(err.Error())
	}
}

/*
	Применяемость:
	- Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
	- Фабричный метод отделяет код производства продуктов от остального кода, который эти продукты использует.
	Благодаря этому, код производства можно расширять, не трогая основной. Так, чтобы добавить поддержку нового
	продукта, вам нужно создать новый подкласс и определить в нём фабричный метод, возвращая оттуда экземпляр нового продукта.

	- Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.
	- Пользователи могут расширять классы вашего фреймворка через наследование. Но как сделать так, чтобы фреймворк
	создавал объекты из этих новых классов, а не из стандартных?
	Решением будет дать пользователям возможность расширять не только желаемые компоненты, но и классы,
	которые создают эти компоненты. А для этого создающие классы должны иметь конкретные создающие методы,
	которые можно определить.

	Плюсы:
	- Избавляет класс от привязки к конкретным классам продуктов.
	- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	- Упрощает добавление новых продуктов в программу.
	- Реализует принцип открытости/закрытости.
	Минусы:
	- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса
	продукта надо создать свой подкласс создателя.
*/
