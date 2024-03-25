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
	GetGameByEpisodeNumber(episodeNumber string) ([]model.JeopardyGameBoxScore, error)
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
	if err := d.gorm.Order("episode_number DESC").First(&mostRecentBoxScore); err != nil {
		return "", fmt.Errorf("failed to retrieve most recent episode number: %v", err)
	}

	return mostRecentBoxScore.EpisodeNumber, nil
}

func (d *DatabaseConnx) GetAllEpisodes(orderBy string) ([]model.Episode, error) {
	var episodes []model.Episode

	sql := fmt.Sprintf(`SELECT DISTINCT ON (episode_number) episode_number, episode_date FROM jeopardy_game_box_scores ORDER BY episode_number, episode_date %s`,
		orderBy)
	if result := d.gorm.Raw(sql).Scan(&episodes); result.Error != nil {
		return nil, result.Error
	}

	return episodes, nil
}

func (d *DatabaseConnx) GetAllContestants(orderBy string) ([]model.Contestant, error) {
	var constestants []model.Contestant

	sql := fmt.Sprintf(`SELECT DISTINCT ON (contestant_last_name, contestant_first_name, home_city, home_state) contestant_first_name, contestant_last_name, home_city, home_state, is_winner FROM jeopardy_game_box_scores ORDER BY contestant_last_name %s`,
		orderBy)
	if result := d.gorm.Raw(sql).Scan(&constestants); result.Error != nil {
		return nil, result.Error
	}

	return constestants, nil
}

func (d *DatabaseConnx) GetGameByEpisodeNumber(episodeNumber string) ([]model.JeopardyGameBoxScore, error) {
	var boxScores []model.JeopardyGameBoxScore
	if err := d.gorm.Where("episode_number = ?", episodeNumber).Find(&boxScores).Error; err != nil {
		return nil, err
	}

	return boxScores, nil
}
