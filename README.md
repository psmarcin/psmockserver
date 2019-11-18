# psmockserver
psmockserver is a MockServer(https://mock-server.com) implementation in Go.


### Features

* [x] Test hooks
* [x] Integration tests https://github.com/psmarcin/psmockserver/pull/8
* [x] Support query strings https://github.com/psmarcin/psmockserver/issues/7
* [x] Base helm chart@v2 setup https://github.com/psmarcin/psmockserver/issues/12
* [ ] Clear mocks using cookies/headers https://github.com/psmarcin/psmockserver/issues/10
* [ ] Validate request body


### How To

#### Install dependencies
1. Go `>1.13` https://golang.org/
1. Realize https://github.com/oxequa/realize - for development 
1. Goreleaser https://goreleaser.com/ - for release
1. Docker https://www.docker.com/ - for test

#### Run server locally
1. `make dev` 

#### Run test locally
1. `make test` 

#### Run docker 
1. `docker build -t psmarcin/psmockserver .`
1. `docker run -d -p 8080:8080 psmarcin/psmockserver`

#### Run mock-tests
1. `make mock-test`

#### Load default mocks from file
psmockserver loads by default file `./default.json`.
