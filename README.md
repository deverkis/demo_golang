# demo_golang
学习Golang
mkdir hello

#To enable dependency tracking for your code, run the go mod init command, giving it the name of the module your code will be in.

cd hello
go mod init example/hello

#use the go mod edit command to edit the example.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is).
go mod edit -replace example.com/greetings=../greetings

#run the go mod tidy command to synchronize the example.com/hello module's dependencies
go mod tidy