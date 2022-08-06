FROM golang:1.16

RUN apt-get update
RUN apt-get install -y vim less wget locales && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8

ENV LANG='ja_JP.UTF-8'
ENV LANGUAGE='ja_JP:ja'
ENV LC_ALL='ja_JP.UTF-8'
ENV TZ JST-9
ENV TERM xterm
