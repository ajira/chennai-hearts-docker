## User service

#### How to build?
 - docker build -t meetup_user_service .

#### Running the built container
 - docker run -it meetup_user_service /bin/sh
 - docker pull redis:3.0.7-alpine
 - docker run -p "6379:6379" redis:3.0.7-alpine
 - docker run -p "8080:8080" --link prickly_sammet:redis --env-file ./.env -it meetup_user_service

#### Using docker-compose
