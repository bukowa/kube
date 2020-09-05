
# templates can behave differently for compiled programs
testing.templated.build:
	docker build -f testing.templated.build.Dockerfile .

gotest:
	go test ./...

test-docker: testing.templated.build


test: gotest test-docker