package pattern

import (
	"fmt"
	"log"
)

// Паттерн "Состояние"

// State определяем интерфейс для состояния
type state interface {
	wakeUpDad() error
	turnOffTV() error
	getState() string
}

// WatchingTVState - конкретное состояние
type watchingTVState struct {
	dad *dad
}

//Если действие не влияет на состояние - возвращаем ошибку, обозначающую текущее состояние папы
func (w *watchingTVState) wakeUpDad() error {
	return fmt.Errorf("отец смотрит телевизор")
}

//Если действие меняет состояние - обозначаем, что состояние изменилось и меняем текущее состояние
//папы через функцию setState()
func (w *watchingTVState) turnOffTV() error {
	fmt.Println("Мы выключили телевизор")
	fmt.Println("Отец возмущался, что телевизор выключили и ушел спать")
	w.dad.setState(w.dad.sleeping)
	return nil
}

func (w *watchingTVState) getState() string {
	return "Папа смотрит телевизор"
}

//SleepingState - конкретное состояние
type sleepingState struct {
	dad *dad
}

//Если действие меняет состояние - обозначаем, что состояние изменилось и меняем текущее состояние
//папы через функцию setState()
func (s *sleepingState) wakeUpDad() error {
	fmt.Println("Мы разбудили папу")
	fmt.Println("Отец возмущался, что его разбудили и пошел смотреть телевизор")
	s.dad.setState(s.dad.watchingTV)
	return nil
}

//Если действие не влияет на состояние - возвращаем ошибку, обозначающую текущее состояние папы
func (s *sleepingState) turnOffTV() error {
	return fmt.Errorf("отец спит")
}

func (s *sleepingState) getState() string {
	return "Папа спит"
}

//Dad - объект, чье состояние изменяется
type dad struct {
	watchingTV   state
	sleeping     state
	currentState state
}

func newDad() *dad {
	newDad := &dad{}
	watchingTV := &watchingTVState{dad: newDad}
	sleeping := &sleepingState{dad: newDad}
	newDad.watchingTV = watchingTV
	newDad.sleeping = sleeping
	newDad.setState(watchingTV)
	return newDad
}

func (d *dad) setState(s state) {
	d.currentState = s
}

//Так как Dad имплементирует тот же интерфейс, что и структура состояния, методы состояния будут
//вызываться полиморфно, в зависимости от текущего состояния папы и менять его при необходимости
func (d *dad) getState() string {
	return d.currentState.getState()
}

func (d *dad) wakeUpDad() error {
	return d.currentState.wakeUpDad()
}

func (d *dad) turnOffTV() error {
	return d.currentState.turnOffTV()
}

//StatePattern - реализация паттерна
func StatePattern() {
	dad := newDad()
	fmt.Println(dad.getState())

	err := dad.turnOffTV()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(dad.getState())

	err = dad.wakeUpDad()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(dad.getState())
}

/*
	Применяемость:
	- Когда код класса содержит множество больших, похожих друг на друга, условных операторов,
	которые выбирают поведения в зависимости от текущих значений полей класса.
	- Паттерн предлагает переместить каждую ветку такого условного оператора в собственный класс.
	Тут же можно поселить и все поля, связанные с данным состоянием.

	- Когда вы сознательно используете табличную машину состояний, построенную на условных операторах,
	но вынуждены мириться с дублированием кода для похожих состояний и переходов.
	- Паттерн Состояние позволяет реализовать иерархическую машину состояний, базирующуюся на
	наследовании. Вы можете отнаследовать похожие состояния от одного родительского класса и вынести
	туда весь дублирующий код.

	Плюсы:
	- Избавляет от множества больших условных операторов машины состояний.
	- Концентрирует в одном месте код, связанный с определённым состоянием.
	- Упрощает код контекста.
	Минусы:
	- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/
