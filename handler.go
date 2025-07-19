package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Waterbootdev/gator/internal/database"
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

func handleGetUsers(s *state, _ command, user database.User) error {

	users, err := s.DB.GetUsers(context.Background())

	if err != nil {
		return err
	}

	curentUserName := user.Name

	for _, user := range users {
		if user.Name == curentUserName {
			fmt.Println(user.Name, "(current)")
		} else {
			fmt.Println(user.Name)
		}
	}

	return nil
}

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return errors.New("please provide a duration")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])

	if err != nil || timeBetweenRequests <= 0 {
		return errors.New("please provide a correct duration")
	}

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		err := s.DB.ScrapeFeeds(user)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func handlerAddFeed(s *state, cmd command, user database.User) error {

	if len(cmd.arguments) < 2 {
		return errors.New("please provide a feedname and url")
	}

	user, err := s.DB.GetUser(context.Background(), user.Name)

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
		fmt.Println(feed.Name, feed.Url, userNames[feed.UserID], feed.LastFetchAt)
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.arguments) < 1 {
		return errors.New("please provide url")
	}

	user, err := s.DB.GetUser(context.Background(), user.Name)

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
func handlerFollowing(s *state, cmd command, user database.User) error {

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.Name)

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}

	return nil
}
func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.arguments) < 1 {
		return errors.New("please provide url")
	}

	feed, err := s.DB.GetFeedByURL(context.Background(), cmd.arguments[0])

	if err != nil {
		return err
	}

	err = s.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}

	return nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {

	posts, err := s.DB.GetPostsByUser(context.Background(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  postLimit(cmd),
	})

	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title, post.PublishedAt)
	}

	return nil
}

func (c *commands) registerHandlers() {

	c.availableCommands = map[string]func(s *state, cmd command) error{}
	c.register("login", handlerLogin)
	c.register("register", handleRegister)
	c.register("reset", handleReset)
	c.register("users", middlewareLoggedIn(handleGetUsers))
	c.register("agg", middlewareLoggedIn(handlerAgg))
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", middlewareLoggedIn(handlerFollowing))
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	c.register("browse", middlewareLoggedIn(handlerBrowse))
}
