package main

func main() {
	err := processHeap()
	if err != nil {
		panic(err)
	}

	err = processSort()
	if err != nil {
		panic(err)
	}
}
