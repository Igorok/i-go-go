### installation of golang version manager
```
xcode-select --install

bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
gvm install go1.14.1 -B
gvm list
gvm use go1.14.2

https://medium.com/golang-notes/%D0%BD%D0%B0%D1%81%D1%82%D1%80%D0%BE%D0%B9%D0%BA%D0%B0-visual-studio-code-%D0%B4%D0%BB%D1%8F-go-647ea94aa795

go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v golang.org/x/lint/golint
go get -u -v github.com/lukehoban/go-find-references
go get -u -v github.com/lukehoban/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/newhook/go-symbols

vscode
{
    "go.buildOnSave": true,
    "go.lintOnSave": true,
    "go.vetOnSave": true,
    "go.buildFlags": [],
    "go.lintFlags": [],
    "go.vetFlags": [],
    "go.useCodeSnippetsOnFunctionSuggest": false,
    "go.formatOnSave": false,
    "go.formatTool": "goreturns"
}

debug
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceRoot}",
            "env": {},
            "args": []
        }
    ]
}

```

### initialization of default mongodb schema
```
go run bin/init/init_mongodb.go
```

### graphiql interface
```
http://localhost:8081/graphql

query signIn($login: String, $password: String) {
  signIn(login: $login, password: $password) {
    token
  }
}

{
  "login": "operator_pizza",
  "password": "operator_pizza"
}
```

### testing
```
go get github.com/smartystreets/goconvey
$GOPATH/bin/goconvey
http://localhost:8080
```
