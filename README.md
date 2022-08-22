go test
go test -v
go test -cover

[Run in Git Bash on Windows]
go test -coverprofile=coverage.out && go tool cover -html=coverage.out