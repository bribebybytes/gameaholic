package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	bitbucketserver "gopkg.in/go-playground/webhooks.v5/bitbucket-server"
	"bribebybytes.in/gameaholic/models"
)

//EventScore holds score for each event that will be added to the player
type EventScore int

const (
	//BitbucketserverRepositoryModifiedPayload event score
	BitbucketserverRepositoryModifiedPayload EventScore = 1
	//BitbucketserverPullRequestMergedPayload event score
	BitbucketserverPullRequestMergedPayload EventScore = 5

	//BitbucketserverPullRequestReviewerApprovedPayload event score
	BitbucketserverPullRequestReviewerApprovedPayload EventScore = 3
)

//CalcScore to calculate score based on events
func CalcScore(externalUser string, event string) {

	return
}

// AddScore
// Identifies player based on webhook
func AddScore(c buffalo.Context, userkey string, provider string, score int) error {

	tx, ok := c.Value("tx").(*pop.Connection)

	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Player
	player := &models.Player{}
	level := &models.Level{}
	achievement := &models.Achievement{}

	// To find the Player the parameter player_id is used.
	var qry string
	qry = ""

	switch provider {
	case "Jira":
		qry = "pemailid = ?"
	case "Bitbucket":
		qry = "pemailid = ?"
	case "Confluence":
		qry = "pemailid = ?"
	}

	//Query Player to add Score
	err := tx.Where(qry, userkey).First(player)

	//Add Score
	player.Pscore = (player.Pscore + score)

	//Query if new score will lead him to new Level
	err = tx.Where("score <= ?", player.Pscore).Order("score desc").First(level)

	if (level.LID != 0) && (player.Plevel != level.Level) {
		player.Plevel = level.Level

		//Query if new level will lead him to new achievement
		err = tx.Where("alevel <= ?", level.Level).Order("alevel desc").First(achievement)

		if (achievement.Alevel != 0) && (achievement.Aname != player.PcurAchive) {
			player.Pavatar = achievement.Aavatar
			player.PcurAchive = achievement.Aname
		}
	}

	// Validate the data from the html form
	err = tx.UpdateColumns(player)

	return err
}

//JiraParser is a simple parser that retrieves JiraEvent and EmailAddress from webhook payload
func JiraParser(JiraPayload string) (JiraEvent string, JiraEmailAddress string) {
	// Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(JiraPayload), &result)

	user := result["user"].(map[string]interface{})

	//Reading each value by its key
	fmt.Println("webhookEvent :", result["webhookEvent"])
	fmt.Println("emailAddress :", user["emailAddress"])

	return result["webhookEvent"].(string), user["emailAddress"].(string)
}

// JiraHandler to handle jira webhooks
func JiraHandler(c buffalo.Context) error {

	fmt.Printf("Actual Request: %+v", c.Request())

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		fmt.Println("Error Parsing Jira payload: %v", err)
		return err
	}

	fmt.Println("Body %s", body)

	JiraEvent, JiraEmailAddress := JiraParser(string(body))

	fmt.Println("JiraEvent: %s", JiraEvent)

	AddScore(c, JiraEmailAddress, "Jira", 5)

	return c.Render(200, r.String("OK"))

}

// Confluence Handler to handle confluence webhooks
func ConfluenceHandler(c buffalo.Context) error {

	fmt.Printf("Actual Request: %+v", c.Request())

	return c.Render(200, r.String("OK"))

}

// BitBucketHandler to handle bitbucket-server webhooks
func BitBucketHandler(c buffalo.Context) error {

	hook, _ := bitbucketserver.New()

	fmt.Printf("Actual Request: %+v", c.Request())

	payload, err := hook.Parse(c.Request(), bitbucketserver.RepositoryModifiedEvent, bitbucketserver.PullRequestOpenedEvent, bitbucketserver.PullRequestMergedEvent, bitbucketserver.PullRequestReviewerApprovedEvent)

	fmt.Printf("Actual Payload: %+v", payload)

	if err != nil {
		if err == bitbucketserver.ErrEventNotFound {
			// ok event wasn;t one of the ones asked to be parsed
		}
	}
	switch payload.(type) {

	case bitbucketserver.RepositoryModifiedPayload:
		release := payload.(bitbucketserver.RepositoryModifiedPayload)
		AddScore(c, release.Actor.EmailAddress, "Bitbucket", 1)
		fmt.Printf("%+v", release)

	case bitbucketserver.PullRequestOpenedPayload:
		pullRequest := payload.(bitbucketserver.PullRequestOpenedPayload)
		AddScore(c, pullRequest.Actor.EmailAddress, "Bitbucket", 3)
		fmt.Printf("%+v", pullRequest)

	case bitbucketserver.PullRequestMergedPayload:
		pullRequest := payload.(bitbucketserver.PullRequestMergedPayload)
		AddScore(c, pullRequest.Actor.EmailAddress, "Bitbucket", int(BitbucketserverPullRequestMergedPayload))
		fmt.Printf("%+v", pullRequest)

	case bitbucketserver.PullRequestReviewerApprovedPayload:
		pullRequest := payload.(bitbucketserver.PullRequestReviewerApprovedPayload)
		AddScore(c, pullRequest.Actor.EmailAddress, "Bitbucket", int(BitbucketserverPullRequestReviewerApprovedPayload))
		fmt.Printf("%+v", pullRequest)
	}
	return c.Render(200, r.String("OK"))

}
