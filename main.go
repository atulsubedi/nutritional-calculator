package main

import "fmt"

func main() {
	ns := Getnutritionalscore(NutritionalData{
		Energy:              Energyfromkcal(0),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(4),
		Sodium:              SodiumMilligram(4), //TODO fix this
		Fruits:              FruitsPercentage(9),
		Fibre:               FibreGram(8),
		Proteins:            ProteinGram(7),
	}, Food)

	fmt.Printf("Nutritional Score:%d\n", ns.value)
	
	// The below getnutriscore() will give you the grade of your food intake
	fmt.Printf("Nutri Score: %s\n", ns.GetNutriScore())
}
