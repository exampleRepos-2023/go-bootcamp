# JIRA Clone - Stage 1
__DAO for Epics and Stories__

Here are the todos:
1. Build Go models given the Test Data.
2. Build DAO, which performs CRUD operations on Epics and Stories.

## Template
In the same folder as this README, you will find the corresponding template to get started with.

Please note - The template structure isn't the ONLY way to build/complete the project. You may:

1. Build the project following the template completely;
2. Ignore the provided template and build everything from the group up, then check the differences;
3. Take a more hybrid approach -- seek inspirations from the template, then establish your own style.

The structure of the template looks like this:
```text
/[stage 1 root folder]      <-- you're here
    |__ /storage
           |__ storage.go
           |__ storage.json
           |__ test_data.json
    |__ /tasks
           |__ tasks.go
    |__ go.mod
    |__ README.md
```

### Test Data

In `/storage/test_data.json`, you can find how the database for this project is structured:
```json
{
  "last_task_id": 2,
  "epics": {
    "0": {
      "id": 0,
      "title": "Epic - Project 1",
      "description": "This is Project 1 for the Bootcamp",
      "status": "IN_PROGRESS",
      "story_ids": [
        2,
        3
      ]
    }
  },
  "stories": {
    "1": {
      "id": 1,
      "title": "Story - Project 1 Solution",
      "description": "Please provide full implement for Project 1",
      "status": "CLOSED"
    },
    "2": {
      "id": 2,
      "title": "Story - Project 1 README",
      "description": "Please create README file for Project 1",
      "status": "IN_PROGRESS"
    }
  }
}
```
These are the components:

`last_task_id`

This is an incrementing integer unique ID assign to an Epic or a Story on its creation.

    
For example, let's say that there are no previously created tasks (Epic/Story), and we are creating the first ever Epic. This Epic will be assigned with an ID of `0`. The next task to be created, can be an Epic or Story, an ID of `1` will be assigned.

`epics`

This is a map with Epic IDs as keys and Epics as values.

An Epic is composed of ID, title, description, status, and Story IDs.

The ID of an Epic is automatically assigned with `last_task_id` + 1.

Epic title and description are user-defined.

Epic status is user-defined and can be `OPEN`, `IN_PROGRESS`, `RESOLVED` or `CLOSED`

Story IDs are IDs of Stories created under this Epic.

`stories`

This is a map with Story IDs as keys and Stories as values.

A Story is composed of ID, title, description, and status.

The ID of an Story is automatically assigned with `last_task_id` + 1.

Story title and description are user-defined.

Story status is user-defined and can be `OPEN`, `IN_PROGRESS`, `RESOLVED` or `CLOSED`

A Story has to exist under one Epic.

### Models

`/tasks/tasks.go` is where you will create the Go database models based on the above JSON file.

There are 3 primary structs: `Epic`, `Story` and `Tasks`. `Tasks` is the composite object to store all the information (about Epics, Stories, etc.) from `storage.json` in runtime.

### Storage
There are 3 files in the `/storage` folder: `storage.go`, `storage.json` and `test_data.json`

In `storage/storage.go`, `JsonDB` is the interface with all the DAO methods, such as selectEpic, createStory. There are a few more important (and already implemented) functions/methods:
* `init()` gets kicked off on initialization of the file and creates a global singleton `var DB JsonDB`. You may want to check the singleton pattern if interested, but for the purpose of this project, it's already implemented and provided for. The gist is -- you can use this global variable to query against the JSON database.
* `JsonDBImpl.loadTasks()` loads tasks from a JSON file into `JsonDBImpl.cached`. If `var mode = Test`, it'll load it from `test_data.json`, which has some Test Data as shown above. If `var mode = Prod`, it'll load it from `/storage/storage.json`, which should be non-existent or empty when starting.
* `JsonDBImpl.saveTasks()` persists `JsonDBImpl.cached` into `/storage/storage.json`. It's supposed to be called upon creation, update and deletion. Otherwise, user can exist the application and lose all the recent changes. Remember -- you're meant to implement the CRUD operations, so keep this in mind.
* `getValuesSortedByKeys()` is a generic function. It takes a map as the parameter, sort the values of the map based the keys of the map, return the sorted values as the result. This function ensures our list select methods will return sorted lists. Here's an example:
    ```go
    m := map[int]string {
        1: "1",
        3: "3",
        2: "2",
    }
    fmt.Println(getValuesSortedByKeys(m)) // Should print out: ["1", "2", "3"], which is in sorted order.
    ```

Additionally, to use successfully use this template, make sure to update the value of `TestDataFilePath` and `DBFilePath` with the following steps:
1. Check the path of your current directory
    ```text
    $ pwd
    /Users/wallace/goworkspace/src/golangdojo.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage1
    ```
2. Update the value of `TestDataFilePath` and `DBFilePath` by appending `/storage/test_data.json` and `/storage/storage.json` to what `$ pwd` returns.
    ```text
    
    const (
        Prod Mode = iota
        Test
        TestDataFilePath = "/Users/wallace/goworkspace/src/golangdojo.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage1/storage/test_data.json"
        DBFilePath       = "/Users/wallace/goworkspace/src/golangdojo.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage1/storage/storage.json"
    )
    ```
3. Update the value of the package variable `mode` to `Test` or `Prod`. Update to `Test` to use the Test Data provided by `/storage/test_data.json`, or update to `Prod` to use `/storage/storage.json` which should be empty when starting fresh.
    ```text
    var mode = Prod
    ```

### Todos
There are `// Todo`'s throughout the template, which you're expected to complete. Again, keep in mind the template structure isn't the ONLY way of implementation. If you feel certain components should be rearranged, do so!

Also, consider discussing why you like (or didn't like) the approach the template suggests with others in the Bootcamp Discord channels!
