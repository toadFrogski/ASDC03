package collections_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/toadFrogski/ASDC03/pkg/collections"
	. "github.com/toadFrogski/ASDC03/pkg/student"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestList(t *testing.T) {
	students, err := readStudentsFromJson()

	if err != nil {
		t.Fatalf("\t%s\tFailed to load fixtures for tests", failed)
	}
	list := &List{}
	t.Log("Given the need to check add function")
	t.Logf("When object add to list")
	{
		list.Add(students[0])
		if list.Get(0) != students[0] {
			t.Fatalf("\t%s\tElement was not added", failed)
		}
	}
	t.Logf("When try to find object in list that exist")
	{
		list.Add(students[1])
		list.Add(students[2])
		if list.Find(students[0]) != students[0] {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
		if list.Find(students[1]) != students[1] {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
		if list.Find(students[2]) != students[2] {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
	}
	t.Logf("When try to find object in list that not exist")
	{
		if list.Find(students[4]) != nil {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
	}
	t.Logf("When try to remove object in list that exist")
	{
		list.Remove(students[3])
		if list.Find(students[3]) != nil {
			t.Fatalf("\t%s\tElement was not removed", failed)
		}
	}
	t.Logf("\t%s\tList test passed successfully", success)
}

func TestQueue(t *testing.T) {
	students, err := readStudentsFromJson()

	if err != nil {
		t.Fatalf("\t%s\tFailed to load fixtures for tests", failed)
	}
	queue := &Queue{}
	t.Logf("When object add to queue")
	{
		queue.Push(students[0])
		if queue.Find(students[0]) != students[0] {
			t.Fatalf("\t%s\tElement was not added", failed)
		}
	}
	t.Logf("When try to find object in queue that exist")
	{
		queue.Push(students[1])
		queue.Push(students[2])
		if queue.Find(students[0]) != students[0] {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
		if queue.Find(students[1]) != students[1] {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
		if queue.Find(students[2]) != students[2] {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
	}
	t.Logf("When try to find object in queue that not exist")
	{
		if queue.Find(students[4]) != nil {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
	}
	t.Logf("When try to pop object in queue that exist")
	{
		student := queue.Pop()
		if queue.Find(students[0]) != nil {
			t.Fatalf("\t%s\tElement was not popped", failed)
		}
		if student != students[0] {
			t.Fatalf("%s\tPop method illegal behavior", failed)
		}
	}
	t.Logf("\t%s\tQueue test passed successfully", success)
}

func TestStack(t *testing.T) {
	students, err := readStudentsFromJson()

	if err != nil {
		t.Fatalf("\t%s\tFailed to load fixtures for tests", failed)
	}
	stack := &Stack{}
	t.Logf("When object add to stack")
	{
		stack.Push(students[0])
		if stack.Find(students[0]) != students[0] {
			t.Fatalf("\t%s\tElement was not added", failed)
		}
	}
	t.Logf("When try to find object in stack that exist")
	{
		stack.Push(students[1])
		stack.Push(students[2])
		if stack.Find(students[0]) != students[0] {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
		if stack.Find(students[1]) != students[1] {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
		if stack.Find(students[2]) != students[2] {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
	}
	t.Logf("When try to find object in stack that not exist")
	{
		if stack.Find(students[4]) != nil {
			t.Fatalf("\t%s\tFind method illegal behavior", failed)
		}
	}
	t.Logf("When try to pop object in stack that exist")
	{
		student := stack.Pop()
		if stack.Find(students[2]) != nil {
			t.Fatalf("\t%s\tElement was not popped", failed)
		}
		if student != students[2] {
			t.Fatalf("%s\tPop method illegal behavior", failed)
		}
	}
	t.Logf("\t%s\tStack test passed successfully", success)
}

func TestBst(t *testing.T) {
	students, err := readStudentsFromJson()

	if err != nil {
		t.Fatalf("\t%s\tFailed to load fixtures for tests", failed)
	}
	bst := &Bst{}
	t.Logf("When object add to stack")
	{
		s := []byte(fmt.Sprint(students[0]))
		bst.Insert(s)
		if !bst.Find(s) {
			t.Fatalf("\t%s\tInsert or Find illegal behavior", failed)
		}
	}
	t.Logf("\t%s\tBst test passed successfully", success)
}

func readStudentsFromJson() ([]Student, error) {
	f, err := os.Open("../data/students.json")

	defer f.Close()

	byteArray, err := ioutil.ReadAll(f)
	var students []Student
	json.Unmarshal(byteArray, &students)
	return students, err
}
