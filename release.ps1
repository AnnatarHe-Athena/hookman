Write-Host "😁 building & uploading"

$env:GOOS = "linux"
$env:GOARCH = "amd64"

go build -ldflags="-s -w" -tags release -o hookman.exe .\main.go
scp .\hookman.exe annatarhe@annatarhe.cn:~/

Write-Host "✨ done"