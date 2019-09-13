.PHONY: default build cui test

TARGET = fb.exe

defaut: build

build:
	go build -ldflags="-H windowsgui" -o $(TARGET)

cui:
	go build -o $(TARGET)

test: cui
	./$(TARGET)
