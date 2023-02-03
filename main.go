package main

import (
	"fmt"

	"example.com/codeTSI/jobCalculations"
	"example.com/codeTSI/jobInfo"
)

func main() {

	//How many walls would you like to paint?
	numberOfWalls := jobInfo.GetNumberOfWalls()

	//List the dimensions of each wall in meters squared
	setOfWalls := jobInfo.GetWallDimensions(numberOfWalls)

	//Do any of these walls have windows/doors
	numberOfWindowsDoors := jobInfo.GetNumberOfWindowsDoors()

	//List the dimensions of each wall in meters squared
	setOfWindowsDoors := jobInfo.GetWindowsDoorsDimensions(numberOfWindowsDoors)

	//Calculate total meters squared to paint
	surfaceAreaToPaint, err := jobCalculations.GetSurfaceAreToPaint(setOfWalls, setOfWindowsDoors)

	if err != nil {
		fmt.Println(err)
	} else {
		//What is the price per litre of paint you would like to use
		priceOfPaintPerLitre := jobInfo.GetPriceOfPaintPerLitre()

		//Calculate total price to paint walls
		priceToPaintWalls, litresToPaintWalls := jobCalculations.GetTotalPriceToPaintWalls(surfaceAreaToPaint, priceOfPaintPerLitre)

		fmt.Printf("\nThe total price to paint %d walls is %.2f and will take ~%.2f litres of paint\n", numberOfWalls, priceToPaintWalls, litresToPaintWalls)
	}

}

// fmt.Println("Number of walls is ", numberOfWalls)
// fmt.Println("Set of walls is ", setOfWalls)
// fmt.Println("Number of windows and doors is ", numberOfWindowsDoors)
// fmt.Println("Set of windows and doors is ", setOfWindowsDoors)
// fmt.Println("Surface area to be painted is ", surfaceAreaToPaint)
// fmt.Println("The price of paint per litre is ", priceOfPaintPerLitre)
