package puppy

import (
	"github.com/sagivO2150/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func from11() string {
	ver := "im from version 1.1.0"
	return ver
}

func from12() string {
	return "im from version 1.2.0"
}
