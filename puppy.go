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

func From11() string {
	ver := "im from version 1.1.0"
	return ver
}

func From12() string {
	return "im from version 1.2.0"
}

func From13() string {
	return "im from version 1.3.0"
}
