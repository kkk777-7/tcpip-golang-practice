FROM golang:latest

RUN apt update -y && \
	apt install -y iptables tcpdump ethtool

CMD ["/bin/bash"]