package service

// import (
// 	"reflect"
// )

// var podSizePointsLookup = map[int]map[int]int{
// 	3: {1: 4, 2: 3, 3: 2, 4: 0, 5: 0, 6: 1},
// 	4: {1: 5, 2: 4, 3: 3, 4: 2, 5: 0, 6: 1},
// 	5: {1: 6, 2: 5, 3: 4, 4: 3, 5: 2, 6: 1},
// }

// func calculatePoints(result Result) int {
// 	var points int

// 	if podSizePoints, ok := podSizePointsLookup[result.PodSize]; ok {
// 		if totalPoints, ok := podSizePoints[result.Place]; ok {
// 			points = totalPoints
// 		}
// 	}

// 	return points
// }

// var bonusPointsLookup = map[string]int{
// 	"TheCouncilOfWizards":    1,
// 	"DavidAndTheGoliaths":    1,
// 	"Untouchable":            3,
// 	"Cleave":                 1,
// 	"ItsFreeRealEstate":      1,
// 	"IAmTimmy":               1,
// 	"BigBiggerHuge":          1,
// 	"CloseButNoCigar":        2,
// 	"JustAsGarfieldIntended": 1,
// }

// func calculateBonusPoints(result Result) int {
// 	var bonusPoints int

// 	// Iterate through boolean fields and add bonus points
// 	rv := reflect.ValueOf(result)
// 	for i := 0; i < rv.NumField(); i++ {
// 		field := rv.Field(i)
// 		if field.Kind() == reflect.Bool && field.Bool() {
// 			fieldName := rv.Type().Field(i).Name
// 			if bonus, ok := bonusPointsLookup[fieldName]; ok {
// 				bonusPoints += bonus
// 			}
// 		}
// 	}

// 	return bonusPoints
// }
