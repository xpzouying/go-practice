package main

import (
	"fmt"
	"os"
)

type person struct {
	Age         int
	Name        string
	Description string
}

type option interface {
	Config(*person) error
}

type optionFn func(*person) error

func (o optionFn) Config(p *person) error {
	return o(p)
}

func withName(name string) option {
	return optionFn(func(p *person) error {
		p.Name = name
		return nil
	})
}

func withAge(age int) option {
	return optionFn(
		func(p *person) error {
			p.Age = age
			return nil
		})
}

func newPerson(opts ...option) (*person, error) {
	p := person{}
	for _, opt := range opts {
		if err := opt.Config(&p); err != nil {
			fmt.Fprintf(os.Stderr, "config person error: %v", err)
			return nil, err
		}
	}

	return &p, nil
}

func main() {
	p, err := newPerson(withName("zouying"), withAge(32))
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("person: %v\n", p)
}
