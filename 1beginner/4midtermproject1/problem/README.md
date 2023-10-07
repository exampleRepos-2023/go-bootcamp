# JIRA Clone

## IMPORTANT NOTE

___Please read the project description thoroughly BEFORE getting started, especially the FAQs section.___

___Re-visit the project description multiple times DURING your design and development process, to ensure you're meeting the project requirements.___

## Problem Statement
We will build a JIRA clone for the terminal.

We will build two primary features in JIRA:
1. Epic CRUD 
2. Story CRUD

## Objective
In this project, we aim to learn and practice the following:
* Primitives
* Composite/custom data types
* Control flow (switch statements, functions/methods including `init()` and `main()`, etc.)
* Type declaration and assertion
* Basic `fmt.Println()`/`fmt.Printf()` `fmt.Scan()`/`fmt.Scanf()` functionalities
* Translating concepts into code
* Navigating and contributing to an existing code base

## FAQs
Here's a list of Frequently Asked Questions:

__Will there a template to build the project on top of?__

Yes. In the same folder as this README, you will find the corresponding templates to get started with. There are additional README's in those template folders to give your further instructions. 

__Should my implementation look exactly like the solution?__

There isn't ONE way to build the project. Ambiguous parts of projects are left for YOU to make corresponding technical decisions on. It's completely OK to come up with different approaches, as long as they are meeting the requirements.

__What if I get stuck and have questions?__

If you haven't already, join our Discord server and the exclusive Bootcamp channels as instructed on the Home page of the Bootcamp. Fire away your questions and find project partners over there!

__Why isn't anyone in the community able to answer my questions?__

Make sure you code is clean and refactored when asking specific implementation questions. Double-check your spelling and grammar. The answers you get will likely be as well as the way you present your questions. Also, try to be courteous, patient and understanding. We are all in this Go learning journey together!


## Terminologies

__Database / Storage / Persistence__

Database, storage and persistence are often used interchangeably. They represent the component we use to store and access information for the application.

__CRUD & DAO__

CRUD stands for actions of creation, read, update & deletion. DAO is often the interface in the application to perform these actions against the database. 

__JIRA, Epic & Story__

JIRA is an industry-standard tool for tracking progress of (not limited to) software projects. An Epic is usually used for entire initiatives, while a Story is a smaller unit for tasks with more specific instructions.

__Terminal & Terminal Applications__

A computer terminal is human-friendly text-based interface program, and it's often used to run other computer programs (may or may not be human-friendly). An application that's meant to run primarily on a terminal is known as a terminal application. 

## Recommendation
Here's a list of recommended action items to do before, during and after the development, to help you more effectively build the project and learn from the project.

Before Development:
* Create an account on JIRA. Try out the basic features (such as Epic & Story) over there. Learning the product and the user experience is CRITICAL for the success of a project.
* Read and understand the project description and requirements thoroughly. Often, mistakes can be avoided by carefully accessing what problems the project is supposed to solve and paying more attention to details.

During Development:
* Frequently check the project description/requirements, to make sure you are building what is asked of you.
* Build and test every step of the way, instead of doing so after implementing the entire project. It's much easier to debug smaller parts than the entire project. If you've failed to build and test frequently and everything is breaking/crashing, remove portions of your code and start over.
* Refactor as you implement. Keep your code clean and compartmentalized. Doing so makes debugging exponentially easier, as your implementation grows.

After Development:
* Run through the provided test cases, and fix any bugs! You are almost done, so finish the project strong!
* Refactor your code once more, and then post it on your GitHub profile. You're a Go developer now!
* Showcase your project to your friends and family (at the very least, to others in the Golang Dojo community)!

## Stages
The project is split into multiple stages. Please keep in mind, some implementation choices are made to minimize the scope of the project, so we can focus on the learning and implementing Go related concepts. Also, you will likely need to go back and refactor code from previous stages, as you add capabilities and encounter more different user scenarios.

### Stage 1

__DAO for Epics and Stories__

In backend development projects, the database design is often the very first task to complete and complete well. The database design determines what and how information is important and stored for repeated usages. While designing the database (what technologies to use, how to model it, etc.), we can very quickly assess if the project is a feasible idea and if we can meet most requirements.

