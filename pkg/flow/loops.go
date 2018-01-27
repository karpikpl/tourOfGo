package flow

func sum() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	return sum
}

func sumOptional() int {
	sum := 1
	// note that init (i:=0) and post (i++) are optional, which basically is a "while" statment
	for sum < 1000 {
		sum += sum
	}
	return sum
}

func infiniteLoop() int {
	sum := 0

	// infinite loop (for without any conditions - while(true)
	for {
		if sum > 45 {
			return sum
		}
		sum++
	}
}
