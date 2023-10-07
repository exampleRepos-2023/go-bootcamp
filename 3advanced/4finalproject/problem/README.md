# StackOverflow Clone

## IMPORTANT NOTE

___Please read the project description thoroughly BEFORE getting started, especially the FAQs section.___

___Re-visit the project description multiple times DURING your design and development process, to ensure you're meeting the project requirements.___

## Problem Statement
We will build the backend for StackOverflow.

We will build two primary features in StackOverflow:
1. Question creation, retrieval & deletion
2. Answer creation, retrieval & deletion

## Objective
In this project, we aim to learn and practice the following:
* Designing SQL models
* Hands-on usage of Postgres
* Designing & building APIs
* Medium-sized project structure
* Translating concepts into code
* Navigating and contributing to an existing code base

## FAQs
Here's a list of Frequently Asked Questions:

__Will there a template to build the project on top of?__

Yes. In the same folder as this README, you will find the corresponding templates to get started with. There are additional README's in those template folders to give your further instructions.

__Should my implementation look exactly like the solution?__

There isn't ONE way to build the project. Ambiguous parts of projects are left for YOU to make corresponding technical decisions on. It's completely OK to come up with different approaches, as long as they are meeting the requirements. However, if you follow the included structure, you get the extra benefit of utilizing the included unit tests. 

__What if I get stuck and have questions?__

If you haven't already, join our Discord server and the exclusive Bootcamp channels as instructed on the Home page of the Bootcamp. Fire away your questions and find project partners over there!

__Why isn't anyone in the community able to answer my questions?__

Make sure you code is clean and refactored when asking specific implementation questions. Double-check your spelling and grammar. The answers you get will likely be as well as the way you present your questions. Also, try to be courteous, patient and understanding. We are all in this Go learning journey together!


## Terminologies

__API / Endpoint__

API and endpoint are often used interchangeably. They are the backend component that acts as the point of contact for frontend clients. The design of the APIs or endpoints determines how the frontend and backend communicate with one another.

__Model__
Models describe how information is organized, transmitted or stored. Data models describe fields of a data table.

__CRUD & DAO__

CRUD stands for actions of creation, read, update & deletion. DAO is often the interface in the application to perform these actions against the database.

__StackOverflow, Question & Answer__

StackOverflow is an industry-standard forum for posting programming related questions and getting crowd-sourced answers. A user posts a question under the user's account, and other users can contribute different answers as they see most appropriate per question. There are additional features on StackOverflow, such as upvoting the question or its corresponding answers. However, question and answer creation, retrieval and deletion are the most primitive.

## Recommendation
Here's a list of recommended action items to do before, during and after the development, to help you more effectively build the project and learn from the project.

Before Development:
* Create an account on StackOverflow. Try out the basic features (such as posting a question & an answer) over there. Learning the product and the user experience is CRITICAL for the success of a project.
* Read and understand the project description and requirements thoroughly. Often, mistakes can be avoided by carefully accessing what problems the project is supposed to solve and paying more attention to details.

During Development:
* Frequently check the project description/requirements, to make sure you are building what is asked of you.
* Build and test every step of the way, instead of doing so after implementing the entire project. It's much easier to debug smaller parts than the entire project. If you've failed to build and test frequently and everything is breaking/crashing, remove portions of your code and start over.
* Utilize the included unit tests to help debug your implementation.
* Refactor as you implement. Keep your code clean and compartmentalized. Doing so makes debugging exponentially easier, as your implementation grows.

After Development:
* Execute the provided unit tests, and fix any bugs! Unit tests often help catch smaller and more granular mistakes. 
* Run through the provided test cases, and fix any bugs! You are almost done, so finish the project strong!
* Refactor your code once more, and then post it on your GitHub profile. You're a Go web developer now!
* Showcase your project to your friends and family (at the very least, to others in the Golang Dojo community)!

## Stages
The project is split into multiple stages. Please keep in mind, some implementation choices are made to minimize the scope of the project, so we can focus on the learning and implementing Go related concepts. Also, you will likely need to go back and refactor code from previous stages, as you add capabilities and encounter more different user scenarios.

### Stage 1

__API (endpoints & models)__

In any sizable project with a frontend interface, API endpoints are among the first things to be designed and designed well. This API design is then used as a contrast for both frontend engineers and backend engineers to enforce during their more independent development.

For this project, the API endpoints are to be implemented by following the design below:

**Question creation** 
```
POST /question
```

Request

