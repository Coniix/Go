package jobCalculations

import "errors"

func GetSurfaceAreToPaint(setOfWalls []float64, setOfWindowsDoors []float64) (float64, error) {

	wallSurfaceArea := 0.0
	windowDoorSurfaceArea := 0.0

	for _, wallSize := range setOfWalls {
		wallSurfaceArea += wallSize
	}

	for _, windowDoorSize := range setOfWindowsDoors {
		windowDoorSurfaceArea += windowDoorSize
	}

	if windowDoorSurfaceArea > wallSurfaceArea {
		return 0, errors.New("Your window/door surface area is greater than your wall surface area. There is no wall space to paint.")
	}

	return wallSurfaceArea - windowDoorSurfaceArea, nil
}

func GetTotalPriceToPaintWalls(surfaceAreaToPaint float64, priceOfPaintPerLitre float64) (float64, float64) {
	//3.785 litres covers 37.16 square meters
	//1 litre covers 9.82 square meters
	litresOfPaint := surfaceAreaToPaint / 9.82

	return litresOfPaint * priceOfPaintPerLitre, litresOfPaint
}
