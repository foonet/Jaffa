#!/bin/bash

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_darwin.amd64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_linux.386 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_linux.amd64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_linux.arm main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_linux.arm64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_linux.mips main.go
CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_linux.mips main.go
CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_linux.mipsle main.go
CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_linux.mips64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_linux.mips64le main.go
CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_freebsd.386 main.go
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_freebsd.amd64 main.go
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_windows.386 main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o jaffa_windows.amd64 main.go
