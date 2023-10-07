# JIRA Clone - Stage 2
__Navigation for Epics and Stories__

Here are the navigable pages to build for this stage:
* Epic List
* Epic Details
* Story Details

## Template
In the same folder as this README, you will find the corresponding template to get started with.

Please note - The template structure isn't the ONLY way to build/complete the project. You may:

1. Build the project following the template completely;
2. Ignore the provided template and build everything from the group up, then check the differences;
3. Take a more hybrid approach -- seek inspirations from the template, then establish your own style.

The structure of the template looks like this:
```text
/[stage 2 root folder]      <-- you're here
    |__ /main
           |__ main.go
    |__ /renderer
           |__ navigator.go
           |__ renderer.go
           |__screen.go
    |__ /storage
           |__ storage.go
           |__ storage.json
           |__ test_data.json
    |__ /tasks
           |__ tasks.go
    |__ go.mod
    |__ README.md
```

### Pages

In `/renderer/pages.go`, there are a few important structs:
* `EpicList` needs `Es []tasks.Epic` to display the list of created Epics.
* `EpicDetails` needs `E tasks.Epic` and `Ss []tasks.Story` to display details of the currently selected Epic and the list of created Stories under this Epic.
* `StoryDetails` needs `E tasks.Epic` for further navigation (going to the previous EpicDetails page, create and delete the story which is tied an Epic, etc.) and `S tasks.Story` to display details of the currently selected Story. 

These structs implement a common interface `Page`. `Page` has an interface method `show()` which (upon execution) will print out the aforementioned Epic/Story information, prompts, buttons etc.

There are helper functions of `newEpicList()`, `newEpicDetails()` and `newStoryDetails` to help create an instance of these structs. Consider using these helper functions instead when trying to create an instance of the corresponding structs. 

### Navigator

In `/renderer/navigator.go`, there are a few important functions/methods:

* `init()` init() gets kicked off on initialization of the file and creates a package singleton `var n Navigator`. You may want to check the singleton pattern if interested, but for the purpose of this project, it's already implemented and provided for. The gist is -- you can use this package variable to navigate through different `Page`'s, given that you've implemented the necessary methods.
* `Navigator.show()` (or its implementation of `NavigatorImpl.show()`) shows Epic/Story information, prompts, buttons etc. of the `CurrentPage`.
* `Navigator.navigate()` (or its implementation of `NavigatorImpl.navigate()`) helps navigate through different `Page`'s, after user reads the prompts by `Navigator.show()` and inputs what the prompts request.

### Renderer

In `/renderer/renderer.go`, there are a few important struct/function/method:
* `Renderer` is a struct that contains a `Navigator` instance.
* `init()` gets kicked off on initialization of the file and creates a global singleton `var R Renderer`, which `/main/main.go` calls to kick off UI rendering of the application. You may want to check the singleton pattern if interested, but for the purpose of this project, it's already implemented and provided for.
* `Renderer.Render()` has an infinite loop on two primary method calls, `Renderer.n.show()` and `Renderer.n.navigate()`. `Renderer.n.show()` give user instructions, and `Renderer.n.navigate()` prompts for user input after instructions are shown.

### Main

Finally, everything gets kicked off in `/main/main.go`, with the method call `Renderer.Render()` of the singleton `var R Renderer` in `/renderer/render.go`.

### Todos
There are `// Todo`'s throughout the template, which you're expected to complete. Again, keep in mind the template structure isn't the ONLY way of implementation. If you feel certain components should be rearranged, do so!

Also, consider discussing why you like (or didn't like) the approach the template suggests with others in the Bootcamp Discord channels!   