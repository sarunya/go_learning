package main

import "fmt"

type ContactInfo struct {
	email string
	phone string
}

type Person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

func main() {
	structDemo()
}

func mapSliceDemo() {
	employeesData := make([]map[string]string, 2)

	employeesData[0] = map[string]string{
		"id":       "1",
		"name":     "sarunya",
		"position": "engineer",
	}

	employeesData[1] = map[string]string{
		"id":       "2",
		"name":     "runya",
		"position": "gineer",
	}

	// append(employeesData, employee1)

	for index, employee := range employeesData {
		fmt.Println("Employee Information ", index)
		for key, value := range employee {
			fmt.Println(key, ":", value)
		}
	}
}

func mapDemo() {
	// colors := map[string]string{
	// 	"name": "saru",
	// 	"age":  "27",
	// }
	colors := make(map[string]string)
	colors["name"] = "saru"

	fmt.Println(colors)
	delete(colors, "name")
	fmt.Println(colors)
}

func structDemo() {
	sarunya := Person{
		firstName: "Sarunya",
		lastName:  "Durai",
	}

	alex := Person{"Alex", "Durai", ContactInfo{"sar", "phone"}}

	fmt.Println(sarunya)
	fmt.Println(alex)

	sarunya.updateName("Saru")

	fmt.Println(sarunya)

	var bob Person

	fmt.Print(bob)
	fmt.Printf("%+v", bob)

	p := []Person{}
	p = append(p, sarunya)
	p = append(p, bob)
	fmt.Println(p)

	for _, person := range p {
		fmt.Println("Person is ", person.firstName)
	}
}

func (p *Person) updateName(firstName string) {
	(*p).firstName = firstName
}
