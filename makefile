.PHONY: default build cui

TARGET = fb.exe

defaut: build

build:
	go build -ldflags="-H windowsgui" -o $(TARGET)

cui:
	go build -o $(TARGET)
