/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package pull

import (
	"context"
	"fmt"
	"github.com/akimdev15/leetcode/internal/database"
	"github.com/akimdev15/leetcode/query"
	"github.com/akimdev15/leetcode/sql"
	"github.com/akimdev15/leetcode/utils"
	"github.com/spf13/cobra"
	"time"
)

var count int

// PullCmd represents the pull command
var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		pullRecentlySolvedQuestions(count)
	},
}

func init() {
	PullCmd.Flags().IntVarP(&count, "count", "c", 0, "how many recent submissions to pull")
}

// pullRecentlySolvedQuestions - pulls all the recently solved questions and adds them to the list
func pullRecentlySolvedQuestions(limit int) {
	if limit == 0 {
		limit = 10
		fmt.Printf("Using default value to pull recent submissions: %d\n", limit)
	}
	submissions, err := query.GetRecentSubmissions(limit)
	if err != nil {
		fmt.Println("get recently solved submissions err:", err)
		return
	}

	dbStruct, err := sql.InitDB()
	if err != nil {
		fmt.Println("Error initializing database. Error: ", err)
		return
	}
	defer dbStruct.DB.Close()

	ctx := context.Background()

	var savedQuestions int
	var skippedQuestions int
	for _, submission := range submissions {
		questionDetail, err := query.GetQuestionDetails(submission.TitleSlug)
		if err != nil {
			fmt.Printf("get question details err %v for the question detail: %v\n", err, questionDetail)
			continue
		}

		// Create new question
		_, err = dbStruct.Queries.CreateQuestion(ctx, database.CreateQuestionParams{
			ID:         questionDetail.QuestionID,
			Name:       submission.Title,
			Url:        utils.ConstructLeetcodeURL(submission.TitleSlug),
			Solved:     "1",
			Difficulty: questionDetail.Difficulty,
			UpdatedAt:  time.Now().UTC(),
			Confidence: "LOW",
		})
		if err != nil {
			skippedQuestions++
		} else {
			savedQuestions++
		}
	}

	fmt.Printf("Saved questions: %d\nSkipped Questions: %d\n", savedQuestions, skippedQuestions)
}
