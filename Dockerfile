FROM golang:latest AS build

# Создаем рабочую директорию и компилим
ADD . /app
WORKDIR /app
RUN go build ./cmd/main/main.go

FROM ubuntu:20.04
COPY . .

# Настройка даты
RUN apt-get -y update && apt-get install -y tzdata
ENV TZ=Russia/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Установка базы
ENV PGVER 12
RUN apt-get -y update && apt-get install -y postgresql-$PGVER
USER postgres

# Настройка базы
RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER andeo WITH SUPERUSER PASSWORD 'andeo';" &&\
    createdb -O andeo andeo &&\
    psql -f ./scripts/init.sql -d andeo &&\
    /etc/init.d/postgresql stop

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGVER/main/postgresql.conf

# Настройка конейтейра
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql", "./logs"]

WORKDIR /usr/src/app

COPY . .
COPY --from=build /app/main .

EXPOSE 5000

ENV POSTGRES_USER andeo
ENV POSTGRES_DB andeo
ENV POSTGRES_PASSWORD andeo
ENV POSTGRES_HOST localhost
ENV POSTGRES_PORT 5432
ENV POSTGRES_SSLMODE disable


USER root
RUN mkdir -p ./logs/
RUN chmod -R 777 ./logs/
CMD service postgresql start && ./main -config-path=./cmd/main/configs/prod.toml