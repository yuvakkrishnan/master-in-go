package main

// N is a shared counter which is BAD
var N int

func main() {
	showN()
	incrN()
	incrN()
	incrN()
	incrN()
	incrN()
	incrN()
	showN()

	incrN()
	incrN()
	incrN()
	incrN()
	incrN()
	incrN()
	incrN()
	incrN()
	showN()
}
