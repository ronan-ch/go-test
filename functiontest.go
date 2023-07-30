package main

/*
go语言中支持函数指针，这个BinFun定义的就是一个二元函数，
接受两个int类型，返回一个int类型
*/
type BinFun func(int, int) int

func add(s1 int, s2 int) int {
	return s1 + s2
}

func sub(s1 int, s2 int) int {
	return s1 - s2
}

func mul(s1 int, s2 int) int {
	return s1 * s2
}

func div(s1 int, s2 int) int {
	return s1 / s2
}

func funtest(input int) int {
	mm := map[string]BinFun{
		"sum": add,
		"sub": sub,
		"mul": mul,
		"div": div,
	}

	result := input
	for _, val := range mm {
		result = val(result, input)
	}
	return result
}

/*
go支持匿名函数，创建了就直接调用，匿名函数没有名字
*/
func anonymousfuntest() {
	result := 0
	for i := 0; i < 10; i++ {
		result = result + func(input int) int {
			return input
		}(i)
	}
}
