package pattern

import "fmt"

//Паттерн "строитель"

//iBuilder общий интерфейс для строителя
type iBuilder interface {
	setDoorType()
	setRoofType()
	setWallAmount()
	getHouse() *house
}

//house структура дома
type house struct {
	doorType   string
	roofType   string
	wallAmount int
}

//woodenBuilder структура строителя из дерева
type woodenBuilder struct {
	doorType   string
	roofType   string
	wallAmount int
}

//brickBuilder структура строителя из кирпича
type brickBuilder struct {
	doorType   string
	roofType   string
	wallAmount int
}

//director структура директора
type director struct {
	builder iBuilder
}

//newDirector конструктор директора
func newDirector(i iBuilder) *director {
	return &director{
		builder: i,
	}
}

//getBuilder определение конкретного конструктора строителя
func getBuilder(prof string) iBuilder {
	switch {
	case prof == "W":
		return newWoodenBuilder()
	case prof == "B":
		return newBrickBuilder()
	default:
		return nil
	}
}

//newWoodenBuilder конструктор строителя из дерева
func newWoodenBuilder() *woodenBuilder {
	return &woodenBuilder{}
}

//newBrickBuilder конструктор строителя из кирпича
func newBrickBuilder() *brickBuilder {
	return &brickBuilder{}
}

//Реализация интерфейса iBuilder для строителя из дерева
func (w *woodenBuilder) setDoorType() {
	w.doorType = "Wooden door"
}

func (w *woodenBuilder) setRoofType() {
	w.roofType = "Wooden roof"
}

func (w *woodenBuilder) setWallAmount() {
	w.wallAmount = 4
}

func (w *woodenBuilder) getHouse() *house {
	return &house{
		doorType:   w.doorType,
		roofType:   w.roofType,
		wallAmount: w.wallAmount,
	}
}

//Реализация интерфейса iBuilder для строителя из кирпича
func (b *brickBuilder) setDoorType() {
	b.doorType = "Brick door"
}

func (b *brickBuilder) setRoofType() {
	b.roofType = "Brick roof"
}

func (b *brickBuilder) setWallAmount() {
	b.wallAmount = 6
}

func (b *brickBuilder) getHouse() *house {
	return &house{
		doorType:   b.doorType,
		roofType:   b.roofType,
		wallAmount: b.wallAmount,
	}
}

//setBuilder смена строителя для директора
func (d *director) setBuilder(i iBuilder) {
	d.builder = i
}

//buildHouse пошаговая реализация постройки дома (дополнительная абстракция с помощью директора)
func (d *director) buildHouse() *house {
	d.builder.setWallAmount()
	d.builder.setRoofType()
	d.builder.setDoorType()
	return d.builder.getHouse()
}

//BuilderPattern реализация паттерна
func BuilderPattern() {
	woodenBuilder := getBuilder("W")
	brickBuilder := getBuilder("B")

	director := newDirector(woodenBuilder)
	woodenHouse := director.buildHouse()

	fmt.Printf("Wooden house door type: %s\n", woodenHouse.doorType)
	fmt.Printf("Wooden house roof type: %s\n", woodenHouse.roofType)
	fmt.Printf("Wooden house wall amount: %d\n", woodenHouse.wallAmount)

	director.setBuilder(brickBuilder)
	brickHouse := director.buildHouse()

	fmt.Println()
	fmt.Printf("Brick house door type: %s\n", brickHouse.doorType)
	fmt.Printf("Brick house roof type: %s\n", brickHouse.roofType)
	fmt.Printf("Brick house wall amount: %d\n", brickHouse.wallAmount)
}

/*
	Применяемость:
	- Когда ваш код должен создавать разные представления какого-то объекта. Например,
	деревянные и железобетонные дома.
	- Строитель можно применить, если создание нескольких представлений объекта состоит из
	одинаковых этапов, которые отличаются в деталях.
	Интерфейс строителей определит все возможные этапы конструирования.
	Каждому представлению будет соответствовать собственный класс-строитель.
	А порядок этапов строительства будет задавать класс-директор.

	-Когда вам нужно собирать сложные составные объекты, например, деревья Компоновщика.
	- Строитель конструирует объекты пошагово, а не за один проход. Более того, шаги строительства
	можно выполнять рекурсивно. А без этого не построить древовидную структуру, вроде Компоновщика.

	Плюсы:
	- Позволяет создавать продукты пошагово.
	- Позволяет использовать один и тот же код для создания различных продуктов.
	- Изолирует сложный код сборки продукта от его основной бизнес-логики.
	Минусы:
	- Усложняет код программы из-за введения дополнительных классов.
	- Клиент будет привязан к конкретным классам строителей, так как в интерфейсе
	директора может не быть метода получения результата.
*/
