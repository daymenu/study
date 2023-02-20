FROM openjdk:8-jre-slim

RUN mkdir /kafka

ENV KAFKA_DATA_DIR=/kafka/data

RUN groupadd -r kafka --gid=1000;\
    useradd -r -g kafka --uid=1000 kafka;\
    mkdir -p "${KAFKA_DATA_DIR}";\
    chown kafka:kafka "${KAFKA_DATA_DIR}"

# RUN set -eux; \
#     apt-get update; \
#     DEBIAN_FRONTEND=noninteractive \
#     apt-get install -y --no-install-recommends \
#         gnupg; \
#     rm -rf /var/lib/apt/lists/*


ADD  ./app/kafka_2.11-1.0.0.tgz /

RUN mv /kafka_2.11-1.0.0 /kafka
WORKDIR /kafka

CMD [ "" ]