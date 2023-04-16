package flow

func BreakFlow() {
	for i := 0; i < 10; i++ {
		println(i)
		if i == 3 {
			break
		}
	}

	num := 2
	switch num {
	case 1:
		println(1)
	case 2:
		println(2)
		break
	default:
		println("默认")
	}
}
