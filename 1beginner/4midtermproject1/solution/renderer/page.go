package renderer

import (
	"fmt"
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/solution/tasks"
)

type Page interface {
	show()
}

type Buttons struct {
	DisplayText []string
}

type EpicList struct {
	Es []tasks.Epic
	Bs Buttons
}

func (s EpicList) show() {
	fmt.Printf("\nEpic List:\n")
	for _, epic := range s.Es {
		fmt.Printf("[%v][%v] %v\n", epic.ID, epic.Status, epic.Title)
	}
	fmt.Printf("\n\nOptions:\n")

	for _, button := range s.Bs.DisplayText {
		fmt.Printf("%v    ", button)
	}
	fmt.Printf("\n\n")
}

func newEpicList(es []tasks.Epic) EpicList {
	el := EpicList{
		Es: es,
		Bs: epicListButtons,
	}
	return el
}

var epicListButtons = Buttons {
	DisplayText: []string {
		"[C] Create an Epic",
		"[D] Details on an Epic",
	},
}

type EpicDetails struct {
	E tasks.Epic
	Ss []tasks.Story
	Bs Buttons
}

func (e EpicDetails) show() {
	fmt.Printf("\nEpic Details:\n")
	fmt.Println("ID: ", e.E.ID)
	fmt.Println("Title: ", e.E.Title)
	fmt.Println("Description: ", e.E.Description)
	fmt.Println("Status: ", e.E.Status)

	fmt.Printf("\nStories:\n")
	for _, story := range e.Ss {
		fmt.Printf("[%v][%v] %v\n", story.ID, story.Status, story.Title)
	}

	fmt.Printf("\nOptions:\n")

	for _, button := range e.Bs.DisplayText {
		fmt.Printf("%v    ", button)
	}
	fmt.Printf("\n\n")
}

func newEpicDetails(e tasks.Epic, ss []tasks.Story) EpicDetails {
	ed := EpicDetails{
		E: e,
		Ss: ss,
		Bs: epicDetailsButtons,
	}
	return ed
}

var epicDetailsButtons = Buttons {
	DisplayText: []string {
		"[P] Previous Page",
		"[U] Update Epic",
		"[R] Remove Epic",
		"[C] Create a Story",
		"[D] Details on a Story",
	},
}

type StoryDetails struct {
	E tasks.Epic
	S tasks.Story
	Bs Buttons
}

func (s StoryDetails) show() {
	fmt.Printf("\nStory Details:\n")
	fmt.Println("ID: ", s.S.ID)
	fmt.Println("Title: ", s.S.Title)
	fmt.Println("Description: ", s.S.Description)
	fmt.Println("Status: ", s.S.Status)

	fmt.Printf("\nOptions:\n")

	for _, button := range s.Bs.DisplayText {
		fmt.Printf("%v    ", button)
	}
	fmt.Printf("\n\n")
}

func newStoryDetails(e tasks.Epic, s tasks.Story) StoryDetails {
	sd := StoryDetails{
		E: e,
		S: s,
		Bs: storyDetailsButtons,
	}
	return sd
}

var storyDetailsButtons = Buttons {
	DisplayText: []string {
		"[P] Previous Page",
		"[U] Update Story",
		"[R] Remove Story",
	},
}
