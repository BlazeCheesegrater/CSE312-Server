package util

type Person struct{
	Name String
	Age int
}

func (p Person) SayHello() string {
	return "Hello " + p
}
