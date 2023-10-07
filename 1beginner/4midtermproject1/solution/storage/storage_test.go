package storage

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/solution/tasks"
	"io/ioutil"
	"testing"
)

var testStoryID = 1
var testStoryTitle = "testStoryTitle"
var testStoryDescription = "testStoryDescription"
var testStoryStatus = tasks.Open
var testStory = tasks.Story{
	ID: testStoryID,
	Title: testStoryTitle,
	Description: testStoryDescription,
	Status: testStoryStatus,
}

var testEpicID = testStoryID+1
var testEpicTitle = "testEpicTitle"
var testEpicDescription = "testEpicDescription"
var testEpicStatus = tasks.Open
var testEpicStoryIDs = []int{testStoryID}
var testEpic = tasks.Epic{
	ID:          testEpicID,
	Title:       testEpicTitle,
	Description: testEpicDescription,
	Status:      testEpicStatus,
	StoryIDs:    testEpicStoryIDs,
}

func TestInit(t *testing.T) {
	// Set mode = Test
	testTasks, err := buildTestTasks(t)

	assert.Nil(t, err)
	lastTaskID := DB.SelectLastTaskID()
	assert.NotEqual(t, lastTaskID, tasks.NewTaskID)
	epics := DB.SelectEpics()
	epicStories := DB.SelectStories(epics[0].ID)
	assert.Equal(t, getValuesSortedByKeys(testTasks.Epics), epics)
	assert.NotEmpty(t, epics)
	assert.Equal(t, getValuesSortedByKeys(testTasks.Stories), epicStories)
	assert.NotEmpty(t, epicStories)
}

func TestJsonDB_SelectEpics(t *testing.T) {
	db:= JsonDBImpl{}
	epics := db.SelectEpics()
	assert.Empty(t, epics)

	db.cached = tasks.Tasks{
		Epics: map[int]tasks.Epic{
			testEpic.ID: testEpic,
		},
	}
	epics = db.SelectEpics()
	assert.Equal(t, len(epics), 1)
	assert.Equal(t, epics[0], testEpic)
}

func TestJsonDB_SelectEpic(t *testing.T) {
	db:= JsonDBImpl{}
	epic, exists := db.SelectEpic(testEpic.ID)
	assert.False(t, exists)

	db.cached = tasks.Tasks{
		Epics: map[int]tasks.Epic{
			testEpic.ID: testEpic,
		},
	}
	epic, exists = db.SelectEpic(testEpic.ID)
	assert.True(t, exists)
	assert.Equal(t, epic, testEpic)
}

func TestJsonDB_SelectStories(t *testing.T) {
	db:= JsonDBImpl{}
	stories := db.SelectStories(testEpic.ID)
	assert.Empty(t, stories)

	db.cached = tasks.Tasks{
		Stories: map[int]tasks.Story{
			testStory.ID: testStory,
		},
	}
	stories = db.SelectStories(testEpic.ID)
	assert.Empty(t, stories)

	db.cached = tasks.Tasks{
		Epics: map[int]tasks.Epic{
			testEpic.ID: testEpic,
		},
		Stories: map[int]tasks.Story{
			testStory.ID: testStory,
		},
	}
	stories = db.SelectStories(testEpic.ID)
	assert.Equal(t, len(stories), 1)
	assert.Equal(t, stories[0], testStory)
}

func TestJsonDB_SelectStory(t *testing.T) {
	db:= JsonDBImpl{}
	story, exists := db.SelectStory(testStory.ID)
	assert.False(t, exists)

	db.cached = tasks.Tasks{
		Stories: map[int]tasks.Story{
			testStory.ID: testStory,
		},
	}
	story, exists = db.SelectStory(testStory.ID)
	assert.True(t, exists)
	assert.Equal(t, story, testStory)

	db.cached = tasks.Tasks{
		Epics: map[int]tasks.Epic{
			testEpic.ID: testEpic,
		},
		Stories: map[int]tasks.Story{
			testStory.ID: testStory,
		},
	}
	story, exists = db.SelectStory(testStory.ID)
	assert.True(t, exists)
	assert.Equal(t, story, testStory)
}

