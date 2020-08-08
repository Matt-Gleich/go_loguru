FROM golang:1.14

# Meta data:
LABEL maintainer="matthewgleich@gmail.com"
LABEL description="🌲 golang port of Delgan's python loguru"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

# Install make
RUN apt-get update && apt-get install make=4.2.1-1.2 -y --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

CMD ["make", "local-test"]
