package main

import (
	"errors"
	"fmt"
	"sort"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println(Hello(" Stefan"))
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("You"))

	if _, err := HelloError(""); err != nil {
		fmt.Println("Handle error ", err.Error())
	}

	var b bool = true
	c := false // Short declaration using := infers the type
	var d, e string = "Hello", "world!"

	m1, m2 := multipleReturn(2, 3)
	_, m3 := multipleReturn(7, 8) // _ ignores return value

	fmt.Printf("\nHello world %v, %v, %v, %v, %v, %v, %v\n\n", b, c, d, e, m1, m2, m3)

	// Struct
	p := Person{"Stefan", "Nilsson", 44, "770822"}
	p.PrintFullName()

	fmt.Printf("Person pointer 0x%p\n", &p)
	p.setPnummerValueReceiver("888888")
	p.PrintFullName()
	p.setPnummerPointerReceiver("777777")
	p.PrintFullName()

	// Slice
	sliceIsList := make([]string, 3, 10) // The make built-in function allocates and initializes an object of type slice, map, or chan (only)
	appended := "8"
	sliceIsList = append(sliceIsList, appended)

	fmt.Println("slice has length ", len(sliceIsList), " and capacity ", cap(sliceIsList))
	for i := 0; i < 100; i++ { // Increment expression using ;
		if sliceIsList[i] == appended {
			fmt.Println("Found ", appended, ", break!")
			break
		}
	}
	for i := 0; ; { // Possible to skip initialization, test or increment ;
		if sliceIsList[i] == appended {
			fmt.Println("Found ", appended, " again, break!")
			break
		} else {
			i++
		}
	}
	for index, value := range sliceIsList { // range using ,
		fmt.Println("At index ", index, " value ", value)
	}
	for _, value := range sliceIsList {
		fmt.Println("Value only loop", value)
	}

	people := []Person{
		{"Bob", "B", 31, "770822"},
		{"John", "B", 42, "770822"},
		{"Michael", "B", 17, "770822"},
		{"Jenny", "B", 26, "770822"},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people)) // sort using ByAge sort struct
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool { // sort using Less function as closure
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

	mySort := MySort{}
	mySort.MethodAddedToStringSlice()

	testConcurrentChannels()
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	pnumber   string // lowercase "package private"
}

// ByAge implements sort.Interface for []Person based on the Age field.
// Interfaces are satisfied implicitly, no need to explicitly say that an interface is implemented just implement all methods
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

/////////////////////////////

func (receiver Person) PrintFullName() { // receiver, think "this" (java)
	fmt.Printf("%s %s %s\n", receiver.FirstName, receiver.LastName, receiver.pnumber)
}

func (receiver *Person) setPnummerPointerReceiver(a string) { // Methods on struct declared "like" extension functions in kotlin
	fmt.Printf("Person pointer in ponter receiver method 0x%p\n", receiver)
	receiver.pnumber = a
}

func (receiver Person) setPnummerValueReceiver(a string) {
	fmt.Printf("Person pointer in value receiver method 0x%p\n", &receiver)
	// p.pnumber = a // not possible since value receiver is a copy of original value
}

func multipleReturn(a, b int) (int, int) {
	return a + b, a * b
}

// MySort You can only define methods on types defined in the same package.
// To add methods to builtin type or struct in other package, define type alias like MyInt because we can’t add methods to built-int type int.
type MySort sort.StringSlice

func (mi MySort) MethodAddedToStringSlice() {
	fmt.Print("MethodAddedToStringSlice")
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func HelloError(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	return Hello(name), nil
}
