FROM httpd

# Update and upgrade repo
RUN apt-get update -y -q && apt-get upgrade -y -q 

# Install tools we might need
RUN DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y -q curl build-essential ca-certificates git 

RUN curl -s https://storage.googleapis.com/golang/go1.18.1.linux-amd64.tar.gz | tar -v -C /usr/local -xz

ENV GOROOT=/usr/local/go
ENV PATH $PATH:/usr/local/go/bin

WORKDIR /var/www/http

COPY . .

RUN go get github.com/ajstarks/svgo

# CMD [“apache2ctl”, “-D”, “FOREGROUND”]