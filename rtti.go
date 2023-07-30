package main

type empty interface {
}

func typeTeller(a empty) {
	switch a.(type) {
	case int:
		println("this is int type")
	case Transformer:
		println("this is a transformer type")
	case Tiger:
		println("this is a Tiger type")
	default:
		println("this is a unidentifed type")
	}
}

func StartTestRtti() {
	typeTeller(Transformer{})
	typeTeller(4)
	typeTeller(Human{})
}
