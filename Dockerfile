FROM alpine:latest

RUN apk update
RUN apk upgrade
RUN apk add gcompat

EXPOSE 8080

ENV WORK_PATH=/usr/poly_day

COPY ./docker $WORK_PATH/
COPY ./.bin/developer_service/poly_day_microservice_developer_service $WORK_PATH/

WORKDIR $WORK_PATH

RUN chmod 777 ./entrypoint.sh
RUN chmod 777 ./export_variables.sh
RUN chmod 777 ./poly_day_microservice_developer_service

ENTRYPOINT ./entrypoint.sh