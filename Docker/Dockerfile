FROM ubuntu

ENV C_ID=cb9173b

WORKDIR /app
RUN apt update && apt install wget jq -y
RUN wget https://github.com/someone-stole-my-name/cfdns/releases/download/${C_ID}/cfdns_linux_amd64_master
RUN chmod 755 cfdns_linux_amd64_master
