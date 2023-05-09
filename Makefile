# Go parameters
GOCMD = go
GOBUILD = GOPRIVATE=$(GOPRIVATE) GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOMOD = $(GOCMD) mod
GOVET = $(GOCMD) vet
GOGENERATE = $(GOCMD) generate
GOINSTALL = $(GOCMD) install

setup/tools:
	$(GOINSTALL) github.com/golang/mock/mockgen@v1.6.0
vet:
	$(GOVET) ./...
test:
	$(GOTEST) -v ./...
