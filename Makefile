export ENDPOINT
export ACCESS_KEY_ID
export ACCESS_KEY_SECRET

test:
	cd example && go build -o test easy_example.go && ./test

build:
	cd example && go build -o test easy_example.go

clean:
	rm -rf example/test