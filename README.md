# Meta2k24-GoLang-Project

### Folder Structure

```bash
Student-DB
├── main.go
├── db.txt
└── features
    └── features.go
```
---

### main.go starter code

```go 
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello, World!")
}

```
---

### features.go 

#### Structs

```go
package features

//Note these are final dependencies and if they give any erro just comment unused ones for time being.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Student struct {
	Name   string
	RollNo string
	Marks  int
}

type Class struct {
	Engineers []Student
}

```
---

#### Taking input from user
  
  ```go
  func inputName() string {
	fmt.Println("Enter your first name")
	var name string
	fmt.Scanln(&name)
	return name
}

func inputRoll() string {
	fmt.Println("Enter your roll no")
	var roll string
	fmt.Scanln(&roll)
	return roll
}

func inputMarks() int {
	fmt.Println("Enter your marks")
	var marks int
	fmt.Scanln(&marks)
	return marks
}
```
--- 

#### New Class Function

```go
func (c *Class) NewClass() {

	// fmt.Println(entry)
	var studentsArray []Student
	file, err := os.Open("db.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		marks, _ := strconv.Atoi(fields[2])

		entry := Student{
			Name:   fields[0],
			RollNo: fields[1],
			Marks:  marks,
		}

		studentsArray = append(studentsArray, entry)

	}

	// fmt.Println(studentsArray)
	c.Engineers = studentsArray

}
```
---

#### Add Student Function

```go
func (c *Class) AddStudent() {

	name := inputName()

	roll := inputRoll()

	marks := inputMarks()

	entry := Student{
		Name:   name,
		RollNo: roll,
		Marks:  marks,
	}

	c.Engineers = append(c.Engineers, entry)

	fmt.Println("Student Added Successfully")
	fmt.Println()

	//Update this data in the db.txt
}
```
---

#### Show Students Function

```go
func (c *Class) ShowStudents() {
	for _, student := range c.Engineers {
		name, roll, marks := student.Name, student.RollNo, student.Marks
		fmt.Printf("Name: %s\nRoll No: %s\nMarks: %d\n", name, roll, marks)
		fmt.Println()
	}
}
```
---

#### Update Student Function

```go
func (c *Class) UpdateStudent() {

	roll := inputRoll()
	// find the student
	index := 0
	for ind, stu := range c.Engineers {
		if stu.RollNo == roll {
			index = ind
			break
		}
	}
	// get the data
	newName := inputName()
	newMarks := inputMarks()
	// change the data

	c.Engineers[index] = Student{
		RollNo: roll,
		Name:   newName,
		Marks:  newMarks,
	}
}

```
---
#### Delete Student Function
```go
func (c *Class) DeleteStudent() {

	// get the data
	roll := inputRoll()
	index := 0

	// find the index
	for idx, student := range c.Engineers {
		if student.RollNo == roll {
			index = idx
		}
	}

	// remove the student
	c.Engineers = append(c.Engineers[:index], c.Engineers[index+1:]...)
}
```
---
#### Write To File Function
```go
func (c *Class) writeToFile(wg *sync.WaitGroup) {

	// convert []Student into string
	var result []string

	for _, val := range c.Engineers {
		current := fmt.Sprintf("%v %v %v", val.Name, val.RollNo, val.Marks)
		result = append(result, current)
	}


	err := os.WriteFile("db.txt", []byte(strings.Join(result, "\n")), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
```
