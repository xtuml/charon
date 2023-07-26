# first stage to build the go binary
FROM golang:bullseye as firstStage
RUN apt-get update && apt-get -y install make 
RUN mkdir /tmp/http_server
WORKDIR /tmp/http_server
COPY . /tmp/http_server/
RUN make dep && make build

# second stage to build the app 
FROM ubuntu:latest

RUN mkdir -p /server/{bin,templates}
RUN mkdir -p /data/{aeo_svdc_config/job_definitions,aerconfig,events,logs/verifier,logs/reception,verifier_processed,verifier_incoming,invariant_store,job_id_store}

COPY --from=firstStage /tmp/http_server/protocol-verifier-http-server /server/bin/protocol-verifier-http-server
WORKDIR /server
RUN chmod +x /server/bin/protocol-verifier-http-server
COPY /templates/index.html /server/templates/index.html

EXPOSE 9000

ENV PATH=$PATH:/server/bin

ENTRYPOINT ["protocol-verifier-http-server"]