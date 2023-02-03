package jobInfo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetNumberOfWalls() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Oops:", r, "Please restart this program to try again")
			os.Exit(8)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	var numberOfWalls int64

	fmt.Println("How many walls would you like to paint?")
	for scanner.Scan() {
		nStr := scanner.Text()
		firstNumber, err := strconv.ParseInt(nStr, 10, 64)

		if err != nil {
			fmt.Printf("%s is not a whole number.\n", nStr)
			fmt.Println("How many walls would you like to paint?")
			continue
		} else {
			if firstNumber == 0 {
				panic("There are no walls to paint!")
			} else if firstNumber < 0 {
				fmt.Println("Oops! That's a negative number, we'll assume you meant the positive number :)")
				numberOfWalls = -firstNumber
			} else {
				numberOfWalls = firstNumber
			}

			break
		}
	}

	return int(numberOfWalls)
}

func GetWallDimensions(numberOfWalls int) []float64 {
	scanner := bufio.NewScanner(os.Stdin)

	setOfWalls := []float64{}

	for i := 0; i < numberOfWalls; i++ {
		fmt.Printf("What size (meter squared) is wall number %d\n", i+1)
		scanner.Scan()
		nStr := scanner.Text()
		wallSize, err := strconv.ParseFloat(nStr, 64)

		if err != nil {
			fmt.Printf("%s is not a number.\n", nStr)
			i--
		} else {
			if wallSize < 0 {
				fmt.Println("Oops! That's a negative wall size, we'll assume you meant the positive equivalent :)")
				setOfWalls = append(setOfWalls, -wallSize)
			} else {
				setOfWalls = append(setOfWalls, wallSize)
			}
		}
	}

	return setOfWalls
}

func GetNumberOfWindowsDoors() (err int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Oops:", r)
			err = 0
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	var numberOfWindowsDoors int64

	fmt.Println("How many doors/windows are on these walls")
	for scanner.Scan() {
		nStr := scanner.Text()
		firstNumber, err := strconv.ParseInt(nStr, 10, 64)

		if err != nil {
			fmt.Printf("%s is not a number.\n", nStr)
			fmt.Println("How many doors/windows are on these walls")
			continue
		} else {
			if firstNumber < 0 {
				panic("Input is less than 0. We will assume there are no windows/doors :)\n")
			}
			numberOfWindowsDoors = firstNumber
			break
		}
	}

	return int(numberOfWindowsDoors)
}

func GetWindowsDoorsDimensions(numberOfWalls int) []float64 {
	scanner := bufio.NewScanner(os.Stdin)

	setOfWindowsDoors := []float64{}

	for i := 0; i < numberOfWalls; i++ {
		fmt.Printf("What size (meter squared) is door/window number %d\n", i+1)
		scanner.Scan()
		nStr := scanner.Text()
		wallSize, err := strconv.ParseFloat(nStr, 64)

		if err != nil {
			fmt.Printf("%s is not a number.\n", nStr)
			i--
		} else {
			setOfWindowsDoors = append(setOfWindowsDoors, wallSize)
		}
	}

	return setOfWindowsDoors
}

func GetPriceOfPaintPerLitre() float64 {
	scanner := bufio.NewScanner(os.Stdin)
	var priceOfPaint float64

	fmt.Println("What is the price per litre of paint you would like to use?")
	for scanner.Scan() {
		nStr := scanner.Text()
		firstNumber, err := strconv.ParseFloat(nStr, 64)

		if err != nil {
			fmt.Printf("%s is not a valid price.\n", nStr)
			fmt.Println("What is the price per litre of paint you would like to use?")
			continue
		} else {
			priceOfPaint = firstNumber
			break
		}
	}

	return priceOfPaint
}
