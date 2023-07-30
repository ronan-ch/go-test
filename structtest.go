package main

import "strconv"

type Person struct {
	name   string
	gender string
	age    int
}

/*
定义成员函数需要把this指针（类似）放在函数名前面，
函数名后面的是成员函数正常的参数
*/
func (p Person) introduce() {
	a := p.name + " " + p.gender + strconv.Itoa(p.age)
	println(a)
}

/*
这里也支持真正的指针传递，这样对struct中的值进行修改
本质上可以认为这个Person对象就是一个函数参数，因此值传递和
指针传递和普通函数参数没有什么区别
*/
func (p *Person) setName(nm string) {
	p.name = nm
}
