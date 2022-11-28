package pattern

//Паттерн "цепочка обязанностей"

import "fmt"

type User struct {
	loginCorrect bool
	phoneCorrect bool
	emailCorrect bool
	registerDone bool
}

//Website Единый интерфейс под звенья цепочки обработчиков запроса
type Website interface {
	execute(*User)
	setNext(Website)
}

//Login Конкретный обработчик корректности логина
type Login struct {
	next Website
}

func (l *Login) execute(u *User) {
	if u.loginCorrect {
		fmt.Println("User already exist!")
		l.next.execute(u)
		return
	}
	fmt.Println("Correct login")
	u.loginCorrect = true
	l.next.execute(u)
}

func (l *Login) setNext(next Website) {
	l.next = next
}

//Phone Конкретный обработчик номера телефона
type Phone struct {
	next Website
}

func (p *Phone) execute(u *User) {
	if u.phoneCorrect {
		fmt.Println("Phone number already exist!")
		p.next.execute(u)
		return
	}
	fmt.Println("Correct phone number")
	u.phoneCorrect = true
	p.next.execute(u)
}

func (p *Phone) setNext(next Website) {
	p.next = next
}

//Email Конкретный обработчик корректности почты
type Email struct {
	next Website
}

func (e *Email) execute(u *User) {
	if u.emailCorrect {
		fmt.Println("Email address already exist!")
		e.next.execute(u)
		return
	}
	fmt.Println("Correct email address")
	u.emailCorrect = true
	e.next.execute(u)
}

func (e *Email) setNext(next Website) {
	e.next = next
}

//Register Конкретный обработчик регистрации пользователя
type Register struct {
	next Website
}

func (r *Register) execute(u *User) {
	if u.registerDone {
		fmt.Println("Register done")
	}
	fmt.Println("Account have been registered!")
}

func (r *Register) setNext(next Website) {
	r.next = next
}

func ChainOfResPattern() {
	register := &Register{}

	login := &Login{}
	login.setNext(register)

	phone := &Phone{}
	phone.setNext(login)

	email := &Email{}
	email.setNext(phone)

	user := &User{}
	email.execute(user)
}

/*
	Применимость:

	- Когда программа должна обрабатывать разнообразные запросы несколькими способами,
	но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
	- С помощью Цепочки обязанностей вы можете связать потенциальных обработчиков в одну цепь
	и при получении запроса поочерёдно спрашивать каждого из них, не хочет ли он обработать запрос.

	- Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	- Цепочка обязанностей позволяет запускать обработчиков последовательно один за другим
	в том порядке, в котором они находятся в цепочке.

	- Когда набор объектов, способных обработать запрос, должен задаваться динамически.
	- В любой момент вы можете вмешаться в существующую цепочку и переназначить связи так,
	чтобы убрать или добавить новое звено.

	Плюсы:
	- Уменьшает зависимость между клиентом и обработчиками.
	- Реализует принцип единственной обязанности.
	- Реализует принцип открытости/закрытости.

	Минусы:
	- Запрос может остаться никем не обработанным.
*/
