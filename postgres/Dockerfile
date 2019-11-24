FROM postgres:12.1

RUN apt update && apt -y upgrade
RUN apt install -y git-all wget

# Install go https://tecadmin.net/install-go-on-debian/
RUN wget -q https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz \
    && tar -xvf go1.13.3.linux-amd64.tar.gz \
    && mv go /usr/local/go

ENV PATH="/usr/local/go/bin:${PATH}"

RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate && cp /root/go/bin/migrate /bin/migrate

# copy migration script to be run at startup
COPY run-migrations.sh /docker-entrypoint-initdb.d
