package api

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/georgedinicola/jeopardy-api/internal/db"

	"github.com/gin-gonic/gin"
)

type JeopardyApiInterface interface {
	GetAllContestants(c *gin.Context)
	GetAllGames(c *gin.Context)
	GetEpisodes(c *gin.Context)
	GetPerformanceForEpisodeNumber(c *gin.Context)
	ExportAllGames(c *gin.Context)
}

type JeopardyApi struct {
	Db db.DatabaseConnx
}

func CreateNewJeopardyApi(db db.DatabaseConnx) *JeopardyApi {
	return &JeopardyApi{Db: db}
}

func (j *JeopardyApi) GetAllContestants(c *gin.Context) {
	allContestants, err := j.Db.GetAllContestants("ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, allContestants)
}

func (j *JeopardyApi) GetAllGames(c *gin.Context) {
	boxScoresAllGames, err := j.Db.GetAllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, boxScoresAllGames)
}

func (j *JeopardyApi) GetAllEpisodes(c *gin.Context) {
	allEpisodes, err := j.Db.GetAllEpisodes("DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, allEpisodes)
}

func (j *JeopardyApi) GetPerformanceForEpisodeNumber(c *gin.Context) {
	episodeNumber := c.Param("episodeNumber")

	episodes, err := j.Db.GetGameByEpisodeNumber(episodeNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, episodes)
}

func (j *JeopardyApi) ExportAllGames(c *gin.Context) {
	fileName := "jeopardy_box_scores.csv"

	boxScoresAllGames, err := j.Db.GetAllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	columnHeaders := []string{
		"Episode Number", "Episode Title", "Episode Date",
		"Contestant Last Name", "Contestant First Name", "Home City", "Home State", "Is Winner",
		"Round One Attempts", "Round One Buzzes", "Round One Buzz Percentage",
		"Round One Correct Answers", "Round One Incorrect Answers", "Round One Correct Answer Percentage",
		"Round One Daily Doubles", "Round One Score",
		"Round Two Attempts", "Round Two Buzzes", "Round Two Buzz Percentage",
		"Round Two Correct Answers", "Round Two Incorrect Answers", "Round Two Correct Answer Percentage",
		"Round Two Daily Double 1", "Round Two Daily Double 2", "Round Two Score",
		"Final Jeopardy Starting Score", "Final Jeopardy Wager", "Final Jeopardy Score",
		"Total Game Attempts", "Total Game Buzzes", "Total Game Buzz Percentage",
		"Total Game Correct Answers", "Total Game Incorrect Answers", "Total Game Correct Answer Percentage",
		"Total Game Daily Doubles Correct", "Total Game Daily Doubles Incorrect", "Total Game Daily Double Winnings",
		"Total Game Score", "Total Triple Stumpers",
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=jeopardy_box_scores.csv")
	c.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(c.Writer)
	writer.Flush()

	writer.Write(columnHeaders)

	for _, score := range boxScoresAllGames {
		writer.Write([]string{
			score.EpisodeNumber,
			score.EpisodeTitle,
			score.EpisodeDate.Format("2006-01-02"),
			score.ContestantLastName,
			score.ContestantFirstName,
			score.HomeCity,
			score.HomeState,
			strconv.FormatBool(score.IsWinner),
			strconv.Itoa(score.RoundOneAttempts),
			strconv.Itoa(score.RoundOneBuzzes),
			strconv.Itoa(score.RoundOneBuzzPercent),
			strconv.Itoa(score.RoundOneCorrectAnswers),
			strconv.Itoa(score.RoundOneIncorrectAnswers),
			strconv.Itoa(score.RoundOneCorrectAnswerPercent),
			strconv.Itoa(score.RoundOneDailyDoubles),
			strconv.Itoa(score.RoundOneScore),
			strconv.Itoa(score.RoundTwoAttempts),
			strconv.Itoa(score.RoundTwoBuzzes),
			strconv.Itoa(score.RoundTwoBuzzPercent),
			strconv.Itoa(score.RoundTwoCorrectAnswers),
			strconv.Itoa(score.RoundTwoIncorrectAnswers),
			strconv.Itoa(score.RoundTwoCorrectAnswerPercent),
			strconv.Itoa(score.RoundTwoDailyDouble1),
			strconv.Itoa(score.RoundTwoDailyDouble2),
			strconv.Itoa(score.RoundTwoScore),
			strconv.Itoa(score.FinalJeopardyStartingScore),
			strconv.Itoa(score.FinalJeopardyWager),
			strconv.Itoa(score.FinalJeopardyScore),
			strconv.Itoa(score.TotalGameAttempts),
			strconv.Itoa(score.TotalGameBuzzes),
			strconv.Itoa(score.TotalGameBuzzPercent),
			strconv.Itoa(score.TotalGameCorrectAnswers),
			strconv.Itoa(score.TotalGameIncorrectAnswers),
			strconv.Itoa(score.TotalGameCorrectAnswerPercent),
			strconv.Itoa(score.TotalGameDailyDoublesCorrect),
			strconv.Itoa(score.TotalGameDailyDoublesIncorrect),
			strconv.Itoa(score.TotalGameDailyDoubleWinnings),
			strconv.Itoa(score.TotalGameScore),
			strconv.Itoa(score.TotalTripleStumpers),
		})
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.File(fileName)
}
