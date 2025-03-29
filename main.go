package main

import (
	"net/http"
	"strconv"
)

// To run this program use "go run main.go", in browser type localhost:8080 and press enter
type foodItem struct {
	name string
	url  string
}

// Add new food items here.
var mealMap = map[string][]foodItem{
	"Breakfast": {{name: "Aval upma", url: "breakfast/Aval%20upma.html"}, {name: "Appam", url: "breakfast/Appam.html"}, {name: "Cucumber salad", url: "breakfast/Cucumber%20salad.html"}, {name: "Soya milk", url: "breakfast/Soya%20milk.html"}, {name: "Dhal fry", url: "breakfast/Dhal%20fry.html"},
		{name: "Dosa", url: "breakfast/dosa.html"},
		{name: "Idli", url: "breakfast/Idli.html"},
		{name: "Pongal", url: "breakfast/Pongal.html"},
		{name: "Chapathi", url: "breakfast/Chapathi.html"},
		{name: "Puttu", url: "breakfast/puttu.html"},
		{name: "Wheat Bread", url: "breakfast/Wheat%20bread.html"},
		{name: "Oats Idli", url: "breakfast/Oats%20idli.html"},
		{name: "Ragi Dosa", url: "breakfast/Ragi%20dosa.html"}},
	"Lunch": {{name: "Chicken Salna", url: "lunch/Chicken%20salna.html"},
		{name: "Curd Rice", url: "lunch/Curd%20rice.html"},
		{name: "Fish fry", url: "lunch/Fish%20fry.html"},
		{name: "Dum aloo", url: "lunch/Dum%20aloo.html"},
		{name: "Paratha", url: "lunch/Paratha.html"},
		{name: "Veg Fried Rice", url: "lunch/Veg%20fried%20rice.html"},
		{name: "Chicken biryani", url: "lunch/Chicken%20biryani.html"},
		{name: "Bean sprout curry", url: "lunch/Bean%20sprout%20curry.html"},
		{name: "Payasam", url: "lunch/Payasam.html"},
		{name: "Tamarind rice", url: "lunch/Tamarind%20rice.html"},
		{name: "Bisibele bath", url: "lunch/Bisibele%20bath.html"},
		{name: "Veg pulao", url: "lunch/Veg%20pulao.html"}},
	"Snacks": {{name: "Soya Milk", url: "snacks/Soya%20Milk.html"}, {name: "Veg Soup", url: "snacks/Veg%20Soup.html"}, {name: "Almonds", url: "snacks/Almonds.html"}, {name: "Apple Juice", url: "snacks/Apple%20Juice.html"}, {name: "Pineapple Juice", url: "snacks/Pineapple%20Juice.html"}, {name: "Buttermilk", url: "snacks/Buttermilk.html"}, {name: "Beetroot Juice", url: "snacks/Beetroot%20Juice.html"}, {name: "Medu Vada", url: "snacks/Medu%20Vada.html"}, {name: "Wheat Bread", url: "snacks/Wheat%20bread.html"}, {name: "Chana Masala", url: "snacks/Chana%20Masala.html"}, {name: "Puttu", url: "snacks/puttu.html"}, {name: "Orange", url: "snacks/Orange%20Juice.html"}, {name: "Rasagulla", url: "snacks/Rasagulla.html"}},
	"Dinner": {{name: "Bisebele bath", url: "dinner/Bisibele%20bath.html"}, {name: "Chicken biryani", url: "dinner/Chicken%20biryani.html"}, {name: "Chapathi", url: "dinner/Chapathi.html"}, {name: "Bean sprout curry", url: "dinner/Bean%20sprout%20curry.html"}, {name: "Dhal fry", url: "dinner/Dhal%20fry.html"}, {name: "Chicken salna", url: "dinner/Chicken%20salna.html"}, {name: "Dosa", url: "dinner/dosa.html"}, {name: "Fish fry", url: "dinner/Fish%20fry.html"}, {name: "Ragi dosa", url: "dinner/Ragi%20dosa.html"}, {name: "Veg fried rice", url: "dinner/Veg%20fried%20rice.html"}, {name: "Idli", url: "dinner/Idli.html"}, {name: "Veg pulao", url: "dinner/Veg%20pulao.html"}, {name: "Matar Paneer Curry", url: "dinner/Matar%20Paneer%20Curry.html"}, {name: "Oats idli", url: "dinner/Oats%20idli.html"}, {name: "Paratha", url: "dinner/Paratha.html"}},
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/evaluatediet", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	// age, _ := strconv.ParseInt(r.Form.Get("age"), 10, 64)
	height, _ := strconv.ParseFloat(r.Form.Get("height"), 64)
	weight, _ := strconv.ParseFloat(r.Form.Get("weight"), 64)
	// gender := r.Form.Get("gender")
	// fastSugarLevel, _ := strconv.ParseFloat(r.Form.Get("fastsugarlevel"), 64)
	// sugarLevel, _ := strconv.ParseFloat(r.Form.Get("sugarlevel"), 64)
	bmi := calculateBMI(height, weight)
	// recommendedFastingSugarLevel, recommendedSugarLevel := calculateRecommendedSugarLevel(age, gender)
	// var recommendedAddedSugar float64
	// if gender == "male" {
	// 	recommendedAddedSugar = getRecommendedSugarForBoy(age)
	// } else {
	// 	recommendedAddedSugar = getRecommendedSugarForGirl(age)
	// }
	breakfastItems := getBreakfastItems()
	lunchItems := getLunchItems()
	snackItems := getSnackItems()
	dinnerItems := getDinnerItems()
	var htmlString = `<!DOCTYPE html>
<html>
  <head>
    <title>Diet Planner</title>
    <link href='https://fonts.googleapis.com/css?family=Open+Sans:400,300,300italic,400italic,600' rel='stylesheet' type='text/css'>
    <link href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700" rel="stylesheet">
    <link href="style/styles.css" rel="stylesheet">
  </head>
  <body>
  <h1>Gluco-M</h1>
    <div class="main-block">
          <div>
            <h3>Name:` + name + `</h3>
            <h3>Age:` + r.Form.Get("age") + `</h3>
            <h3>Height:` + r.Form.Get("height") + `</h3>
            <h3>Weight: ` + r.Form.Get("weight") + `</h3>
			<h3>Gender: ` + r.Form.Get("gender") + `</h3>
            <h3>Fasting Sugar Level: ` + r.Form.Get("fastsugarlevel") + `</h3>
            <h3>Normal Sugar Level(Post Prandial): ` + r.Form.Get("sugarlevel") + `</h3>
			<h3>BMI: ` + strconv.FormatFloat(bmi, 'f', 2, 64) + `</h3>
			<h3>` + bmiResult(bmi) + `</h3>
		  	<h2> Recommended Breakfast Items</h2>
		  	<div>
		  	<h3>` + breakfastItems + `<h3>
		  	</div>
		  	<h2> Recommended Lunch Items</h2>
		  	<div>
		  	<h3>` + lunchItems + `<h3>
		  	</div>
		  	<h2> Recommended Snacks Items</h2>
		  	<div>
		  	<h3>` + snackItems + `<h3>
		  	</div>
		 	 <h2> Recommended Dinner Items</h2>
		  	<div>
		  	<h3>` + dinnerItems + `<h3>
		  </div>
    </div> 
  </body>
</html>`
	w.Write([]byte(htmlString))
}
func getBreakfastItems() string {
	return getResultFoods("Breakfast")

}
func getLunchItems() string {
	return getResultFoods("Lunch")
}
func getSnackItems() string {
	return getResultFoods("Snacks")
}
func getDinnerItems() string {
	return getResultFoods("Dinner")
}
func getResultFoods(mealTime string) string {
	resultList := make([]foodItem, 0)
	for _, breakfastItem := range mealMap[mealTime] {
		resultList = append(resultList, breakfastItem)
	}
	var resultString string
	for _, v := range resultList {
		resultString += `<h3><a href="` + v.url + `">` + v.name + `</a></h3>`
	}
	return resultString
}
func getRecommendedSugarForBoy(age int64) float64 {
	if age <= 3 {
		return 1060 / 8
	} else if age <= 6 {
		return 1350 / 8
	} else if age <= 9 {
		return 1690 / 8
	} else if age <= 12 {
		return 2190 / 8
	} else if age <= 15 {
		return 2750 / 8
	} else if age <= 17 {
		return 3020 / 8
	} else {
		return 2320 / 8
	}
}
func getRecommendedSugarForGirl(age int64) float64 {
	if age <= 3 {
		return 1060 / 8
	} else if age <= 6 {
		return 1350 / 8
	} else if age <= 9 {
		return 1690 / 8
	} else if age <= 12 {
		return 2010 / 8
	} else if age <= 15 {
		return 2330 / 8
	} else if age <= 17 {
		return 2440 / 8
	} else {
		return 2320 / 8
	}
}
func calculateBMI(height, weight float64) float64 {
	return weight / (height * height)

}
func calculateRecommendedSugarLevel(age int64, gender string) (string, string) {
	if age <= 6 {
		return "80-180 mg/dl", "< 180 mg/dl"
	} else if age <= 12 {
		return "80-180 mg/dl", "< 140 mg/dl"
	} else if age <= 19 {
		return "70-150 mg/dl", "< 140 mg/dl"
	} else {
		return "< 100 mg/dl", "< 180 mg/dl"
	}
}
func bmiResult(bmi float64) string {
	if bmi < 19 {
		return "Under Weight"
	} else if bmi < 25 {
		return "Healthy Weight"
	} else if bmi < 30 {
		return "Moderately Obese"
	} else {
		return "Obese"
	}
}
