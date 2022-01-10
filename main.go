package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter size of plateau")
	fmt.Printf(" (ex: 5 5): ")
	scanner.Scan()
	size := scanner.Text()
	fmt.Printf("Enter size of plateau: %q \n", size)

	resp := "y"
	x := 1
	for x < 2 {
		if resp == "y" {

			scan := bufio.NewScanner(os.Stdin)
			fmt.Printf("Enter location of rover")
			fmt.Printf(" (ex: 1 2 N): ")
			scan.Scan()
			location := scan.Text()
			fmt.Printf("Size of plateau: %q \n", location)

			sc := bufio.NewScanner(os.Stdin)
			fmt.Printf("Enter direction of rover")
			fmt.Printf(" (ex: LMRMMRMRM): ")
			sc.Scan()
			direction := sc.Text()
			fmt.Printf("Direction of plateau: %q \n", direction)

			scn := bufio.NewScanner(os.Stdin)
			fmt.Printf("Do you want to contine?")
			fmt.Printf(" (y/n): ")
			scn.Scan()
			ret := scn.Text()
			fmt.Printf("Your response: %q \n", ret)
			if ret == "y" {
				resp = "y"
			} else if ret == "n" {
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
			fmt.Printf("Thank you!\n")
			os.Exit(0)
		} else {
			fmt.Printf("Something went wrong!")
		}

	}
}
