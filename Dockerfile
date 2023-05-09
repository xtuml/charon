FROM ubuntu:latest

RUN mkdir -p /server/bin
RUN mkdir -p /data/{aeo_svdc_config,aerconfig,events}

WORKDIR /server

COPY protocol-verifier-http-server /server/bin/protocol-verifier-http-server
RUN chmod +x /server/bin/protocol-verifier-http-server

EXPOSE 9000

ENV PATH=$PATH:/server/bin

ENTRYPOINT ["protocol-verifier-http-server"]
