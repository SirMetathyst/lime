watch:
	npm run tailwind:build:watch

debug:
	go build -o lime *.go
	./lime 2> ./lime.log