In this project, we will persist Epic and Story records into a JSON file to keep things as simple as possible. The primary components are as follows:
* `last_task_id` - Integer ID counter for both Epics and Stories. Epics and Stories are both categorized as Tasks. Each Epic/Story creation will increment the counter.
* `epics` - A mapping between Epic IDs and the actual Epics. An Epic will consist a list of Stories in forms of Story IDs.
* `stories` - A mapping between Story IDs and the actual Stories.
* `epic` and `story` both have `id`, `title`, `description` and `status`.
* `status` can be `OPEN`, `IN_PROGRESS`, `RESOLVED` or `CLOSED`. 

Note - loading from and saving into the JSON file is already implemented and provided as part of the implementation template. We haven't learned marshalling and unmarshalling JSON files yet.

Here's an example for how that JSON file will look like:
```json
{
  "last_task_id": 2,
  "epics": {
    "0": {
      "id": 0,
      "name": "Epic - Project 1",
      "description": "This is Project 1 for the Bootcamp",
      "status": "IN_PROGRESS",
      "story_ids": [
        1,
        2
      ]
    }
  },
  "stories": {
    "1": {
      "id": 1,
      "name": "Story - Project 1 Solution",
      "description": "Please provide full implement for Project 1",
      "status": "CLOSED"
    },
    "2": {
      "id": 2,
      "name": "Story - Project 1 README",
      "description": "Please create README file for Project 1",
      "status": "IN_PROGRESS"
    }
  }
}
```


Here are the todos for Stage 1:

__1. Translate the JSON above into Go objects.__

Based on the above JSON file, we know this is the structure for Epic and Story:
```
Epic {
    ID
    Title
    Description
    Status
    StoryIDs
}

Story {
    ID
    Title
    Description
    Status
}
```
Please translate this structure into Go structs.

For example, we will need to create a Go struct to represent Story, which consists of `id`, `title`, `description` and `status`.

The Go struct for Story will look something similar to (may not be exactly the same as) this:
```go
type Story struct {
	ID int
	Title string
	Description string
	Status string
}
```
Hint - type declaration may be useful here.

__2. Create CRUD for Epics and Stories.__

Now that we have the Go objects/models created, we need to create a CRUD DAO to query these records from and to the database. Methods you need may include (but not limited to):
* SelectEpic
* InsertEpic
* UpdateEpic
* DeleteEpic
* SelectStory
* InsertStory
* UpdateStory
* DeleteStory

Depending on your exactly implementations, you could also select a list instead of individual records (eg. SelectEpics). Parameters passed in may vary. Upsert (insert + update) is also a common DAO method.

Note - `last_task_id`, the integer ID counter for Epics and Stories, needs to increment upon insertions of Epic/Story.

### Stage 2
__Navigation for Epics and Stories__

Here are the navigable pages for this project:
* Epic List
* Epic Details
* Story Details

__Epic List__

The Epic List page consists of 2 primary components:
1. List of Epics. This is the list of Epics created. Each list item will show the Epic task ID and the Epic title. The entire list will look like this:
    ```text
    Epic List:
    [1][IN_PROGRESS] Epic - Project 1 for the Bootcamp
    [4][OPEN] Epic - Project 2 for the Bootcamp
    [5][OPEN] Epic - Project 3 for the Bootcamp
    ```
    This list should be printed out as `stdout`, which can be achieved using `fmt.Println()` or `fmt.Printf()`.

2. Further navigation. This is a list of options for the user to further navigate through the application. The user can: create a new epic, or view more details on the Epic listed above: 
    
    ```text
    Epic List:
    [1][IN_PROGRESS] Epic - Project 1 for the Bootcamp
    [4][OPEN] Epic - Project 2 for the Bootcamp
    [5][OPEN] Epic - Project 3 for the Bootcamp
   
    Options:
    [C] Create an Epic    [D] Details on an Epic
    ```
    This user input prompt statement should be printed out as `stdout`, which can be achieved using `fmt.Println()` or `fmt.Printf()`.
 
    To create a new Epic, the user will need to input the Epic title and description.

    Getting the user input and storing it into Go variables can be achieved using `fmt.Scan()` or `fmt.Scanf()`.

    Note - status is defaulted to OPEN on task creation.

    ```text
    Options:
    [C] Create an Epic    [D] Details on an Epic
   
    Please select an option:
    C
   
    Please enter Epic title:
    "New Epic title"
   
    Please enter Epic description:
    "New Epic description"
    ```
    Note - Consider using literal `%q` (`fmt.Scanf("%q", input)`) to scan for inputs with double quotes.   

    After a new Epic is created, the Epic List page will be presented again:
    ```text
    Epic List:
    [1][IN_PROGRESS] Epic - Project 1 for the Bootcamp
    [4][OPEN] Epic - Project 2 for the Bootcamp
    [5][OPEN] Epic - Project 3 for the Bootcamp
    [6][OPEN] Epic - New Epic title
       
    Options:
    [C] Create an Epic    [D] Details on an Epic
    ```
    Alternatively, user can select to view more details on one of the listed Epics as follows:
    ```text
    Epic List:
    [1][IN_PROGRESS] Epic - Project 1 for the Bootcamp
    [4][OPEN] Epic - Project 2 for the Bootcamp
    [5][OPEN] Epic - Project 3 for the Bootcamp
   
    Options:
    [C] Create an Epic    [D] Details on an Epic

    Please select an option:
    D
   
    Please select Epic ID:
    1
    ```
    After an Epic is selected for more details. The Epic Details page will be presented next.

