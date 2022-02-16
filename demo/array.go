package demo

import "fmt"

type User struct {
	name string
	age  int
}

func (user User) SayHello() {
	fmt.Printf("Hello %s", user.name)
}

func sumArr(arr *[5]int) int {
	var sum int = 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func ArrayDemo() {
	var arr0 [5]int = [5]int{1, 2, 3}
	var arr1 = [5]int{1, 2, 3}

	var arr2 = [...]int{1, 3, 4, 5}
	var arr3 = [...]string{3: "hello", 4: "world"}
	var arr4 = [...]User{{"user1", 1}, {"user2", 2}}
	fmt.Println(arr0, arr1, arr2, arr3, arr4)
	a := User{"zhangsan", 30}
	a.SayHello()
	for i, user := range arr4 {
		user.SayHello()
		fmt.Println(i)
	}
	fmt.Println(sumArr(&arr1))
	fmt.Println(1 << 2)
	fmt.Println(fmt.Sprintf("hello %v", "world"))
	fmt.Printf("hello%+v", a)
	fmt.Printf("全局变量：arr %v\n", arr2)
}
