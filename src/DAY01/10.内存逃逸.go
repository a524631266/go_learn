package main

/**
 * 内存逃逸，是为，在栈上的元素，被移动到堆上
 */
func main() {
	city2 := showLocalCity2()
	println("city2:::", city2)
	println("*city2:::", *city2)
}
func showLocalCity2() *string {
	city := "杭州"
	ptr := &city
	return ptr
}
