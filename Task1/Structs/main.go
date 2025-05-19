package main

import "fmt"

//Struct definition
type student struct {
	name            string
	roll            int
	schoolOfstudent school
}

type school struct {
	noOfgroundsInschool int
	nameOfschool        string
}

//Embedded struct
type car struct {
	make  string
	model int
}
type truck struct {
	car
	truckType string
}

func test(s student) {

	fmt.Printf("Name of student is: %v\n", s.name)
	fmt.Printf("Roll no of student is: %v\n", s.roll)

}

func main() {

	test(student{
		name: "Nigga",
		roll: 01,
	})

	//Anonymous struct
	myNigga := struct {
		name string
	}{
		name: "Nigga1",
	}

	//defining embedded struct
	mytruck := truck{
		truckType: "16 wheeler",
		car: car{
			make:  "haval",
			model: 2022,
		},
	}

	myStudent := student{}
	myStudent.name = "Nigga 2"
	myStudent.roll = 2
	myStudent.schoolOfstudent.noOfgroundsInschool = 2
	myStudent.schoolOfstudent.nameOfschool = "Nigga school"

	fmt.Println(myStudent)
	fmt.Println(myNigga)
	fmt.Println(mytruck)

}
