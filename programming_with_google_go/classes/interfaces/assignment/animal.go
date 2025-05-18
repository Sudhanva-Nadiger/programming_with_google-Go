package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Animal interface {
	Eat() string
	Speak() string
	Move() string
}

type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (c *Cow) Eat() string   { return c.food }
func (c *Cow) Speak() string { return c.noise }
func (c *Cow) Move() string  { return c.locomotion }

type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (b *Bird) Eat() string   { return b.food }
func (b *Bird) Speak() string { return b.noise }
func (b *Bird) Move() string  { return b.locomotion }

type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (s *Snake) Eat() string   { return s.food }
func (s *Snake) Speak() string { return s.noise }
func (s *Snake) Move() string  { return s.locomotion }

func query(animalName string, animalAction string, mp map[string]Animal) {
	animal, ok := mp[animalName]
	if !ok {
		fmt.Println("Animal not found")
		return
	}

	switch animalAction {
	case "eat":
		fmt.Println(animal.Eat())
	case "speak":
		fmt.Println(animal.Speak())
	case "move":
		fmt.Println(animal.Move())
	default:
		fmt.Println("Invalid action")
	}
}

var allowedAnimals = []string{"cow", "bird", "snake"}
var allowedActions = []string{"eat", "speak", "move"}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mp := make(map[string]Animal)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		splits := strings.Split(input, " ")

		inputType := splits[0]

		if inputType == "newanimal" {
			animalName := splits[1]
			animalType := splits[2]

			if !slices.Contains(allowedAnimals, animalType) {
				fmt.Println("Invalid animal type")
				continue
			}

			switch animalType {
			case "cow":
				cow := &Cow{
					food:       "grass",
					locomotion: "walk",
					noise:      "moo",
					name:       animalName,
				}

				mp[animalName] = cow

			case "bird":
				bird := &Bird{
					food:       "worms",
					locomotion: "fly",
					noise:      "peep",
					name:       animalName,
				}

				mp[animalName] = bird
			case "snake":
				snake := &Snake{
					food:       "mice",
					locomotion: "slither",
					noise:      "hsss",
					name:       animalName,
				}

				mp[animalName] = snake
			}

			fmt.Println("Created it!")

		} else if inputType == "query" {
			animalName := splits[1]
			animalAction := splits[2]

			if !slices.Contains(allowedActions, animalAction) {
				fmt.Println("Invalid action")
				continue
			}

			query(animalName, animalAction, mp)
		} else {
			fmt.Println("Invalid command")
		}
	}
}
