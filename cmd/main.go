package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/toadFrogski/ASDC03/pkg/collections"
	. "github.com/toadFrogski/ASDC03/pkg/student"
)

func main() {
	f, err := os.Open("./data/students.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	byteArray, err := ioutil.ReadAll(f)
	var students []Student
	json.Unmarshal(byteArray, &students)

	var studentsList ListInterface
	studentsList = &List{}
	for i := 0; i < len(students); i++ {
		studentsList.Add(students[i])
	}
	var studentsSequence SequencesInterface
	studentsSequence = &Queue{}
	for i := 0; i < len(students); i++ {
		studentsSequence.Push(students[i])
	}
	for i := 0; i < len(students); i++ {
		fmt.Printf("Removed : %v\n", studentsSequence.Pop())
	}
	studentsSequence.Display()

	studentsSequence = &Stack{}
	for i := 0; i < len(students); i++ {
		studentsSequence.Push(students[i])
	}
	for i := 0; i < len(students); i++ {
		fmt.Printf("Removed : %v\n", studentsSequence.Pop())
	}

	bst := Bst{}
	for i := 0; i < len(students); i++ {
		bst.Insert([]byte(fmt.Sprint(students[i])))
	}
	bst.Display()
	bst.DisplayReverso()
	bst.DisplaySymmetrical()
}
