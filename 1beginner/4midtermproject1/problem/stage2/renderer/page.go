package renderer

import (
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage2/tasks"
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
	// Todo
	/*
	EpicList.show() should print out:

	Epic List:
	[1][IN_PROGRESS] Epic - Project 1 for the Bootcamp
	[4][OPEN] Epic - Project 2 for the Bootcamp
	[5][OPEN] Epic - Project 3 for the Bootcamp

	Options:
	[C] Create an Epic    [D] Details on an Epic
	*/
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
	// Todo
	/*
	EpicDetails.show() should print out:

	Epic Details:
	ID:  1
	Title:  Epic - Project 1
	Description:  This is Project 1 for the Bootcamp
	Status:  IN_PROGRESS

	Stories:
	[2][OPEN] Story - Project 1 Solution
	[3][OPEN] Story - Project 1 README

	Options:
	[P] Previous Page    [U] Update Epic    [R] Remove Epic    [C] Create a Story    [D] Details on a Story
	*/
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
	// Todo
	/*
	StoryDetails.show() should print out:

	Story Details:
	ID:  9
	Title:  New Story title
	Description:  New Story description
	Status:  OPEN

	Options:
	[P] Previous Page    [U] Update Story    [R] Remove Story
	*/
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
