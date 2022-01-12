package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type Zone struct {
	X int
	Y int
}

type Rover struct {
	CurrentX         int
	CurrentY         int
	CurrentDirection string
	Zone             Zone
}

func main() {
	resp := "y"
	for {
		if resp == "y" {
			zone := ReadZoneDimensions()

			xStart, yStart := zone.ReadRoverStartingCoOrdinates()

			startingDirection := ReadStartingDirection()

			rover := Rover{xStart, yStart, startingDirection, zone}

			move := ReadRoverMoveOrders()

			rover.MoveRover(move)

			log.Println("--------Rover ending co-ordinates--------")
			log.Printf("X Axis: %v", rover.CurrentX)
			log.Printf("Y Axis: %v", rover.CurrentY)
			log.Printf("Direction facing: %v", strings.ToUpper(rover.CurrentDirection))
			scn := bufio.NewScanner(os.Stdin)
			fmt.Printf("Do you want to contine?")
			fmt.Printf(" (y/n): ")
			scn.Scan()
			answer := scn.Text()
			fmt.Printf("Your response: %q \n", answer)
			if answer == "y" {
				resp = "y"
			} else if answer == "n" {
				resp = "n"
			} else {
				scn := bufio.NewScanner(os.Stdin)
				fmt.Printf("Do you want to contine?")
				fmt.Printf(" (y/n): ")
				scn.Scan()
				ret := scn.Text()
				fmt.Printf("Your response: %q \n", ret)
			}
		} else if resp == "n" {
			os.Exit(0)
		} else {
			log.Panic()
		}
	}
}

func (r *Rover) MoveRover(movement string) {

	for _, m := range movement {

		if strings.EqualFold(r.CurrentDirection, "N") {
			if strings.EqualFold(string(m), "M") {
				if r.CurrentX < r.Zone.X {
					r.CurrentX++
					continue
				} else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover has exited current zone. Signal lost!"))
				}
			} else if strings.EqualFold(string(m), "R") {
				r.CurrentDirection = "E"
				continue
			} else if strings.EqualFold(string(m), "L") {
				r.CurrentDirection = "W"
				continue
			}
		}
		if strings.EqualFold(r.CurrentDirection, "E") {
			if strings.EqualFold(string(m), "M") {
				if r.CurrentY < r.Zone.Y {
					r.CurrentY++
					continue
				} else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover has exited current zone. Signal lost!"))
				}
			} else if strings.EqualFold(string(m), "R") {
				r.CurrentDirection = "S"
				continue
			} else if strings.EqualFold(string(m), "L") {
				r.CurrentDirection = "N"
				continue
			}
		}
		if strings.EqualFold(r.CurrentDirection, "S") {
			if strings.EqualFold(string(m), "M") {
				if r.CurrentX > 0 {
					r.CurrentX--
					continue
				} else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover is leaving current zone. Signal lost!"))
				}
			} else if strings.EqualFold(string(m), "R") {
				r.CurrentDirection = "W"
				continue
			} else if strings.EqualFold(string(m), "L") {
				r.CurrentDirection = "E"
				continue
			}
		}
		if strings.EqualFold(r.CurrentDirection, "W") {
			if strings.EqualFold(string(m), "M") {
				if r.CurrentX > 0 {
					r.CurrentY--
					continue
				} else {
					log.Println("Rover is leaving current zone. Signal lost!")
					panic(errors.New("Rover is leaving current zone. Signal lost!"))
				}
			} else if strings.EqualFold(string(m), "R") {
				r.CurrentDirection = "N"
				continue
			} else if strings.EqualFold(string(m), "L") {
				r.CurrentDirection = "S"
				continue
			}
		}
	}
}

func ReadZoneDimensions() Zone {

	var x int
	var y int

	log.Println("Please enter Zone X axis length:")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println(err)
	}

	log.Println("Please enter Zone Y axis length:")
	_, err = fmt.Scanf("%d", &y)
	if err != nil {
		fmt.Println(err)
	}

	zone := Zone{x, y}

	return zone
}

func (z *Zone) ReadRoverStartingCoOrdinates() (int, int) {
	var x int
	var y int

	for {
		log.Println("Please enter Rover starting X axis:")
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			fmt.Println(err)
		}
		if x > z.X {
			log.Printf("This point is outside of the current zone X axis of: %v", z.X)
			continue
		}
		break
	}

	for {
		log.Println("Please enter Rover starting Y axis:")
		_, err := fmt.Scanf("%d", &y)
		if err != nil {
			fmt.Println(err)
		}
		if y > z.Y {
			log.Printf("This point is outside of the current zone Y axis of: %v", z.Y)
			continue
		}
		break
	}
	return x, y
}

func ReadStartingDirection() string {
	var direction string
	reader := bufio.NewReader(os.Stdin)

	for {
		log.Println("Please enter Rover starting direction:")

		direction, _ = reader.ReadString('\n')
		direction = strings.TrimSuffix(direction, "\n")

		if !ConfirmRoverStartingDirectionValidity(direction) {
			log.Println("Rover staring direction not correct. Only 'N', 'E', 'S' or 'W' permitted")
			continue
		}
		break
	}
	return direction
}

func ReadRoverMoveOrders() string {

	var move string
	reader := bufio.NewReader(os.Stdin)

	for {
		log.Println("Please enter Rover movement orders:")

		move, _ = reader.ReadString('\n')
		move = strings.TrimSuffix(move, "\n")

		if !ConfirmRoverMoveOrdersValidity(move) {
			log.Println("Rover movement order input is not correct. Only 'M', 'R' or 'L' permitted")
			continue
		}
		break
	}
	return move
}

func ConfirmRoverMoveOrdersValidity(orders string) bool {
	const alpha = "mlr"

	for _, char := range orders {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

func ConfirmRoverStartingDirectionValidity(direction string) bool {
	const direc = "nesw"

	for _, char := range direction {
		if !strings.Contains(direc, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
