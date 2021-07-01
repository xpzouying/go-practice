package main

import (
	"log"
	"unsafe"
)

func bytesToString() {

	data := []byte("zouying")

	// 强制转换为 unsafe.Pointer
	p := unsafe.Pointer(&data)
	name := *(*string)(p)
	_ = name
}

// 在struct中转换
func studentToPerson() {

	type student struct {
		name string
		age  int

		profession string
		major      string
	}

	type person struct {
		name string
		age  int
	}

	s1 := student{
		name:       "eden",
		age:        18,
		profession: "CS",
		major:      "software",
	}

	p := unsafe.Pointer(&s1)

	person1 := *(*person)(p)

	log.Printf("student: %+v", s1)
	log.Printf("person: %+v", person1)
}

func main() {
	bytesToString()

	studentToPerson()
}
