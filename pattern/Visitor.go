package pattern

//Паттерн "Посетитель"

import (
	"fmt"
	"math"
)

//Общий интерфейс фигур
type Shape interface {
	getType() string
	accept(Visitor)
}

//Интерфейс "посетителя"
type Visitor interface {
	visitToCircle(*Circle)
	visitToSquare(*Square)
	visitToPentagon(*Rectangle)
}

//Структура конкретной фигуры
type Rectangle struct {
	side1, side2, angles int
}

func (r *Rectangle) accept(visitor Visitor) {
	visitor.visitToPentagon(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

//Структура конкретной фигуры
type Circle struct {
	radius float64
}

func (c *Circle) accept(visitor Visitor) {
	visitor.visitToCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

//Структура конкретной фигуры
type Square struct {
	side, angles int
}

func (s *Square) accept(visitor Visitor) {
	visitor.visitToSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

//Конкретный посетитель
type AreaOfShape struct {
	area float64
}

func (a *AreaOfShape) visitToSquare(s *Square) {
	a.area = float64(s.side * s.side)
	fmt.Println("Calculating area for square")
}

func (a *AreaOfShape) visitToCircle(c *Circle) {
	a.area = c.radius * c.radius * math.Pi
	fmt.Println("Calculating area for circle")
}

func (a *AreaOfShape) visitToPentagon(r *Rectangle) {
	a.area = float64(r.side2 * r.side1)
	fmt.Println("Calculating area for rectangle")
}

//Конкретный посетитель
type PerimeterOfShape struct {
	perimeter float64
}

func (pr *PerimeterOfShape) visitToSquare(s *Square) {
	pr.perimeter = float64(s.side * 4)
	fmt.Println("Calculating perimeter for square")
}

func (pr *PerimeterOfShape) visitToCircle(c *Circle) {
	pr.perimeter = 2 * c.radius * math.Pi
	fmt.Println("Calculating perimeter for circle")
}

func (pr *PerimeterOfShape) visitToPentagon(r *Rectangle) {
	pr.perimeter = float64(r.side1*2 + r.side2*2)
	fmt.Println("Calculating perimeter for rectangle")
}

//Реализация паттерна
func VisitPattern() {
	square := &Square{side: 10}
	circle := &Circle{radius: 5}
	rectangle := &Rectangle{side1: 5, side2: 6}

	area := &AreaOfShape{}

	circle.accept(area)
	fmt.Println(area.area)
	square.accept(area)
	fmt.Println(area.area)
	rectangle.accept(area)
	fmt.Println(area.area)

	perimeter := &PerimeterOfShape{}

	fmt.Println()
	circle.accept(perimeter)
	fmt.Println(perimeter.perimeter)
	square.accept(perimeter)
	fmt.Println(perimeter.perimeter)
	rectangle.accept(perimeter)
	fmt.Println(perimeter.perimeter)
}

/*
	Применимость:
	- Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры
	объектов, например, деревом.
	- Посетитель позволяет применять одну и ту же операцию к объектам различных классов.

	- Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между
	собой операции, но вы не хотите «засорять» классы такими операциями.
	- Посетитель позволяет извлечь родственные операции из классов, составляющих структуру объектов,
	поместив их в один класс-посетитель. Если структура объектов является общей для нескольких
	приложений, то паттерн позволит в каждое приложение включить только нужные операции.

	- Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.
	- Посетитель позволяет определить поведение только для этих классов, оставив его пустым для
	всех остальных.

	Плюсы:
	- Упрощает добавление операций, работающих со сложными структурами объектов.
	- Объединяет родственные операции в одном классе.
	- Посетитель может накапливать состояние при обходе структуры элементов.
	Минусы:
	- Паттерн не оправдан, если иерархия элементов часто меняется.
	- Может привести к нарушению инкапсуляции элементов.
*/
