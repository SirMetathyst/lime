watch:
	npm run tailwind:build:watch

debug:
	go build -o lime ./cmd/lime/*.go
	./lime 2> ./lime.log