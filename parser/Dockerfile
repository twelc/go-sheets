FROM python:3.11.6

WORKDIR /usr/src/app
COPY . .

RUN pip install -r requirements.txt

RUN apt update -y && apt upgrade -y
RUN apt install tor obfs4proxy -y
RUN echo "SocksPort 9150" >> /etc/tor/torrc
RUN echo "ClientTransportPlugin obfs4 exec /usr/bin/obfs4proxy" >> /etc/tor/torrc
RUN echo "Bridge obfs4 5.75.242.46:62511 4813B8925EE2124B0AED15BAB7460EA83F5E9EC9 cert=isijAs/iLPpqTbCwxUhyXvCYpLqynHd09WoruIUf1sQ0MYG1oRdC1ejg5M7OGRCwRJiOZQ iat-mode=0" >> /etc/tor/torrc
RUN echo "UseBridges 1" >> /etc/tor/torrc
