mkdir "./dist" 2> NUL

set GODEBUG=madvdontneed=1

:: Windows
set GOOS=windows
set GOARCH=amd64
go build -o ./dist/windows.exe ./example/main.go

:: Linux (x64)
set GOOS=linux
set GOARCH=amd64
go build -o ./dist/linux_amd64 ./example/main.go

:: Linux (ARM 64)
set GOOS=linux
set GOARCH=arm64
go build -o ./dist/linux_arm64 ./example/main.go

echo Finished!