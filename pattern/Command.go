package pattern

import "fmt"

//Паттерн "Команда"

// iCommand интерфейс команды
type iCommand interface {
	execute()
}

// device интерфейс устройства
type device interface {
	on()
	off()
}

//button структура кнопки включения/выключения
type button struct {
	command iCommand
}

//commandOn структура команды включение
type commandOn struct {
	device device
}

//commandOff структура команды выключения
type commandOff struct {
	device device
}

//tv структура устройства (телевизор)
type tv struct {
	isRunning bool
}

//press реализация нажатия на кнопку включения/выключения
func (b *button) press(on, off iCommand) {
	b.command.execute()
	switch {
	case b.command == on:
		b.setCommand(off)
	case b.command == off:
		b.setCommand(on)
	}
}

func (b *button) setCommand(c iCommand) {
	b.command = c
}

//Реализация интерфейса команд
func (on *commandOn) execute() {
	on.device.on()
}

func (off *commandOff) execute() {
	off.device.off()
}

//Реализация интерфейса устройства
func (tv *tv) on() {
	tv.isRunning = true
	fmt.Println("TV is on")
}

func (tv *tv) off() {
	tv.isRunning = false
	fmt.Println("TV is off")
}

//CommandPattern реализация паттерна
func CommandPattern() {
	tv := &tv{}
	commandOn := &commandOn{device: tv}
	commandOff := &commandOff{device: tv}

	button := &button{commandOn}

	for i := 1; i < 11; i++ {
		fmt.Printf("%v ", i)
		button.press(commandOn, commandOff)
	}
}

/*
	Применяемость:
	- Когда вы хотите параметризовать объекты выполняемым действием.
	- Команда превращает операции в объекты. А объекты можно передавать, хранить и взаимозаменять внутри других объектов.

	- Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
	- Как и любые другие объекты, команды можно сериализовать, то есть превратить в строку, чтобы потом сохранить в
	файл или базу данных. Затем в любой удобный момент её можно достать обратно, снова превратить в объект команды и
	выполнить. Таким же образом команды можно передавать по сети, логировать или выполнять на удалённом сервере.

	Плюсы:
	- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	- Позволяет реализовать простую отмену и повтор операций.
	- Позволяет реализовать отложенный запуск операций.
	- Позволяет собирать сложные команды из простых.
	- Реализует принцип открытости/закрытости.
	Минусы:
	- Усложняет код программы из-за введения множества дополнительных классов.

*/
