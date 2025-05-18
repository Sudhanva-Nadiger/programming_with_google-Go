package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string   { return a.food }
func (a *Animal) Speak() string { return a.noise }
func (a *Animal) Move() string  { return a.locomotion }

// Inheritance
type Cow struct {
	Animal
}

type Bird struct {
	Animal
}

type Snake struct {
	Animal
}

var allowedAnimals = []string{"cow", "bird", "snake"}
var allowedActions = []string{"eat", "speak", "move"}

func main() {
	cow := Cow{
		Animal: Animal{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
	}
	bird := Bird{
		Animal: Animal{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
	}
	snake := Snake{
		Animal: Animal{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		scanner.Scan()
		input := scanner.Text()
		splits := strings.Split(input, " ")

		animal := splits[0]
		action := splits[1]

		if !slices.Contains(allowedAnimals, animal) {
			fmt.Println("Invalid animal")
			continue
		}

		if !slices.Contains(allowedActions, action) {
			fmt.Println("Invalid action")
			continue
		}

		switch animal {
		case "cow":
			switch action {
			case "speak":
				fmt.Println(cow.Speak())
			case "move":
				fmt.Println(cow.Move())
			case "eat":
				fmt.Println(cow.Eat())
			}
		case "bird":
			switch action {
			case "speak":
				fmt.Println(bird.Speak())
			case "move":
				fmt.Println(bird.Move())
			case "eat":
				fmt.Println(bird.Eat())
			}
		case "snake":
			switch action {
			case "speak":
				fmt.Println(snake.Speak())
			case "move":
				fmt.Println(snake.Move())
			case "eat":
				fmt.Println(snake.Eat())
			}
		}
	}
}
