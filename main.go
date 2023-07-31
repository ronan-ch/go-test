package main

import (
	"strconv"
)

func main() {
	var number = "10"
	parseInt, _ := strconv.ParseInt(number, 10, 64)

	parseInt++
	println(parseInt)

	//定义一个切片，也可以认为是数组
	var aa = []int{1, 2, 3}
	aa = append(aa, 4, 5, 6, 7, 8, 9, 10)

	ints := aa[1:] //对于aa做一个子切片，这个和字符串一样的
	for i := range ints {
		print(ints[i])
	}
	var name string = "Ronan Chen"
	subname := name[2:]
	println(subname)

	m := map[string]int{
		"ronan":  1,
		"yaojie": 2,
	}
	println(m["ronan"])
	//loop1()

	//fun1(1, 2, 3, 4)
	//println(funtest(10))
	val1 := 1
	val2 := 1
	bypoint(&val1, val2)
	println(val1)
	println(val2)

	person := Person{
		name:   "ronan",
		gender: "male",
		age:    47,
	}
	person.introduce()

	person.setName("Leona")
	println(person.name)

	//println("test interface")
	//StartTest()

	//println("test rtti")
	//StartTestRtti()

	//testimport()

	//visitChannel()

	//channelSwitch2()
	//waitGroupTest()
	//Play()
	testHttpClient()
}

func fun1(vals ...int) int {
	for _, val := range vals {
		println(val)
	}
	return len(vals)
}
