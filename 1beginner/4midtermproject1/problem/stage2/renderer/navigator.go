package renderer

import (
	"errors"
	"fmt"
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage2/storage"
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage2/tasks"
)

type Navigator interface {
	show() error
	navigate() error
}

var n Navigator

func init() {
	db, err := storage.NewTasksDB()
	if err != nil {
		fmt.Println("Error initializing initialize db")
	}
	n = &NavigatorImpl{
		DB: db,
		CurrentPage: newEpicList(db.SelectEpics()),
	}
}

type NavigatorImpl struct {
	DB storage.JsonDB
	CurrentPage Page
}

func (n *NavigatorImpl) show() error {
	n.CurrentPage.show()
	return nil
}

func (n *NavigatorImpl) navigate() error {
	fmt.Println("Please select an option:")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error navigating", err)
		return err
	}
	switch n.CurrentPage.(type) {
	case EpicList:
		switch input {
		case "C":
			return n.createEpic()
		case "D":
			return n.viewEpicDetails()
		default:
			fmt.Println("Please select from available options.")
		}
	case EpicDetails:
		switch input {
		case "P":
			return n.epicDetailsPreviousPage()
		case "U":
			return n.updateEpic()
		case "R":
			return n.removeEpic()
		case "C":
			return n.createEpicStory()
		case "D":
			return n.storyDetails()
		default:
			fmt.Println("Please select from available options.")
		}
	case StoryDetails:
		switch input {
		case "P":
			return n.storyDetailsPreviousPage()
		case "U":
			return n.updateEpicStory()
		case "R":
			return n.removeEpicStory()
		default:
			fmt.Println("Please select from available options.")
		}
	}
	err = errors.New("error asserting CurrentPage")
	fmt.Println("Error asserting CurrentPage", err)
	return err
}

func (n *NavigatorImpl) createEpic() error {
	// Todo
	var epics []tasks.Epic
	n.CurrentPage = newEpicList(epics)
	return nil
}

func (n *NavigatorImpl) viewEpicDetails() error {
	// Todo
	var epic tasks.Epic
	var epicStories []tasks.Story
	n.CurrentPage = newEpicDetails(epic, epicStories)
	return nil
}

func (n *NavigatorImpl) epicDetailsPreviousPage() error {
	// Todo
	var epics []tasks.Epic
	n.CurrentPage = newEpicList(epics)
	return nil
}

func (n *NavigatorImpl) updateEpic() error {
	// Todo
	var updatedEpic tasks.Epic
	var updatedEpicStories []tasks.Story
	n.CurrentPage = newEpicDetails(updatedEpic, updatedEpicStories)
	return nil
}

func (n *NavigatorImpl) removeEpic() error {
	// Todo
	var epics []tasks.Epic
	n.CurrentPage = newEpicList(epics)
	return nil
}

func (n *NavigatorImpl) createEpicStory() error {
	// Todo
	var updatedEpic tasks.Epic
	var updatedEpicStories []tasks.Story
	n.CurrentPage = newEpicDetails(updatedEpic, updatedEpicStories)
	return nil
}

func (n *NavigatorImpl) storyDetails() error {
	// Todo
	var epic tasks.Epic
	var epicStory tasks.Story
	n.CurrentPage = newStoryDetails(epic, epicStory)
	return nil
}

func (n *NavigatorImpl) storyDetailsPreviousPage() error {
	// Todo
	var epic tasks.Epic
	var epicStories []tasks.Story
	n.CurrentPage = newEpicDetails(epic, epicStories)
	return nil
}

func (n *NavigatorImpl) updateEpicStory() error {
	// Todo
	var updatedEpic tasks.Epic
	var updatedEpicStory tasks.Story
	n.CurrentPage = newStoryDetails(updatedEpic, updatedEpicStory)
	return nil
}

func (n *NavigatorImpl) removeEpicStory() error {
	// Todo
	var updatedEpic tasks.Epic
	var updatedEpicStories []tasks.Story
	n.CurrentPage = newEpicDetails(updatedEpic, updatedEpicStories)
	return nil
}