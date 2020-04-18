Write-Host "😁 building & uploading"

$env:GOOS = "linux"
$env:GOARCH = "amd64"

go build -tags release -o hookman.exe .\main.go
scp .\hookman.exe annatarhe@annatarhe.cn:~/

Write-Host "✨ done"