__Epic Details__

The Epic Details page consists of 3 primary components:

1. Details on the selected Epic (selected from the previous Epic List page). Details on the selected Epic, such as title, description, etc. will be listed like this:
    ```text
    Epic Details:
    ID:  1
    Title:  Epic - Project 1
    Description:  This is Project 1 for the Bootcamp
    Status:  IN_PROGRESS
    ```
2. List of Stories tied to the selected Epic. This is the list of Stories created for the selected Epic. Each list item will show the Story task ID and the Story title. The entire list will look like this:
    ```text
    Epic Details:
    ID:  1
    Title:  Epic - Project 1
    Description:  This is Project 1 for the Bootcamp
    Status:  IN_PROGRESS
   
    Stories:
    [2][OPEN] Story - Project 1 Solution
    [3][OPEN] Story - Project 1 README
    ```
    
3. Further navigation. This is a list of options for the user to further navigate through the application. The user can: go back the previous page, update the currently selected Epic, remove the currently selected Epic, create a new Story tied to this Epic, or view more details on one of the Stories listed above:
    ```text
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
    ```
    Selecting previous page should bring us back to the previous page -- the Epic List page.

    Next, updatable fields of the currently selected Epic are : title, description, and status:
    ```text
    Options:
    [P] Previous Page    [U] Update Epic    [R] Remove Epic    [C] Create a Story    [D] Details on a Story
    
    Please select an option:
    U
    
    Epic fields to update:
    [T] Title    [D] Description    [S] Status
    ```
    Updating title and description should be straight forward:
    ```text
    Epic fields to update:
    [T] Title    [D] Description    [S] Status
   
    Please select an option:
    T
    
    Current Epic description: 
    "Epic - Project 1"
    New Epic description:
    "Epic - Project 1 Updated"
   
    Epic Details:
    ID:  1
    Title:  Project 1 Updated
    Description:  This is Project 1 for the Bootcamp
    Status:  IN_PROGRESS

    Options:
    [P] Previous Page    [U] Update Epic    [R] Remove Epic    [C] Create a Story    [D] Details on a Story
    
    ```
    And updating status should list statuses to select from:
    ```text
    Epic fields to update:
    [T] Title    [D] Description    [S] Status
   
    Please select an option:
    S
    
    Current status: 
    OPEN
    New status: [O] OPEN    [I] IN_PROGRESS    [R] RESOLVED    [C] CLOSED
    R
   
    Epic Details:
    ID:  1
    Title:  Project 1 Updated
    Description:  This is Project 1 for the Bootcamp
    Status:  RESOLVED

    Options:
    [P] Previous Page    [U] Update Epic    [R] Remove Epic    [C] Create a Story    [D] Details on a Story
    ```

    Next, when removing the currently selected Epic, there needs to be a confirmation:
    ```text
    Options:
    [P] Previous Page    [U] Update Epic    [R] Remove Epic    [C] Create a Story    [D] Details on a Story    
    
    Please select an option:
    R
    
    Confirm to remove Epic: [1] Epic - Project 1
    (Note - All Stories under this Epic will also be removed)
    [Y] Yes    [n] No
    ```
    Note - Removing the Epic should also remove all the Stories tied to that Epic.
    
    If the currently selected Epic is confirmed to be deleted, the Epic List page should be presented again, now with one less Epic:
    ```text
    Epic List:
    [4][OPEN] Epic - Project 2 for the Bootcamp
    [5][OPEN] Epic - Project 3 for the Bootcamp
    ```
    There should be similar prompts when creating a new Story tied to the currently selected Epic:
    ```text
    Options:
    [P] Previous Page    [U] Update Epic    [R] Remove Epic    [C] Create a Story    [D] Details on a Story    
    
    Please select an option:
    C
    
    Please enter Story title:
    "New Story title"
    
    Please enter Story description:
    "New Story description"
    ```
    After a new Story is created, we should see Epic Details page again:
    ```text
    Epic Details:
    ID:  4
    Title:  Epic - Project 2
    Description:  This is Project 2 for the Bootcamp
    Status:  OPEN
    
    Stories:
    [7][OPEN] Story - Project 2 Solution
    [8][OPEN] Story - Project 2 README
    [9][OPEN] New Story title
    
    Options:
    [P] Previous Page    [U] Update Epic    [R] Remove Epic    [C] Create a Story    [D] Details on a Story    
    ```
    Last but not least, you can view more details on any of the created Stories tied to the currently selected Epic:
    ```text
    Options:
    [P] Previous Page    [U] Update Epic    [R] Remove Epic    [C] Create a Story    [D] Details on a Story    
    
    Please select an option:
    D
    
    Please select Story ID:
    9
    ```
    And the Story Details should be presented next.

