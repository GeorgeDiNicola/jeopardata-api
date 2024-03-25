package db

import (
	"fmt"
	"georgedinicola/jeopardy-api/internal/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	CreateGormDbConnection() (*gorm.DB, error)
	GetAllEpisodes() ([]model.Episode, error)
	GetMostRecentEpisodeNumber() (string, error)
}

type DatabaseConnx struct {
	gorm *gorm.DB
}

func CreateNewGormDbConnection() (DatabaseConnx, error) {
	dbHost, dbUsername, dbPassword := os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
	dbName, dbPort, dbTimezone := os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_TIMEZONE")

	gormDB, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=%s", dbHost, dbPort, dbUsername, dbPassword, dbName, dbTimezone)), &gorm.Config{})
	if err != nil {
		return DatabaseConnx{}, err
	}

	return DatabaseConnx{gormDB}, nil
}

func (d *DatabaseConnx) GetMostRecentEpisodeNumber() (string, error) {
	var mostRecentBoxScore model.JeopardyGameBoxScore

	result := d.gorm.Order("episode_number DESC").First(&mostRecentBoxScore)
	if result.Error != nil {
		return "", result.Error
	}

	return mostRecentBoxScore.EpisodeNumber, nil
}

func (d *DatabaseConnx) GetAllEpisodes() ([]model.Episode, error) {
	var episodes []model.Episode
	if result := d.gorm.Model(&model.JeopardyGameBoxScore{}).Select("EpisodeNumber", "EpisodeDate").Find(&episodes).Order("episode_date, episode_number DESC"); result.Error != nil {
		return nil, result.Error
	}

	return episodes, nil
}
