default: out/painter

clean:
	rm -rf out

test:
	go test ./...

out/painter:
	mkdir -p out
	go build -o out/example ./cmd/example