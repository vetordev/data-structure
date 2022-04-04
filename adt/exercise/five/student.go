package exercise

import (
	"fmt"
	"math/rand"
	"time"
)

var studentMap map[int]*Student = make(map[int]*Student)

type Student struct {
	name      string
	registry  string
	birthDate time.Time
}

func insertStudent(name string, registry string, birthDate time.Time) *Student {
	student := &Student{name, registry, birthDate}

	id := rand.Int()
	studentMap[id] = student

	return student
}

func getStudent(id int) *Student {
	return studentMap[id]
}

func deleteStudent(id int) {
	delete(studentMap, id)
}

func showStudentByBirth(date *time.Time) {
	for id, student := range studentMap {
		if student.birthDate.Before(*date) {
			fmt.Printf("%d ---> %v\n", id, &student)
		}
	}
}

func showStudentByRegistry(registry string) {
	for id, student := range studentMap {
		if student.registry == registry {
			fmt.Printf("%d ---> %v\n", id, &student)
			break
		}
	}
}
