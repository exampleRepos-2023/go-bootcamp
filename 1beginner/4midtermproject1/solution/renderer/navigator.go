package renderer

import (
	"errors"
	"fmt"
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/solution/storage"
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/solution/tasks"
	"strconv"
)

type Navigator interface {
	show() error
	navigate() error
}

var n Navigator

func init() {
	n = &NavigatorImpl{
		DB: storage.DB,
		CurrentPage: newEpicList(storage.DB.SelectEpics()),
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
	newEpic := tasks.Epic{
		ID: tasks.NewTaskID,
		Status: tasks.Open,
	}
	fmt.Println("\nPlease enter Epic title:")
	var input string
	_, err := fmt.Scanf("%q\n", &input)
	if err != nil {
		fmt.Println("Invalid Epic title.")
		return nil
	}
	newEpic.Title = input
	fmt.Println("\nPlease enter Epic description:")
	_, err = fmt.Scanf("%q\n", &input)
	if err != nil {
		fmt.Println("Invalid Epic description.")
		return nil
	}
	newEpic.Description = input
	_, err = n.DB.UpsertEpic(newEpic)
	if err != nil {
		fmt.Println("Error upserting initialize Epic")
		return err
	}
	n.CurrentPage = newEpicList(n.DB.SelectEpics())
	return nil
}

func (n *NavigatorImpl) viewEpicDetails() error {
	fmt.Println("\nPlease select Epic ID:")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid Epic ID.")
		return nil
	}

	epicId, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Failed to convert Epic ID.")
		return nil
	}

	epic, exists := n.DB.SelectEpic(epicId)
	if !exists {
		fmt.Println("No matching Epic ID.")
		return nil
	}

	stories := n.DB.SelectStories(epicId)

	n.CurrentPage = newEpicDetails(epic, stories)
	return nil
}

func (n *NavigatorImpl) epicDetailsPreviousPage() error {
	n.CurrentPage = newEpicList(n.DB.SelectEpics())
	return nil
}

func (n *NavigatorImpl) updateEpic() error {
	epicDetails, ok := n.CurrentPage.(EpicDetails)
	if !ok {
		return errors.New("error asserting EpicDetails")
	}

	fmt.Println("\nEpic fields to update:")
	fmt.Println("[T] Title    [D] Description    [S] Status")
	fmt.Println("\nPlease select an option:")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid epic field.")
		return nil
	}
	newEpic := tasks.Epic{
		ID:          epicDetails.E.ID,
		Title:       epicDetails.E.Title,
		Description: epicDetails.E.Description,
		Status:      epicDetails.E.Status,
		StoryIDs:    epicDetails.E.StoryIDs,
	}
	switch input {
	case "T":
		fmt.Println("\nCurrent Epic title: ")
		fmt.Printf("\"%s\"", epicDetails.E.Title)
		fmt.Println("\nNew Epic title: ")
		var input string
		_, err := fmt.Scanf("%q", &input)
		if err != nil {
			fmt.Println("Invalid initialize title.")
			return nil
		}

		newEpic.Title = input
	case "D":
		fmt.Println("\nCurrent Epic description: ")
		fmt.Printf("\"%s\"", epicDetails.E.Description)
		fmt.Println("\nNew Epic description: ")
		var input string
		_, err := fmt.Scanf("%q", &input)
		if err != nil {
			fmt.Println("Invalid initialize description.")
			return nil
		}

		newEpic.Description = input
	case "S":
		fmt.Println("\nCurrent Epic status: ")
		fmt.Println(epicDetails.E.Status)
		fmt.Println("New Epic status: [O] OPEN    [I] IN_PROGRESS    [R] RESOLVED    [C] CLOSED")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid initialize description.")
			return nil
		}
		switch input {
		case "O":
			newEpic.Status = tasks.Open
		case "I":
			newEpic.Status = tasks.InProgress
		case "R":
			newEpic.Status = tasks.Resolved
		case "C":
			newEpic.Status = tasks.Closed
		default:
			fmt.Println("Please select from available options.")
		}
	}
	_, err = n.DB.UpsertEpic(newEpic)
	if err != nil {
		fmt.Println("Error upserting initialize Epic.")
		return err
	}
	n.CurrentPage = newEpicDetails(newEpic, epicDetails.Ss)
	return nil
}

func (n *NavigatorImpl) removeEpic() error {
	epicDetails, ok := n.CurrentPage.(EpicDetails)
	if !ok {
		return errors.New("error asserting EpicDetails")
	}

	fmt.Printf("\nConfirm to remove Epic: [%v] %v", epicDetails.E.ID, epicDetails.E.Title)
	fmt.Println("\n    (Note - All Stories under this Epic will also be removed)")
	fmt.Println("[Y] Yes    [n] No")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return err
	}

	if input != "Y" {
		fmt.Println("\nCanceled to remove Epic.")
		return nil
	}

	for _, storyId := range epicDetails.E.StoryIDs {
		err := n.DB.DeleteStory(storyId)
		if err != nil {
			fmt.Println("Error deleting Story.")
			return err
		}
	}
	err = n.DB.DeleteEpic(epicDetails.E.ID)
	if err != nil {
		fmt.Println("Error deleting Epic.")
		return err
	}
	n.CurrentPage = newEpicList(n.DB.SelectEpics())
	return nil
}

