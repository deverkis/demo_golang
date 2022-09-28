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

#This topic introduces two additional commands for building code:

The go build command compiles the packages, along with their dependencies, but it doesn't install the results.
The go install command compiles and installs the packages.

#You can discover the install path by running the go list command, as in the following example:

$ go list -f '{{.Target}}'

#run the go install command to compile and install the package
go install 


