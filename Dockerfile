FROM ubuntu:latest

RUN mkdir -p /server/{bin,templates}
RUN mkdir -p /data/{aeo_svdc_config/job_definitions,aerconfig,events,logs/verifier,logs/reception,verifier_processed,verifier_incoming,invariant_store,job_id_store}

WORKDIR /server

COPY protocol-verifier-http-server /server/bin/protocol-verifier-http-server
RUN chmod +x /server/bin/protocol-verifier-http-server
COPY /templates/index.html /server/templates/index.html

EXPOSE 9000

ENV PATH=$PATH:/server/bin

ENTRYPOINT ["protocol-verifier-http-server"]