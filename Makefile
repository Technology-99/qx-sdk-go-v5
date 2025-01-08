export ENDPOINT
export ACCESS_KEY_ID
export ACCESS_KEY_SECRET

test:
	go build -o example easy_example.go && ./example

build:
	go build -o example easy_example.go

clean:
	rm -rf build