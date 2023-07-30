package main

/*
定义一个Animal的接口，这里walk是他的成员函数
*/
type Animal interface {
	Walk()
	Talk()
}

type Machine interface {
	Charge()
}

/*
可以把两个interface组合在一起变成另外一个接口。这样ToyAnimal
接口就有了三个函数，walk，talk，charge
*/
type ToyAnimal interface {
	Animal
	Machine
}

/*
在go语言中，没有继承这个概念，因此，只要实现了walk和talk这两个函数，
就可以认为它是Animal接口的一个实现。需要注意的是，如果只实现了其中的
一个函数，则不能认为是接口的一个实现类。
*/
type Tiger struct {
}

func (a Tiger) Walk() {
	println("tiger walk with 4 feet")
}

/*
	func (a Tiger) Talk() {
		println("tiger can't talk")
	}
*/
type Human struct {
}

func (a Human) Walk() {
	println("human walk with 2 feet")
}
func (a Human) Talk() {
	println("Human can talk")
}

type Transformer struct {
}

func (t Transformer) Walk() {
	println("transformer can walk")
}

func (t Transformer) Talk() {
	println("transformer can Talk")
}

func (t Transformer) Charge() {
	println("transformer can be charged")
}

/*
和java，c++语言一样，go语言提倡在函数参数中使用接口，这样
传入的是可以是任何实现该接口函数的所有对象
*/
func TestInterface(a Animal) {
	a.Walk()
}

func StartTest() {
	/*
		这里由于Tiger没有实现Talk类，因此这行代码无法编译，因为Tiger
		不被go编译器认为是一个合法的Animal的实现类
	*/
	//TestInterface(Tiger{})
	TestInterface(Human{})
	TestInterface2(Transformer{})
}

func TestInterface2(t Transformer) {
	t.Talk()
	t.Walk()
	t.Charge()
}
