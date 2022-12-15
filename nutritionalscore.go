package main

type ScoreType int

const(
	Food = iota
	Water
	Beverage
	Cheese
)

type NutritionalScore struct {
	value     int
	positive  int
	negative  int
	Scoretype ScoreType
}

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarLevels = []float64{45, 60, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360,270, 180, 90}
var saturatedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarLevelsBeverage = []float64{13.5, 12.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

type EnergyKJ float64

type SugarGram float64

type SaturatedFattyAcids float64

type SodiumMilligram float64

type ProteinGram float64

type FibreGram float64

type FruitsPercentage float64

type NutritionalData struct {
	Sugars  SugarGram
	Energy  EnergyKJ
	Proteins ProteinGram
	SaturatedFattyAcids SaturatedFattyAcids 
	Fibre FibreGram 
	Sodium SodiumMilligram 
	Fruits FruitsPercentage 
	Iswater bool
}

	func (e EnergyKJ)GetPoints(st ScoreType) int{
		if st == Beverage{
			return getPointsFromRange(float64(e),energyLevelsBeverage)
		}
		return getPointsFromRange(float64(e),energyLevels)
	}

	func (s SugarGram)GetPoints(st ScoreType) int {

		if st == Beverage{
			return getPointsFromRange(float64(s),sugarLevelsBeverage)
		}
		return getPointsFromRange(float64(s),energyLevels)
	}
	
	func (sfa SaturatedFattyAcids)GetPoints(st ScoreType) int {
			return getPointsFromRange(float64(sfa),saturatedFattyAcidsLevels)
	}

	func (sdm SodiumMilligram)GetPoints(st ScoreType) int {
		return getPointsFromRange(float64(sdm),sodiumLevels)
	}

	func (fp FruitsPercentage) GetPoints(st ScoreType) int {
		if st == Beverage{
			if fp > 80 {
				return 10
			} else if fp > 60 {
				return 4
			} else if fp > 40 {
				return 2
			}
			return 0
		}	
			if fp > 80 {
				return 5
			} else if fp > 60 {
				return 2
			} else if fp > 40 {
				return 1 
			}
			return 0
	}

	func (fg FibreGram) GetPoints(st ScoreType) int {
		return getPointsFromRange(float64(fg),fibreLevels)
	}

	func (p ProteinGram) GetPoints(st ScoreType) int {
		return getPointsFromRange(float64(p),proteinLevels)	
	}

	func Energyfromkcal(kcal float64) EnergyKJ{
		return EnergyKJ(kcal * 4.184)
	} 
		
	func SodiumfromSalt(saltMg float64) SodiumMilligram{
		return SodiumMilligram(saltMg/2.5)
	}

func Getnutritionalscore(n NutritionalData, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != Water{
		FruitPoint := n.Fruits.GetPoints(st)
		FibrePoint := n.Fibre.GetPoints(st)

		negative = n.Sugars.GetPoints(st) + n.Energy.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)  
		positive = FruitPoint + FibrePoint + n.Proteins.GetPoints(st)
	if st == Cheese{
		value
	}
	}

	return NutritionalScore{
	value: value,
	positive: positive,
	negative: negative,
	Scoretype: st,
}}

func (ns NutritionalScore) GetNutriScore() string {

}

func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)
	for i, l:= range steps {
		if v > l {
			return lenSteps - i
		}
	}
	return 0;
}