func TestJsonDB_UpsertEpic(t *testing.T) {
	db:= JsonDBImpl{}
	db.cached = tasks.Tasks{
		Epics: map[int]tasks.Epic{},
		Stories: map[int]tasks.Story{},
	}
	epics := db.SelectEpics()
	assert.Empty(t, epics)

	insertedEpic, err := db.UpsertEpic(testEpic)
	assert.Nil(t, err)
	assert.Equal(t, insertedEpic, testEpic)
	insertedEpic, exists := db.SelectEpic(testEpic.ID)
	assert.True(t, exists)
	assert.Equal(t, insertedEpic, testEpic)

	toBeUpdatedEpicTitle := testEpic.Title + "Updated"
	toBeUpdatedEpicDescription := testEpic.Title + "Updated"
	var toBeUpdatedEpicStatus tasks.TaskStatus
	if testEpic.Status == tasks.Open {
		toBeUpdatedEpicStatus = tasks.InProgress
	} else {
		toBeUpdatedEpicStatus = tasks.Open
	}
	var toBeUpdatedEpicStoryIDs []int
	if len(toBeUpdatedEpicStoryIDs) == 0 {
		toBeUpdatedEpicStoryIDs = []int{testStoryID}
	} else {
		toBeUpdatedEpicStoryIDs = []int{}
	}
	toBeUpdatedEpic := tasks.Epic{
		ID:          testEpic.ID,
		Title:       toBeUpdatedEpicTitle,
		Description: toBeUpdatedEpicDescription,
		Status:      toBeUpdatedEpicStatus,
		StoryIDs:    toBeUpdatedEpicStoryIDs,
	}
	updatedEpic, err := db.UpsertEpic(toBeUpdatedEpic)
	assert.Nil(t, err)
	assert.Equal(t, updatedEpic, toBeUpdatedEpic)
	updatedEpic, exists = db.SelectEpic(testEpic.ID)
	assert.True(t, exists)
	assert.Equal(t, updatedEpic, toBeUpdatedEpic)
}

func TestJsonDB_UpsertStory(t *testing.T) {
	db:= JsonDBImpl{}
	db.cached = tasks.Tasks{
		Epics: map[int]tasks.Epic{},
		Stories: map[int]tasks.Story{},
	}
	stories := db.SelectStories(testEpic.ID)
	assert.Empty(t, stories)

	insertedStory, err := db.UpsertStory(testStory)
	assert.Nil(t, err)
	assert.Equal(t, insertedStory, testStory)
	insertedStory, exists := db.SelectStory(testStory.ID)
	assert.True(t, exists)
	assert.Equal(t, insertedStory, testStory)

	toBeUpdatedStoryTitle := testStory.Title + "Updated"
	toBeUpdatedStoryDescription := testStory.Title + "Updated"
	var toBeUpdatedStoryStatus tasks.TaskStatus
	if testStory.Status == tasks.Open {
		toBeUpdatedStoryStatus = tasks.InProgress
	} else {
		toBeUpdatedStoryStatus = tasks.Open
	}
	toBeUpdatedStory := tasks.Epic{
		ID: testStory.ID,
		Title: toBeUpdatedStoryTitle,
		Description: toBeUpdatedStoryDescription,
		Status: toBeUpdatedStoryStatus,
	}
	updatedStory, err := db.UpsertEpic(toBeUpdatedStory)
	assert.Nil(t, err)
	assert.Equal(t, updatedStory, toBeUpdatedStory)
	updatedStory, exists = db.SelectEpic(testStory.ID)
	assert.True(t, exists)
	assert.Equal(t, updatedStory, toBeUpdatedStory)
}

func TestJsonDB_DeleteEpic(t *testing.T) {
	db:= JsonDBImpl{}
	db.cached = tasks.Tasks{
		Epics: map[int]tasks.Epic{},
		Stories: map[int]tasks.Story{},
	}
	epics := db.SelectEpics()
	assert.Empty(t, epics)

	insertedEpic, err := db.UpsertEpic(testEpic)
	assert.Nil(t, err)
	assert.Equal(t, insertedEpic, testEpic)
	insertedEpic, exists := db.SelectEpic(testEpic.ID)
	assert.True(t, exists)
	assert.Equal(t, insertedEpic, testEpic)

	err = db.DeleteEpic(testEpic.ID)
	assert.Nil(t, err)
	_, exists = db.SelectEpic(testEpic.ID)
	assert.False(t, exists)
}

func TestJsonDB_DeleteStory(t *testing.T) {
	db:= JsonDBImpl{}
	db.cached = tasks.Tasks{
		Epics: map[int]tasks.Epic{},
		Stories: map[int]tasks.Story{},
	}
	stories := db.SelectStories(testEpic.ID)
	assert.Empty(t, stories)

	insertedStory, err := db.UpsertStory(testStory)
	assert.Nil(t, err)
	assert.Equal(t, insertedStory, testStory)
	insertedStory, exists := db.SelectStory(testStory.ID)
	assert.True(t, exists)
	assert.Equal(t, insertedStory, testStory)

	err = db.DeleteStory(testStory.ID)
	assert.Nil(t, err)
	_, exists = db.SelectStory(testStory.ID)
	assert.False(t, exists)
}

func buildTestTasks(t *testing.T) (tasks.Tasks, error){
	content, err := ioutil.ReadFile(TestDataFilePath)
	if err != nil {
		t.Log("Error when opening file: ", err)
		return tasks.Tasks{}, err
	}

	var payload tasks.Tasks
	err = json.Unmarshal(content, &payload)
	if err != nil {
		t.Log("Error during unmarshal(): ", err)
		return tasks.Tasks{}, err
	}
	return payload, nil
}