| Name        | Type   | Description                               |
|-------------|--------|-------------------------------------------|
| title       | String | Title of the question to be created       |      
| description | String | Description of the question to be created |

Response

| Name        | Type     | Description                                                    |
|-------------|----------|----------------------------------------------------------------|
| question    | Question | Question created. See below for the data structure of Question |      
| status_code | Int      | 201 for success, 503 for failure                               |      

Question:

| Name          | Type   | Description                                   |
|---------------|--------|-----------------------------------------------|
| question_uuid | UUID   | Question UUID generated per question creation |      
| title         | String | Title of the question created                 |      
| description   | String | Description of the question created           |      
| created_at    | Long   | Timestamp of the question created             |      


Question retrieval
```
GET /questions
```

Request

| Name | Type | Description |
|------|------|-------------|
| N/A  | N/A  | N/A         |

Response

| Name        | Type       | Description                                                               |
|-------------|------------|---------------------------------------------------------------------------|
| questions   | []Question | Slice of all questions created, each consisting a title and a description |
| status_code | Int        | 200 for success, 503 for failure                                          |      



Question deletion
```
DELETE /question
```

Request

| Name          | Type   | Description                         |
|---------------|--------|-------------------------------------|
| question_uuid | UUID   | Title of the question to be deleted |      

Response

| Name        | Type | Description                                                                |
|-------------|------|----------------------------------------------------------------------------|
| status_code | Int  | 200 for success, 404 for question to be deleted not found, 503 for failure |      

For this stage, we are to implement the API endpoints and models as described above, and get to the point where we can print out requests via the web server and return some dummy responses.

### Stage 2

__Persistence (models & connection)__

In backend development projects, after the API contract is established, the database design is often the very next task to complete and complete well. The database design determines what and how information is important and stored for repeated usages. While designing the database (what technologies to use, how to model it, etc.), we can very quickly assess if the project is a feasible idea and if we can meet most requirements.

In this project, we will persist Question and Answer records into a SQL database. The primary components are as follows:

Question

| Name          | Type         | Description                                  |
|---------------|--------------|----------------------------------------------|
| question_uuid | VARCHAR(36)  | Generated identifier unique to each question |
| title         | VARCHAR(255) | Title of the question                        |
| description   | VARCHAR(255) | Description of the question                  |
| created_at    | BIGINT       | Creation timestamp of the question           |

Answer

| Name        | Type         | Description                                |
|-------------|--------------|--------------------------------------------|
| answer_uuid | VARCHAR(36)  | Generated identifier unique to each answer |
| content     | VARCHAR(255) | Content of the answer                      |
| created_at  | BIGINT       | Creation timestamp of the answer           |

QuestionAnswers

| Name          | Type         | Description                                  |
|---------------|--------------|----------------------------------------------|
| question_uuid | VARCHAR(36)  | Generated identifier unique to each question |
| answer_uuid   | VARCHAR(36)  | Generated identifier unique to each answer   |

Note - The SQL queries for creating and dropping tables are included in the template.

For this stage, here are the todos:
* Create persistence models as described above.
* Establish Postgres database connection.
* Drop, if exists, and create Question, Answer & QuestionAnswers tables.

### Stage 3
__Persistence (DAOs)__

Now that we have the models and connection created, it's time to create the DAOs for creation, retrieval & deletion for Question, Answer & QuestionAnswers.

Note - The SQL queries are included in the template, but the where clauses are to be added. For example, let's say you need to search for a Question given the question_uuid, `SELECT question_uuid, title, description, created_at from questions` is provided but `WHERE question_uuid IN (a2f5338f-c7a6-472a-8e9b-1c8192665bb0)` needs to be appended.

### Stage 4
__Connecting endpoints with DAOs__

Now that both endpoints and DAOs are created, it's time to connect these two major components.

Remember endpoints are the entry point for receive frontend client requests, we need to query the database based on the details provided in the requests to perform creation, retrieval & deletion. These operations will take place on top of the DAOs created.

### Stage 5
__Bug Bash__

This is the stage to more thoroughly test out your implementation and fix any additionally discovered bugs.

