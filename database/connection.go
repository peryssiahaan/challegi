package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Content    string    `json:"content"`
	Difficulty uint8     `json:"difficulty" gorm:"check:difficulty >= 1 AND difficulty <= 5"`
	GroupID    uint      `json:"group_id"`                    // Foreign key to Group
	Answers    []*Answer `gorm:"many2many:question_answers;"` // Many-to-many relationship with answers
}

type Answer struct {
	gorm.Model
	Content   string      `json:"content"`
	Type      string      `json:"type"`                        // Type field for categorizing answers
	GroupID   uint        `json:"group_id"`                    // Foreign key to Group
	Questions []*Question `gorm:"many2many:question_answers;"` // Many-to-many relationship with questions
}

type Group struct {
	gorm.Model
	Name      string     `json:"name" gorm:"unique"` // Ensure group names are unique
	Questions []Question `json:"questions" gorm:"foreignKey:GroupID"`
	Answers   []Answer   `json:"answers" gorm:"foreignKey:GroupID"`
}

// Join table for many-to-many relationship between Question and Answer
type QuestionAnswer struct {
	QuestionID uint `gorm:"primaryKey"`
	AnswerID   uint `gorm:"primaryKey"`
}

var db *gorm.DB

func Init() error {
	var err error
	connectionString := os.Getenv("POSTGRES_CS")
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Group{}, &Question{}, &Answer{}, &QuestionAnswer{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return err
	}

	return nil
}
