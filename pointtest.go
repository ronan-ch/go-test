package main

/*
由于val1是按照指针传递，因此可以在bypoint函数中被改变
val2是默认的值传递（即copy传递），因此在bypoint里修改不会
改变main函数中的值
*/
func bypoint(val1 *int, val2 int) {
	println("start test by point")
	*val1++
	val2++
	println("end test by point")
}