Make sure your implementation follows these requirements:
* When you create a question/answer, a random UUID is assigned and persisted with that question/answer.
* When you create a question/answer, a timestamp for when the question/answer is created is assigned and persisted with that question/answer.
* When you create/delete a question/answer, it is indeed create and a retrieval reflects of said question/answer reflects that database record change.
* When you view questions via `GET /questions`, all questions created are listed.
* When you view answers of a question via `GET /answers`, all answers created under that question are listed.
* When you delete a question, all answers tied to that question are also deleted.
* User can terminate the application at any point, and previous creations & deletions persist.
* Test your application against the provided Test Cases, and the resulting database records match 100% with expectations.
* Fix any bugs!

#### Sample Requests
Create question

```text

curl --request POST \
  --url http://localhost/question \
  --header 'Accept: application/json' \
  --data '{
	"title": "Title",
	"description": "Description"
}'
```

Get questions

```text
curl --request GET \
  --url http://localhost/questions \
  --header 'Accept: application/json'
```

Delete question

```text
curl --request DELETE \
  --url http://localhost/question \
  --header 'Accept: application/json' \
  --data '{
	"question_uuid": "[UUID of a created question]"
}'
```

Create answer

```text
curl --request POST \
  --url http://localhost/answer \
  --header 'Accept: application/json' \
  --data '{
	"question_uuid": [UUID of a created question],
	"content": "Content"
}'
```

Get answers

```text
curl --request GET \
  --url http://localhost/answers \
  --header 'Accept: application/json' \
  --data '{
	"question_uuid": [UUID of a created question]
}'
```

Delete answer

```text
curl --request DELETE \
  --url http://localhost/answer \
  --header 'Accept: application/json' \
  --data '{
	"answer_uuid": [UUID of a created answer]
}'
```

#### Test Cases

__Get empty questions__
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /handlers
      |__ /main
      |__ /managers
      |__ /models
      |__ /persistence
      |__ ...
    ```
* Run `cd main`
* Run `go run main.go`
* Get questions, which should be empty.

__Create question__
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /handlers
      |__ /main
      |__ /managers
      |__ /models
      |__ /persistence
      |__ ...
    ```
* Run `cd main`
* Run `go run main.go`
* Get questions, which should be empty.
* Create one question, which should return question created.
* Get questions, which should return one question created.
* Create another question, which should return question created.
* Get questions, which should return two questions created.

__Delete question__
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /handlers
      |__ /main
      |__ /managers
      |__ /models
      |__ /persistence
      |__ ...
    ```
* Run `cd main`
* Run `go run main.go`
* Get questions, which should be empty.
* Create one question, which should return question created.
* Get questions, which should return one question created.
* Delete non-existing question, which shouldn't do anything.
* Create one question, which should return question created.
* Delete existing question.
* Get questions, which should be empty.

__Get empty answers__
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /handlers
      |__ /main
      |__ /managers
      |__ /models
      |__ /persistence
      |__ ...
    ```
* Run `cd main`
* Run `go run main.go`
* Get answers with non-existing question, which should be empty
* Create one question, which should return question created.
* Get answers with existing question, which should be empty.

__Create answer__
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /handlers
      |__ /main
      |__ /managers
      |__ /models
      |__ /persistence
      |__ ...
    ```
* Run `cd main`
* Run `go run main.go`
* Create one question, which should return question created.
* Get answers with non-existing question, which should be empty.
* Get answers with existing question with no answers, which should be empty.
* Create answer with non-existing question, which shouldn't do anything.
* Get answers with non-existing question, which should be empty.
* Create answer with existing question, which should return answer created.
* Get answers with existing question, which should return answer created.

__Delete answer__
* `cd` into the root folder of project
    ```text
    /[root folder]    <--here
      |__ /handlers
      |__ /main
      |__ /managers
      |__ /models
      |__ /persistence
      |__ ...
    ```
* Run `cd main`
* Run `go run main.go`
* Create one question, which should return question created.
* Get answers with existing question, which should be empty.
* Create one answer with existing question, which should return answer created.
* Get answers with existing question, which should return one answer created.
* Create another answer with existing question, which should return answer created.
* Get answers with existing question, which should return two answers created.
* Delete non-existing answer, which shouldn't do anything.
* Get answers with existing question, which should return two answers created.
* Delete one existing answer under existing question.
* Get answers with existing question, which should return one answers created.
* Delete another existing answer under existing question.
* Get answers with existing question, which should be empty.

## Final Message
Check this -- you're a Go web developer now!

Showcase your implementation and struggles you've faced along the way to others in the Golang Dojo community.

More importantly, teaching is the best way to learn. Any questions posted by others in the Discord channels are opportunities for you to answer and truly internalize your knowledge.

Congrats! And let's get started with the Masterclasses!