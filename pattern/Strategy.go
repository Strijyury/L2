package pattern

//Паттерн "стратегия"

import "fmt"

//EvictionAlgo - Общий интерфейс стратегий
type EvictionAlgo interface {
	evict(c *Cache)
}

//Fifo Конкретная стратегия
type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

//Lru Конкретная стратегия
type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

//Lfu Конкретная стратегия
type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}

//Структура контекста включает в себя интерфейс стратегии. За счет этого в структуру
//контекста можно передавать новую стратегию. В конкретно этом примере реализуются три разные
//стратегии по очистке кэша

//Cache - пример контекста(кэш)
type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

//initCache - конструктор кэша
func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

//setEvictionAlgo - замена стратегии
func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

//Add - добавление в кэш элемента
func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

//Get - кэш возвращает элемент с его последующим удалением
func (c *Cache) get(key string) {
	delete(c.storage, key)
}

//Evict - реализует стратегию удаление, определяя нужную функцию evict через интерфейсное значение
//в структуре cashe
func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

//StrategyPattern - Реализация паттерна
func StrategyPattern() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")
}

/*
	Применимость:
	- Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
	- Стратегия позволяет варьировать поведение объекта во время выполнения программы,
	подставляя в него различные объекты-поведения (например, отличающиеся балансом скорости и
	потребления ресурсов).

	- Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
	- Стратегия позволяет вынести отличающееся поведение в отдельную иерархию классов,
	а затем свести первоначальные классы к одному, сделав поведение этого класса настраиваемым.

	- Когда вы не хотите обнажать детали реализации алгоритмов для других классов.
	- Стратегия позволяет изолировать код, данные и зависимости алгоритмов от других объектов,
	скрыв эти детали внутри классов-стратегий.

	- Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
	Каждая ветка такого оператора представляет собой вариацию алгоритма.
	- Стратегия помещает каждую лапу такого оператора в отдельный класс-стратегию.
	Затем контекст получает определённый объект-стратегию от клиента и делегирует ему работу.
	Если вдруг понадобится сменить алгоритм, в контекст можно подать другую стратегию.

	Плюсы:
	- Горячая замена алгоритмов на лету.
	- Изолирует код и данные алгоритмов от остальных классов.
	- Уход от наследования к делегированию.
	- Реализует принцип открытости/закрытости.

	Минусы:
	- Усложняет программу за счёт дополнительных классов.
	- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/
