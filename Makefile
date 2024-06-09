build:
	.\tailwindcss -i src/input.css -o public/output.css
	@templ generate
	@go build -o bin/$(APP_NAME) cmd/main.go
	@echo "application compiled successfully" 
	

run:build
	air


templ-watch:
	templ generate --watch


tailwind-watch:
	.\tailwindcss -i src/input.css -o public/output.css --watch