func (n *NavigatorImpl) createEpicStory() error {
	epicDetails, ok := n.CurrentPage.(EpicDetails)
	if !ok {
		return errors.New("error asserting EpicDetails")
	}

	newStory := tasks.Story{
		ID: tasks.NewTaskID,
		Status: tasks.Open,
	}
	fmt.Println("\nPlease enter Story title:")
	var input string
	_, err := fmt.Scanf("%q\n", &input)
	if err != nil {
		fmt.Println("Invalid Story title.")
		return nil
	}
	newStory.Title = input
	fmt.Println("\nPlease enter Story description:")
	_, err = fmt.Scanf("%q\n", &input)
	if err != nil {
		fmt.Println("Invalid Story description.")
		return nil
	}
	newStory.Description = input
	insertedStory, err := n.DB.UpsertStory(newStory)
	if err != nil {
		fmt.Println("Error upserting initialize Story")
		return err
	}

	updatedEpicStoryIDs := []int{insertedStory.ID}
	for _, storyID := range epicDetails.E.StoryIDs {
		updatedEpicStoryIDs = append(updatedEpicStoryIDs, storyID)
	}
	updatedEpic := tasks.Epic{
		ID:          epicDetails.E.ID,
		Title:       epicDetails.E.Title,
		Description: epicDetails.E.Description,
		Status:      epicDetails.E.Status,
		StoryIDs:    updatedEpicStoryIDs,
	}
	_, err = n.DB.UpsertEpic(updatedEpic)
	if err != nil {
		fmt.Println("Error upserting initialize Epic")
		return err
	}
	n.CurrentPage = newEpicDetails(updatedEpic, n.DB.SelectStories(updatedEpic.ID))
	return nil
}

func (n *NavigatorImpl) storyDetails() error {
	epicDetails, ok := n.CurrentPage.(EpicDetails)
	if !ok {
		return errors.New("error asserting EpicDetails")
	}

	fmt.Println("\nPlease select Story ID:")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid Story ID.")
		return nil
	}
	story, exists := n.DB.SelectStory(input)
	if !exists {
		fmt.Println("No matching Story ID.")
		return nil
	}
	n.CurrentPage = newStoryDetails(epicDetails.E, story)
	return nil
}

func (n *NavigatorImpl) storyDetailsPreviousPage() error {
	storyDetails, ok := n.CurrentPage.(StoryDetails)
	if !ok {
		return errors.New("error asserting EpicDetails")
	}

	n.CurrentPage = newEpicDetails(storyDetails.E, n.DB.SelectStories(storyDetails.E.ID))
	return nil
}

func (n *NavigatorImpl) updateEpicStory() error {
	storyDetails, ok := n.CurrentPage.(StoryDetails)
	if !ok {
		return errors.New("error asserting EpicDetails")
	}

	fmt.Println("\nStory fields to update:")
	fmt.Println("[T] Title    [D] Description    [S] Status")
	fmt.Println("\nPlease select an option:")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid story field.")
		return nil
	}
	newStory := tasks.Story{
		ID: storyDetails.S.ID,
		Title: storyDetails.S.Title,
		Description: storyDetails.S.Description,
		Status: storyDetails.S.Status,
	}
	switch input {
	case "T":
		fmt.Println("\nCurrent Story title: ")
		fmt.Printf("\"%s\"", storyDetails.S.Title)
		fmt.Println("\nNew title: ")
		var input string
		_, err := fmt.Scanf("%q", &input)
		if err != nil {
			fmt.Println("Invalid initialize title.")
			return nil
		}

		newStory.Title = input
	case "D":
		fmt.Println("\nCurrent Story description: ")
		fmt.Printf("\"%s\"", storyDetails.S.Description)
		fmt.Println("\nNew Story description: ")
		var input string
		_, err := fmt.Scanf("%q", &input)
		if err != nil {
			fmt.Println("Invalid initialize description.")
			return nil
		}

		newStory.Description = input
	case "S":
		fmt.Println("Current status: ")
		fmt.Println(storyDetails.S.Status)
		fmt.Println("New status: [O] OPEN    [I] IN_PROGRESS    [R] RESOLVED    [C] CLOSED")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid initialize description.")
			return nil
		}
		switch input {
		case "O":
			newStory.Status = tasks.Open
		case "I":
			newStory.Status = tasks.InProgress
		case "R":
			newStory.Status = tasks.Resolved
		case "C":
			newStory.Status = tasks.Closed
		default:
			fmt.Println("Please select from available options.")
		}
	}
	_, err = n.DB.UpsertStory(newStory)
	if err != nil {
		fmt.Println("Error upserting initialize story")
		return err
	}
	n.CurrentPage = newStoryDetails(storyDetails.E, newStory)
	return nil
}

func (n *NavigatorImpl) removeEpicStory() error {
	storyDetails, ok := n.CurrentPage.(StoryDetails)
	if !ok {
		return errors.New("error asserting EpicDetails")
	}

	fmt.Printf("\nConfirm to remove Story: [%v] %v", storyDetails.S.ID, storyDetails.S.Title)
	fmt.Println("\n[Y] Yes    [n] No")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return err
	}

	if input != "Y" {
		fmt.Println("\nCanceled to remove Story.")
		return nil
	}

	err = n.DB.DeleteStory(storyDetails.S.ID)
	if err != nil {
		fmt.Println("Error deleting story")
		return err
	}

	var newEpicStoryIDs []int
	for _, epicStoryID := range storyDetails.E.StoryIDs {
		if epicStoryID != storyDetails.S.ID {
			newEpicStoryIDs = append(newEpicStoryIDs, epicStoryID)
		}
	}
	newEpic := tasks.Epic{
		ID:          storyDetails.E.ID,
		Title:       storyDetails.E.Title,
		Description: storyDetails.E.Description,
		Status:      storyDetails.E.Status,
		StoryIDs:    newEpicStoryIDs,
	}

	_, err = n.DB.UpsertEpic(newEpic)
	if err != nil {
		fmt.Println("Error upserting initialize Epic")
		return err
	}

	n.CurrentPage = newEpicDetails(newEpic, n.DB.SelectStories(storyDetails.E.ID))
	return nil
}