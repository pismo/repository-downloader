Download all organization repos to a folder

`go get github.com/pismo/repository-downloader`

Run as:
`repository-downloader https://api.github.com/orgs/ORGANIZATION/repos YOURGITHUBTOKEN /home/MYUSER/folder`

If you dont have GOBIN set up in your PATH, then do as follow:
`go run $GOPATH/src/github.com/pismo/repository-downloader/application.go https://api.github.com/orgs/ORGANIZATION/repos YOURGITHUBTOKEN /home/MYUSER/folder`

