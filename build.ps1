if(-not (Test-Path -Path .\build)) {
    New-Item -Path "build" -ItemType Directory
}
go build -o build/app.exe # denotar que es un binario Windows
if ((Test-Path -Path .\fonts) -and  -not (Test-Path -Path .\build\fonts )) {
    Copy-Item -Path .\fonts -Destination .\build\fonts -Recurse
}