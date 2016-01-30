## User service

#### How to build?
 - docker build -t meetup_user .

#### Running the built container
 - docker run -it meetup_user /bin/sh
 - docker run -p "8080:8080" --env-file ./.env -it meetup_user /bin/sh
