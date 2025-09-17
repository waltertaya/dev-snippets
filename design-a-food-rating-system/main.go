package main

import (
	"fmt"
	"strings"
)

type Food struct {
	Food     string
	Cousines string
	Ratings  int
}

type FoodRatings struct {
	Foods []Food
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {

	var result []Food

	for i := 0; i < len(foods); i++ {
		food := Food{
			foods[i],
			cuisines[i],
			ratings[i],
		}
		result = append(result, food)
	}

	return FoodRatings{result}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	for i, f := range this.Foods {
		if f.Food == food {
			this.Foods[i].Ratings = newRating
			break
		}
	}
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	var result string
	var rating int = -1 // start lower than any valid rating

	for _, f := range this.Foods {
		if f.Cousines == cuisine {
			if f.Ratings > rating ||
				(f.Ratings == rating && strings.Compare(f.Food, result) < 0) {
				rating = f.Ratings
				result = f.Food
			}
		}
	}
	return result
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
func main() {

	// fmt.Println('r' > 's')

	// Test 1
	foods := []string{
		"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi",
	}
	cuisines := []string{
		"korean", "japanese", "japanese", "greek", "japanese", "korean",
	}
	ratings := []int{
		9, 12, 8, 15, 14, 7,
	}

	obj := Constructor(foods, cuisines, ratings)
	fmt.Println(obj.HighestRated("korean"))   // kimchi
	fmt.Println(obj.HighestRated("japanese")) // ramen
	obj.ChangeRating("sushi", 16)
	fmt.Println(obj.HighestRated("japanese")) // sushi
	obj.ChangeRating("ramen", 16)
	fmt.Println(obj.HighestRated("japanese")) // ramen

	// Test 2
	foods = []string{
		"cpctxzh", "bryvgjqmj", "wedqhqrmyc", "ee", "lafzximxh", "lojzxfel", "flhs",
	}
	cuisines = []string{
		"wbhdgqphq", "wbhdgqphq", "mxxajogm", "wbhdgqphq", "wbhdgqphq", "mxxajogm", "mxxajogm",
	}
	ratings = []int{
		15, 5, 7, 16, 16, 10, 13,
	}

	obj = Constructor(foods, cuisines, ratings)
	obj.ChangeRating("lojzxfel", 1)
	fmt.Println(obj.HighestRated("mxxajogm"))  // flhs
	fmt.Println(obj.HighestRated("wbhdgqphq")) // ee
	fmt.Println(obj.HighestRated("mxxajogm"))  // flhs
}
