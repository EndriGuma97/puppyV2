package puppy

import (
	dog "github.com/EndriGuma97/dogoV2"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BarksInt() int {
	return dog.Dog() + 34
}

func Barks2() string {
	return "Woof! Woof! Woof123!"
}