NAME=reversi-back
VERSION=1.0

.PHONY: build

build:
	DOCKER_BUILDKIT=1 docker build --secret id=credential,src=.gitconfig ./build -t $(NAME):$(VERSION)

start:
	docker run -itd -p 80:80 --name $(NAME) $(NAME):$(VERSION)

stop:
	docker rm -f $(NAME)

logs:
	docker logs $(NAME)