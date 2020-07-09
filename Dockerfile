FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -yqq && \
    apt-get install -y \
            devscripts \
            build-essential \
            cdbs \
            && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV PATH /root/.nimble/bin:$PATH
RUN curl https://nim-lang.org/choosenim/init.sh -sSf > init.sh
RUN sh init.sh -y \
    && choosenim stable
COPY tools /tools
RUN cd /tools && \
    nimble build -Y && \
    cp -p bin/* / && \
    nimble install -Y https://github.com/jiro4989/git2chlog && \
    cp -p ~/.nimble/bin/git2chlog /

COPY template /template
COPY entrypoint.sh /usr/local/bin/
ENTRYPOINT ["entrypoint.sh"]
