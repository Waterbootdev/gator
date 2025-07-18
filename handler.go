package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Waterbootdev/gator/internal/database"
	"github.com/Waterbootdev/gator/internal/feeds"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("please provide a username")
	}

	user, err := s.DB.GetUser(context.Background(), cmd.arguments[0])

	if err != nil {
		fmt.Println("no user found")
		os.Exit(1)
	}

	if user.Name != cmd.arguments[0] {
		fmt.Println("Fatal error occured while logging in")
		os.Exit(1)
	}

	err = s.Config.SetUser(cmd.arguments[0])

	if err != nil {
		return err
	}

	fmt.Println("Login successful")

	return nil
}

func handleRegister(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return errors.New("please provide a username")
	}

	currentTime := time.Now()

	currentUser, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.arguments[0],
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return s.Config.SetUser(currentUser.Name)
}

func handleReset(s *state, _ command) error {
	return s.DB.DeleteALLUsers(context.Background())
}

func handleGetUsers(s *state, _ command) error {

	users, err := s.DB.GetUsers(context.Background())

	if err != nil {
		return err
	}

	curentUserName := s.Config.CurrentUserName

	for _, user := range users {
		if user.Name == curentUserName {
			fmt.Println(user.Name, "(current)")
		} else {
			fmt.Println(user.Name)
		}
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"

	feed, err := feeds.FetchFeed(context.Background(), url)

	if err != nil {
		return err
	}
	fmt.Println(feed.Channel.Title)

	fmt.Println(feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.Description)
		fmt.Println(item.PubDate)
	}

	return nil
}

func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.arguments) < 2 {
		return errors.New("please provide a feedname and url")
	}

	user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)

	if err != nil {
		return err
	}

	currentTime := time.Now()

	feedId := uuid.New()

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        feedId,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
	})

	if err != nil {
		return err
	}

	_, err = s.DB.CreateFollower(context.Background(), database.CreateFollowerParams{
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	})

	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.DB.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	userNames, err := s.getUserNames()

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name, feed.Url, userNames[feed.UserID])
	}

	return nil
}

func handlerFollow(s *state, cmd command) error {

	if len(cmd.arguments) < 1 {
		return errors.New("please provide url")
	}

	user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)

	if err != nil {
		return err
	}

	feed, err := s.DB.GetFeedByURL(context.Background(), cmd.arguments[0])

	if err != nil {
		return err
	}

	currentTime := time.Now()

	feedFollowRow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	})

	if err != nil {
		return err
	}

	fmt.Println(feedFollowRow)

	return nil
}
func handlerFollowing(s *state, cmd command) error {

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), s.Config.CurrentUserName)

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}

	return nil
}
func (c *commands) registerHandlers() {

	c.availableCommands = map[string]func(s *state, cmd command) error{}
	c.register("login", handlerLogin)
	c.register("register", handleRegister)
	c.register("reset", handleReset)
	c.register("users", handleGetUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", handlerAddFeed)
	c.register("feeds", handlerFeeds)
	c.register("follow", handlerFollow)
	c.register("following", handlerFollowing)
}
