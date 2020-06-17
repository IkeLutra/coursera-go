package main

import (
	"testing"
)

func TestAnimal(t *testing.T) {
	animal := Animal{
		food:       "hay",
		locomotion: "hop",
		noise:      "thump",
	}
	if animal.food != "hay" {
		t.Errorf("not hay")
	}
	if animal.locomotion != "hop" {
		t.Errorf("Not hop")
	}
	if animal.noise != "thump" {
		t.Errorf("Not thump")
	}
	if animal.Eat() != "hay" {
		t.Errorf("This animal does not eat hay: %v", animal.Eat())
	}
	if animal.Move() != "hop" {
		t.Errorf("This animal does not hop: %v", animal.Move())
	}
	if animal.Speak() != "thump" {
		t.Errorf("This animal does not thump: %v", animal.Speak())
	}

}
