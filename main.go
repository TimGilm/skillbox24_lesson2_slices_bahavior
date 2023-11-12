// поведение срезов
package main

import "fmt"

func main() {
	testSlice := make([]int, 0, 0)
	/*testSlice = append(testSlice, 1) //добавляем в срез цифру 1
	fmt.Println(testSlice)
	for i := 0; i < 10; i++ {
		testSlice = append(testSlice, i)
	}
	debugSlice(testSlice)*/
	s := make([]int, 0, 0)
	s = append(s, 1)
	s = append(s, 2)
	debugSlice(s)
	testSlice = append(testSlice, s...) //... -это spread или оператор распространения
	debugSlice(testSlice)               //оператор распространения нужен, чтобы добавить сразу весь срез

	//как делать обход среза?
	for i := 0; i < len(s); i++ {
		v := s[i]
		fmt.Println(i, v)
	}
	//но обычно используют оператор range
	for i, v := range s { //i -индекс, v - значение
		fmt.Println(i, v)
	}

}

func debugSlice(input []int) {
	fmt.Printf("data: %+v\n", input)
	fmt.Printf("len: %+v\n", len(input))
	fmt.Printf("cap: %+v\n", cap(input))
}
