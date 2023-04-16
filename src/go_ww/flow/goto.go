package flow

func GotoFlow() {
	const num int = 1
	if num == 2 {
		println("2")
	} else {
		goto LABEL1
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 && j == 2 {
				goto LABEL2
			}
		}
	}

LABEL1:
	println("goto跳过来了")
	println("3")

LABEL2:
	println("goto2")
}
