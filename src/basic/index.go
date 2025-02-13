package basic

import "fmt"

func Basic() {

	// variables

	// int
	var number1 int = 0
	var number2 = 2
	number3 := 3

	// float
	var number4 float32 = 0.34
	var number5 float64 = 0.78

	// string
	var str string = "Hello"
	str2 := "world!"

	// boolean
	bool1 := false
	bool2 := true

	fmt.Println(number1, number2, number3, number4, number5)
	fmt.Println(str, str2)
	fmt.Println(bool1, bool2)

	// array

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [5]string{"Hey", "there,", "How ", "are", "you!"}

	fmt.Println(arr1, arr2)

	// slice

	slice1 := []int{5, 6, 7, 8, 9}
	slice2 := []string{"Hello", "world!", "Good"}

	fmt.Println(slice1, slice2)

	// string format
	fmt.Print("Hello, world! \n") // never break string at the end.. explicitly add \n to break;
	fmt.Println("Break line automatically")
	fmt.Printf("Type of %v, %v, %v, %v is %T, %T, %T, %T \n", arr1, arr2, slice1, slice2, arr1, arr2, slice1, slice2)
	strFormat := fmt.Sprintf("Type of slice and Array is %T, %T", slice1, arr1)
	fmt.Println(strFormat)

	// array operations
	fmt.Println(append(slice2, "bye!"))
	fmt.Println(arr2[1:], "-", arr2[:2], "-", arr2[0:3], "-", len(arr1))

	// loops

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	// can't mutate array value as it is making duplicate copy of value in side arr2
	for i, value := range arr2 {
		fmt.Printf("Index: %d, Value: %s\n", i, value)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// mutate array in loops
	for i := 0; i < len(arr2); i++ {
		arr2[i] = "new Data"
	}

	fmt.Println(arr2)

	// map

	map1 := map[string]string{
		"name":    "Shaun",
		"surName": "bali",
		"age":     "25",
	}
	fmt.Println((map1["name"]))

	for k, v := range map1 {
		fmt.Printf("%v - %v \n", k, v)
	}
}
