package main

func loop1() {
	println("enter loop1")
	var vals []int = []int{1, 2, 3, 4, 5, 6}
	var intset map[int]bool = map[int]bool{}
	for _, val := range vals {
		intset[val] = true
	}
	println(intset[3])
	println("exit loop1")
}
