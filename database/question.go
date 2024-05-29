package database

func CreateQuestion(question *Question) error {
	result := db.Create(question)
	return result.Error
}

func GetQuestions() ([]Question, error) {
	var questions []Question
	result := db.Find(&questions)
	return questions, result.Error
}
