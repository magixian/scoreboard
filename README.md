# Scoreboard

This is a simple command line application to calculate the scoreboard given results from games played.

The application currently will take results from a `.txt` file which will have the following format, 
```
[team a name] [score], [team b name] [score]
[team c name] [score], [team b name] [score]
...
```
A sample result file can be seen on this [sample results file](./sample_results.txt)

## How to run
There are 2 ways you can run the application, from the source code using golang or using a binary. There
is a prebuilt binary in the source `./bin/scoreboard` directory

1. Using source code
- Firstly you need to have golang installed on your machine.
- cd into the source code and run go mod tidy to make sure all the dependencies are available
- You will need to pass the path to your txt file as an argument
```
cd ./scoreboard
go mod tidy
go run cli/main.go ./sample_results.txt 
```

2. Using binary
- You use the prebuilt binary in the `./bin` directory or you can built on your local machine using go
```
go build -o ./bin/scoreboard ./cli/main.go
./bin/scoreboard ./sample_results.txt 
```