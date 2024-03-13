package features

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Student struct {
	Name   string
	RollNo string
	Marks  int
}

type Class struct {
	Engineers []Student
}

func (c *Class) NewClass() {
	var studentArray []Student
	file, err := os.Open("db.txt")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		roll := fields[0]
		name := fields[1]
		marks, err := strconv.Atoi(fields[2])

		if err != nil {
			panic(err)
		}
		entry := Student{
			Name:   name,
			RollNo: roll,
			Marks:  marks,
		}

		studentArray = append(studentArray, entry)
	}

	c.Engineers = studentArray

}

//CRUD

func inputName() string {
	var name string
	fmt.Println("Enter your first name")
	fmt.Scanln(&name)
	return name
}

func inputRoll() string {
	var roll string
	fmt.Println("Enter your roll no")
	fmt.Scanln(&roll)
	return roll
}

func inputMarks() int {
	var marks int
	fmt.Println("Enter Your Marks")
	fmt.Scanln(&marks)
	return marks
}

func (c *Class) AddStudent() {
	var wg sync.WaitGroup
	name := inputName()
	roll := inputRoll()
	marks := inputMarks()
	entry := Student{
		Name:   name,
		RollNo: roll,
		Marks:  marks,
	}

	c.Engineers = append(c.Engineers, entry)

	wg.Add(1)
	c.writeToFile(&wg)
	wg.Wait()
}

func (c *Class) ShowStudents() {
	for _, val := range c.Engineers {
		fmt.Printf("Name=%s Roll No=%s Marks=%d\n", val.Name, val.RollNo, val.Marks)
	}
}

func (c *Class) DeleteStudent() {
	var wg sync.WaitGroup
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
	// shivam utkarsh aditya
	//    0    1        2
	// 1
	// shivam aditya
	c.Engineers = append(c.Engineers[:index], c.Engineers[index+1:]...)

	wg.Add(1)
	c.writeToFile(&wg)
	wg.Wait()
}

func (c *Class) UpdateStudent() {

	var wg sync.WaitGroup

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

	wg.Add(1)
	go c.writeToFile(&wg)

	wg.Wait()
}

func (c *Class) writeToFile(wg *sync.WaitGroup) {

	defer wg.Done()

	// convert []Student into string
	var result []string

	for _, val := range c.Engineers {
		// {roll: 1, name:utkarsh, marks:404}
		// "utkasrsh 1 404"
		current := fmt.Sprintf("%v %v %v", val.RollNo, val.Name, val.Marks)
		result = append(result, current)
	}

	// read write execute
	// 1     1      1
	//  owner group other
	// 6 6 4
	// 100

	err := os.WriteFile("db.txt", []byte(strings.Join(result, "\n")), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
