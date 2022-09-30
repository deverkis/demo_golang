#init
go mod init example/fuzz

#Excute
go run .

#Test
go test

#Test fuzz

go test -fuzz=Fuzz