.PHONY: run dock drun

run:
	go run *.go -v

dock:
	docker build -t parzello .

drun:
	docker run -it \
		-v ~/.config/gcloud/:/gcloud \
		-e GOOGLE_APPLICATION_CREDENTIALS=/gcloud/application_default_credentials.json \
		parzello