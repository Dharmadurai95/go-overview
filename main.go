package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("hello world")
	hel, wo := returnMultiVal()
	fmt.Println(hel, wo, "hello world")

	//uint -> we cannot define negative values
	var ageOne int = 25
	const ageTwo uint = 28
	ageThree := 30
	fmt.Println(ageOne, ageTwo, ageThree)

	// formatedStr()
	// arraySlice()
	// standardLib()
	// goLoops()
	// functionOverview()
	//passByValBassByRef()

	//bill summary
	// billSummary := newBill("Dharmadurai")
	// billSummary.addItem("onion soup", 7.48)
	// billSummary.addItem("puttin", 6.48)
	// billSummary.addItem("pizza", 8.48)

	// billSummary.addTip(8.55)
	// fmt.Println(billSummary.formatBill())
	createBill()

}

func sayGoodMor(n string) {
	fmt.Printf("Good morning %v \n", n)
}
func sayGoodBye(n string) {
	fmt.Printf("Good bye %v ", n)
}
func printMsg(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

func passByValBassByRef() {
	//passby val string,int,float,boolean,struct,array
	//passby ref -function ,slice,map
}

func returnHelperMethos(msg string) (func(r float32) float32, func(num int) int) {
	return func(r float32) float32 {
		fmt.Println(msg, "hello world")
		return math.Pi * r * r
	}, func(num int) int { return num * 2 }
}

func functionOverview() {
	printMsg([]string{"Dharma", "durai", "kvb bank", "response", "hello"}, sayGoodMor)
	circleArea, multiplyTwo := returnHelperMethos("Hi this is dharmadurai")
	fmt.Println(circleArea(10.5))
	fmt.Println(multiplyTwo(10))

}

func returnMultiVal() (string, string) {
	return "hello", "world"
}

// standartLibrary
func standardLib() {
	// greeting := "hi there friends"
	// fmt.Println(strings.Contains(greeting, "hi"))
	// fmt.Println(strings.ReplaceAll(greeting, "hi", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, " "))
	// fmt.Println(strings.Split(greeting, " "))

	// //original string
	// fmt.Println("Original string", greeting)

	// ***********sort *******
	ages := []int{1000, 21, 34, 25, 43, 63, 18, 10, 30}
	sort.Ints(ages) //it will mutate original val which mean this is pass by ref
	fmt.Println(ages)
	//search
	index := sort.SearchInts(ages, 100)
	fmt.Println(index) // if you provide invalid search val like 25852 it will give length +1 val for ex length is 5 if you give invalid input the output will be 6

	names := []string{"azhagar", "Ram", "zeal", "Dharma", "Anand"}
	sort.Strings(names)
	fmt.Println(names)
	nameInd := sort.SearchStrings(names, "zeal")
	fmt.Println(nameInd)
}

func arraySlice() {
	//array
	var marks = [5]int{45, 56, 54, 42, 100} // for array we need to specify the length and datatype we cannot change length run time;
	// append(marks,10) this will throw error becase append method is available
	fmt.Println(marks[1:])  //SLICE THE ARRAY FROM 1 INDEX TO END OF ARRAY
	fmt.Println(marks[2:4]) //slice the array from 2 index to 3 third index because second index will not include
	// change item val
	marks[4] = 65
	fmt.Println(marks)
	fmt.Println("capacity or arr", cap(marks))

	//slice
	fmt.Println("****************slice ****************")
	names := []string{"Dharma", "Durai", "Muthusamy"} // in slice we don't need to specify the length we can change in runtime;
	names = append(names, "Illusion")                 // this will not modify oiriginal array so we need to update original variable;
	fmt.Println("Length of slice", len(names))
	fmt.Println("Capacity of slice", cap(names))
	fmt.Println(names[0:2])           //  fmt.Println(names[:2])  fetch first two item  both are correct
	fmt.Println(names[len(names)-2:]) //fetch last two items

}

func formatedStr() {
	var name = "Dharma"
	var age = 28

	fmt.Print("Hello there ...")
	fmt.Print("Hello Ninja ...\n")
	// above both strings are print in same line  if you want print separate line you \n inside the string ex --> fmt.Print("Hello there ... \n")

	fmt.Println("it will print the message in separate line") // it will print new line

	fmt.Println("hello mr", name, "you are", age, "year old boy")
	fmt.Printf("Hello mr %v you are %v year old boy \n", name, age) // formate the strings  v - variables
	fmt.Printf("Hello mr %q you are %q year old boy\n", name, age)  // formate the strings  q - quotes (it will work only for strings)

	// it will return formated string but above statement will print in console.
	var formatStr = fmt.Sprintf("Hello mr %v you are %v year old boy\n", name, age)
	fmt.Println("formatted string", formatStr)

	fmt.Printf("Find the type %T \n", 20)
	fmt.Printf("Find the type %f \n", 2000.25)
	fmt.Printf("fixed length %0.2f \n", 20.0025) // fixed 2  ex 20.00

}

// loops
func goLoops() {
	//while loop
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is :", x)
	// 	x++
	// }

	//for looop
	// for x := 0; x < 5; x++ {
	// 	fmt.Println("value of x is :", x)

	// }

	//loops the slices or array
	names := []string{"dharma", "durai", "muthusamy", "priya", "gokula"}
	// for ind := 0; ind < len(names); ind++ {
	// 	fmt.Println("iterate name is:", names[ind])
	// }

	for ind, val := range names {
		fmt.Printf("the value is %v  and the index is %v \n", val, ind)
	}

}
