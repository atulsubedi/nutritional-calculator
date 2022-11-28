package main

	func main() {
		ns := Getnutritionalscore(NutritionalData){
			Energy : Energyfromkcal(0),
			Sugars : Sugargram(10),
			SaturatedFattyAcids : SaturatedFattyAcids(4),
			Sodium : SodiumMilligram(4), //TODO fix this
			Fruits : FruitsPercentage(9),
			Fibre : FibreGram(8),
			Protein : ProteinGram(7),
		}
	}

