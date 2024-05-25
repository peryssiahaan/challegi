package database

import "challegi/models"

func CreateQuestion(question *models.Question) error {
	result := db.Create(question)
	return result.Error
}

func GetQuestions() ([]models.Question, error) {
	var questions []models.Question
	result := db.Find(&questions)
	return questions, result.Error
}