__Story Details__

The Story Details page consists of 2 primary components:

1. Details on the selected Story (selected from the previous Epic Details page). Details on the selected Story, such as title, description, etc. will be listed like this:

    ```text
    Story Details:
    ID:  9
    Title:  New Story title
    Description:  New Story description
    Status:  OPEN
    ```
2. Further navigation. Options are similar to the Epic Details page:

    ```text
    Story Details:
    ID:  9
    Title:  New Story title
    Description:  New Story description
    Status:  OPEN
    
    Options:
    [P] Previous Page    [U] Update Story    [R] Remove Story
    ```
    Selecting previous page should bring us back to the previous page -- the Epic Details page.
    Updating and removing Story options should operate similarly to updating and remove Epic options. 

### Stage 3
__Bug Bash__

This is the stage to more thoroughly test out your implementation and fix any additionally discovered bugs.

Make sure your implementation follows these requirements:
* When you create a task (Epic/Story), `last_task_id` + 1 is assigned to the Epic or Story as the ID.
* When you create a task (Epic/Story), `last_task_id` in the database (our JSON file) also gets incremented.
* When you create/update/delete tasks (Epic/Story), `epics` and/or `stories` in the database (our JSON file) also gets updated accordingly.
* When you view details on an Epic, all Stories tied to that Epic are listed.
* When you delete an Epic, all Stories tied to that Epic are also deleted.
* Your implementation works perfectly with the JSON database design already provided.
* User can terminate the application at any point, and previous creations, updates and deletions persist.
* Test your application against the provided Test Cases, and the resulting JSON file matches 100% with expectations.
* Fix any bugs!

## Test Cases
__Create Epic__

Steps:
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /main
      |__ /renderer
      |__ /storage
      |__ ...
    ```
* Set `mode=Prod` in `storage/storage.go`
* Run `rm storage/storage.json`
* Run `go run main/main.go`
* Input `C` to create a new Epic
* Input `"New Epic title"` as Epic title
* Input `"New Epic description"` as Epic description
* Check if storage.json matches with the following:
     ```json
    {"last_task_id":0,"epics":{"0":{"id":0,"title":"New Epic title","description":"New Epic description","status":"OPEN","story_ids":null}},"stories":{}}
     ```

__Create Story__

Steps:
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /main
      |__ /renderer
      |__ /storage
      |__ ...
    ```
* Set `mode=Prod` in `storage/storage.go`
* Run `rm storage/storage.json`
* Run `go run main/main.go`
* Input `C` to create a new Epic
* Input `"New Epic title"` as Epic title
* Input `"New Epic description"` as Epic description
* Input `D` to view details on an Epic
* Input `0` to select the created Epic
* Input `C` to create a new Story in the selected Epic
* Input `"New Story title"` as Story title
* Input `"New Story description"` as Story description
* Check if storage.json matches with the following:
    ```json
    {"last_task_id":1,"epics":{"0":{"id":0,"title":"New Epic title","description":"New Epic description","status":"OPEN","story_ids":[1]}},"stories":{"1":{"id":1,"title":"New Story title","description":"New Story description","status":"OPEN"}}}
    ```

__Update Epic__

