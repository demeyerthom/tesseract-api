.PHONY: generate

generate: # Generates the server
	swagger generate server \
		--server-package ./internal \
		--model-package ./internal/models \
		--target= ./internal \
		--name tesseract-api \
		--exclude-main