package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage1/tasks"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var mode = Test

type Mode int

const (
	Prod Mode = iota
	Test
	TestDataFilePath = "/Users/wallace/goworkspace/src/golangdojo.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage1/storage/test_data.json"
	DBFilePath       = "/Users/wallace/goworkspace/src/golangdojo.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage1/storage/storage.json"
)

type JsonDB interface {
	SelectLastTaskID() int
	SelectEpics() []tasks.Epic
	SelectEpic(epicID int) (tasks.Epic, bool)
	SelectStories(epicID int) []tasks.Story
	SelectStory(storyID int) (tasks.Story, bool)
	UpsertEpic(epic tasks.Epic) (tasks.Epic, error)
	UpsertStory(story tasks.Story) (tasks.Story, error)
	DeleteEpic(epicID int) error
	DeleteStory(storyID int) error
}

var DB JsonDB

func init() {
	dbImpl := JsonDBImpl{}
	err := dbImpl.loadTasks()
	if err != nil {
		fmt.Println("Error initializing DB", err)
		return
	}

	DB = &dbImpl
}

type JsonDBImpl struct {
	cached tasks.Tasks
}

func (db *JsonDBImpl) loadTasks() error {
	emptyTasks := tasks.Tasks{
		LastTaskID: tasks.NewTaskID,
		Epics: map[int]tasks.Epic{},
		Stories: map[int]tasks.Story{},
	}
	db.cached = emptyTasks

	switch mode {
	case Prod:
		initTasks, err := initProdDB()
		if err != nil {
			fmt.Println("Error initializing Prod DB", err)
			return err
		}
		db.cached = initTasks
	case Test:
		initTasks, err := initTestDB()
		if err != nil {
			fmt.Println("Error initializing Test DB", err)
			return err
		}
		db.cached = initTasks
		err = db.saveTasks()
		if err != nil {
			fmt.Println("Error saving test tasks", err)
			return err
		}
	default:
		err := errors.New("error asserting mode")
		fmt.Println("Error loading tasks", err)
		return err
	}
	return nil
}

func initProdDB() (tasks.Tasks, error) {
	emptyTasks := tasks.Tasks{
		LastTaskID: tasks.NewTaskID,
		Epics: map[int]tasks.Epic{},
		Stories: map[int]tasks.Story{},
	}
	emptyTasksJson, err := json.Marshal(emptyTasks)

	if _, err := os.Stat(DBFilePath); err != nil {
		err = ioutil.WriteFile(DBFilePath, emptyTasksJson, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating new storage.json", err)
			return emptyTasks, err
		}
	}

	content, err := ioutil.ReadFile(DBFilePath)
	if err != nil {
		fmt.Println("Error when opening file: ", err)
		return emptyTasks, err
	}

	if len(strings.Trim(string(content), " ")) == 0 {
		err = ioutil.WriteFile(DBFilePath, emptyTasksJson, os.ModePerm)
		if err != nil {
			fmt.Println("Error initializing empty storage.json", err)
			return emptyTasks, err
		}
		content, err = ioutil.ReadFile(DBFilePath)
		if err != nil {
			fmt.Println("Error reading initialized storage.json", err)
			return emptyTasks, err
		}
	}

	var contentTasks tasks.Tasks
	err = json.Unmarshal(content, &contentTasks)
	if err != nil {
		fmt.Println("Error during unmarshal()", err)
		return emptyTasks, err
	}

	return contentTasks, err
}

func initTestDB() (tasks.Tasks, error) {
	emptyTasks := tasks.Tasks{
		LastTaskID: tasks.NewTaskID,
		Epics: map[int]tasks.Epic{},
		Stories: map[int]tasks.Story{},
	}

	content, err := ioutil.ReadFile(TestDataFilePath)
	if err != nil {
		fmt.Println("Error when opening file: ", err)
		return emptyTasks, err
	}

	var contentTasks tasks.Tasks
	err = json.Unmarshal(content, &contentTasks)
	if err != nil {
		fmt.Println("Error during unmarshal()", err)
		return emptyTasks, err
	}

	return contentTasks, nil
}

func (db *JsonDBImpl) saveTasks() error {
	tasksJson, err := json.Marshal(db.cached)
	if err != nil {
		fmt.Println("Error marshalling tasks to be saved", err)
		return err
	}
	err = ioutil.WriteFile(DBFilePath, tasksJson, os.ModePerm)
	if err != nil {
		fmt.Println("Error writing tasks to be saved", err)
		return err
	}
	return nil
}

func (db *JsonDBImpl) SelectLastTaskID() int {
	// Todo
	return -1
}

func (db *JsonDBImpl) SelectEpics() []tasks.Epic {
	return getValuesSortedByKeys(db.cached.Epics)
}

func (db *JsonDBImpl) SelectEpic(epicID int) (tasks.Epic, bool) {
	// Todo
	return tasks.Epic{}, false
}

func (db *JsonDBImpl) SelectStories(epicID int) []tasks.Story {
	// Todo
	epicStories := make(map[int]tasks.Story)
	return getValuesSortedByKeys(epicStories)
}

func (db *JsonDBImpl) SelectStory(storyID int) (tasks.Story, bool) {
	// Todo
	return tasks.Story{}, false
}

func (db *JsonDBImpl) UpsertEpic(epic tasks.Epic) (tasks.Epic, error) {
	// Todo
	return tasks.Epic{}, errors.New("not implemented")
}

func (db *JsonDBImpl) UpsertStory(story tasks.Story) (tasks.Story, error) {
	// Todo
	return tasks.Story{}, errors.New("not implemented")
}

func (db *JsonDBImpl) DeleteEpic(epicID int) error {
	// Todo
	return errors.New("not implemented")
}

func (db *JsonDBImpl) DeleteStory(storyID int) error {
	// Todo
	return errors.New("not implemented")
}

func getValuesSortedByKeys[T any](m map[int]T) []T {
	keys := make([]int, 0, len(m))
	for key := range m{
		keys = append(keys, key)
	}
	sort.Ints(keys)

	sortedValues := make([]T, 0, len(m))
	for _, key := range keys {
		sortedValues = append(sortedValues, m[key])
	}
	return sortedValues
}

func example() {
	m := map[int]string {
		1: "1",
		3: "3",
		2: "2",
	}
	fmt.Println(getValuesSortedByKeys(m)) // Should print out: {"1", "2", "3"}, which is in sorted order.
}