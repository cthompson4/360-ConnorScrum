package main

//Created by Connor Thompson
import "fmt"
import "io/ioutil"
import "log"
import "net/http"
import "encoding/json"
import "strconv"

type Questions struct {
	Questions []Question `json:"question"`
}
type Question struct {
	Question			string	`json:"question"`
	Answer1				string	`json:"answer1"`
	Answer2				string	`json:"answer2"`
	Answer3				string	`json:"answer3"`
	Answer4				string	`json:"answer4"`
	CorrectAnswer		string  `json:"correctanswer"`
}
//Function pulls from online. Tutorial from: https://blog.logrocket.com/making-http-requests-in-go/
func main() {
	var questionBank string
	fmt.Println("Welcome to QuizMaster!")
	fmt.Println("There are two question banks available, would you like 1) Math or 2) Geography?")
	if _, err := fmt.Scan(&questionBank); err != nil {
		fmt.Println("Please enter 1 or 2!")
	}
	url := "https://victorious-grass-035ec5810.1.azurestaticapps.net/questionbank" + questionBank + ".json"
	resp, err := http.Get(url)
	if err != nil {
   		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var questions Questions
	json.Unmarshal(body, &questions)
	for i := 0; i < len(questions.Questions); i++ {
		fmt.Println("Question: " + questions.Questions[i].Question)
		fmt.Println("1) " + questions.Questions[i].Answer1)
		fmt.Println("2) " + questions.Questions[i].Answer2)
		fmt.Println("3) " + questions.Questions[i].Answer3)
		fmt.Println("4) " + questions.Questions[i].Answer4)
		correctanswer, err := strconv.Atoi(questions.Questions[i].CorrectAnswer)
		if err != nil {
			log.Fatalln(err)
		}
		var answer int
		if _, err := fmt.Scan(&answer); err != nil {
			fmt.Println("Please enter 1, 2, 3, or 4!")
		}
		if answer != correctanswer {
			fmt.Println("Incorrect!")
			fmt.Println()
		} else {
			fmt.Println("Correct!")
			fmt.Println()
		}
	}
}
