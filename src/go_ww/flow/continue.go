package flow

func ContinueFlow() {
	for i := 0; i < 10; i++ {
		println("i:", i)
		if i == 2 {
			continue
		}
	}

label:
	for v := 0; v < 5; v++ {
		for j := 0; j < 5; j++ {
			if v == 2 && j == 2 {
				continue label
			}
			println(v, j)
		}
	}
}
