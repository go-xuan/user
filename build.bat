: linux
set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=0
go build -o quan_admin


: windows
set GOARCH=amd64
set GOOS=windows
set CGO_ENABLED=0
go build -o quan_admin.exe