Steps:
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /main
      |__ /renderer
      |__ /storage
      |__ ...
    ```
* Set `mode=Prod` in `storage/storage.go`
* Run `rm storage/storage.json`
* Run `go run main/main.go`
* Input `C` to create a new Epic
* Input `"New Epic title"` as Epic title
* Input `"New Epic description"` as Epic description
* Input `D` to view details on an Epic
* Input `0` to select the created Epic
* Input `U` to update the selected Epic
* Input `T` to select title to update
* Input `"Updated Epic title"` as updated title
* Input `U` to update the selected Epic
* Input `D` to select description to update
* Input `"Updated Epic description"` as updated description
* Input `U` to update the selected Epic
* Input `S` to select status to update
* Input `I` to select `IN_PROGRESS` as updated status
* Check if storage.json matches with the following:
    ```json
    {"last_task_id":0,"epics":{"0":{"id":0,"title":"Updated Epic title","description":"Updated Epic description","status":"IN_PROGRESS","story_ids":null}},"stories":{}}
    ```

__Update Story__

Steps:
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /main
      |__ /renderer
      |__ /storage
      |__ ...
    ```
* Set `mode=Prod` in `storage/storage.go`
* Run `rm storage/storage.json`
* Run `go run main/main.go`
* Input `C` to create a new Epic
* Input `"New Epic title"` as Epic title
* Input `"New Epic description"` as Epic description
* Input `D` to view details on an Epic
* Input `0` to select the created Epic
* Input `C` to create a new Story in the selected Epic
* Input `"New Story title"` as Story title
* Input `"New Story description"` as Story description
* Input `D` to view details on a Story
* Input `1` to select the created Story
* Input `U` to update the selected Story
* Input `T` to select title to update
* Input `"Updated Story title"` as updated title
* Input `U` to update the selected Story
* Input `D` to select description to update
* Input `"Updated Story description"` as updated description
* Input `U` to update the selected Story
* Input `S` to select status to update
* Input `I` to select `IN_PROGRESS` as updated status
* Check if storage.json matches with the following:
    ```json
    {"last_task_id":1,"epics":{"0":{"id":0,"title":"New Epic title","description":"New Epic description","status":"OPEN","story_ids":[1]}},"stories":{"1":{"id":1,"title":"Updated Story title","description":"Updated Story description","status":"IN_PROGRESS"}}}
    ```

__Remove Epic__

Steps:
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /main
      |__ /renderer
      |__ /storage
      |__ ...
    ```
* Set `mode=Prod` in `storage/storage.go`
* Run `rm storage/storage.json`
* Run `go run main/main.go`
* Input `C` to create a new Epic
* Input `"New Epic title"` as Epic title
* Input `"New Epic description"` as Epic description
* Input `D` to view details on an Epic
* Input `0` to select the created Epic
* Input `C` to create a new Story in the selected Epic
* Input `"New Story title"` as Story title
* Input `"New Story description"` as Story description
* Input `R` to remove the selected Epic
* Input `Y` to confirm removal
* Check if storage.json matches with the following:
    ```json
    {"last_task_id":1,"epics":{},"stories":{}}
    ```

__Remove Story__

Steps:
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /main
      |__ /renderer
      |__ /storage
      |__ ...
    ```
* Set `mode=Prod` in `storage/storage.go`
* Run `rm storage/storage.json`
* Run `go run main/main.go`
* Input `C` to create a new Epic
* Input `"New Epic title"` as Epic title
* Input `"New Epic description"` as Epic description
* Input `D` to view details on an Epic
* Input `0` to select the created Epic
* Input `C` to create a new Story in the selected Epic
* Input `"New Story title"` as Story title
* Input `"New Story description"` as Story description
* Input `D` to view details on a Story
* Input `1` to select the created Story
* Input `R` to remove the selected Story
* Input `Y` to confirm removal
* Check if storage.json matches with the following:
    ```json
    {"last_task_id":1,"epics":{"0":{"id":0,"title":"New Epic title","description":"New Epic description","status":"OPEN","story_ids":null}},"stories":{}}
    ```
  
## Final Message
Check this -- you're a Go developer now!

This is a pretty elaborate project for a beginner. You should be proud of your progress if you've gotten this far.

Showcase your implementation and struggles you've faced along the way to others in the Golang Dojo community.

More importantly, teaching is the best way to learn. Any questions posted by others in the Discord channels are opportunities for you to answer and truly internalize your knowledge. 

Congrats! And let's get started with the next modules and corresponding projects! 