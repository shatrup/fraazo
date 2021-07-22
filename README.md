# fraazo
Web application which loads a CSV on startup and then exposes an API to explore the file.

### clone the project
  `git clone https://github.com/FoodDarzee/take-home.git`
  
### build the docker image
first go to the root directory of the project then type   
`docker build -t fraazo .`

### run the docker image
`docker run -it -p 8082:8082 fraazo`

### build and run the project
`go build && go run main.go`

### Run Tests
For details refer: https://stackoverflow.com/questions/16353016/how-to-go-test-all-testings-in-my-project

This should run all tests in current directory and all of its subdirectories:
`go test ./...`
