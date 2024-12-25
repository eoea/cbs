build:
	go build -o ~/.local/bin/cbs ./src/main.go

driver_update:
	go run github.com/playwright-community/playwright-go/cmd/playwright@latest install firefox
