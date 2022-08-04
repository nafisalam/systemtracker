# Project

This is a system monitoring application to display a list of files and size in a directory in json formate. This application uses a cli framwork -COBRA.

This application has below folders structure.

1. .conf : It is a configuration related folder and it has all the environment and application related config parameters. Currently. there is parameters defined in the directory are not used anywhere.

2. cmd : It has the commands

3. pkg : It has the implementation of the commands.

Prerequisite:

1. Install Go library in you system.
2. Setup the Go path.

Instructions to run the application locally:

 1. Check out the code from github.

 2. Go to the root folder of the application

 3. Run the command go mod vendor

 4. Run the command go run main.go /track /[foldername] . you need to provide the complete folder  path in the above command. e.g - go run --race main.go /track /Users/nafisalam/go

 6. To run the unit test , Please run below commands
    1. go to the folder pkg/tracker
    2. go test -run tracker_test.go 
    2. go test ./... -coverprofile cover.out
    3. go tool cover -func cover.out

 
