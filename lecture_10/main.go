package main

import "fmt"

type dog struct {
	needEat   int
	forWeight int
	weightDog float64
}

func (c dog) String() string {
	return fmt.Sprintf("dog weighing %.3f kilograms", c.weightDog)
}

func (c dog) getFoodForWeight() float64 {
	return needFoodForWeight(c.needEat, c.forWeight, c.weightDog)
}

type cat struct {
	needEat   int
	forWeight int
	weightCat float64
}

func (c cat) String() string {
	return fmt.Sprintf("cat weighing %.3f kilograms", c.weightCat)
}

func (c cat) getFoodForWeight() float64 {
	return needFoodForWeight(c.needEat, c.forWeight, c.weightCat)
}

type cow struct {
	needEat   int
	forWeight int
	weightCow float64
}

func (c cow) String() string {
	return fmt.Sprintf("cow weighing %.3f kilograms", c.weightCow)
}

func (c cow) getFoodForWeight() float64 {
	return needFoodForWeight(c.needEat, c.forWeight, c.weightCow)
}

type animals interface {
	nedEatForOneAnimal
	fmt.Stringer
}

type nedEatForOneAnimal interface {
	getFoodForWeight() float64
}

func needFoodForWeight(needEat, forWeight int, weight float64) float64 {
	return float64(needEat/forWeight) * weight
}

func main() {
	var countOllFarmFood float64
	var countOllAnimals int

	c := []animals{
		dog{needEat: 10, forWeight: 5, weightDog: 12.231},
		cat{needEat: 7, forWeight: 1, weightCat: 3.123},
		cow{needEat: 25, forWeight: 1, weightCow: 96.122},
		cat{needEat: 4, forWeight: 2, weightCat: 6.266},
		dog{needEat: 3, forWeight: 1, weightDog: 12.231},
		cow{needEat: 63, forWeight: 3, weightCow: 96.122},
	}

	for _, res := range c {
		fmt.Printf("To feed a %s, you need %.3f kilograms of food per month.\n", res.String(), res.getFoodForWeight())

		countOllFarmFood += res.getFoodForWeight()
		countOllAnimals++
	}
	fmt.Printf("Total per month yoo need %.3f kilograms of food, for %d animals.", countOllFarmFood, countOllAnimals)
}
