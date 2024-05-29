package database

func CreateAnswer(answer *Answer) error {
	result := db.Create(answer)
	return result.Error
}

func GetAnswers() ([]Answer, error) {
	var answers []Answer
	result := db.Find(&answers)
	return answers, result.Error
}
