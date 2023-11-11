FROM python:slim
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources


RUN apt-get update
RUN apt-get install -y socat

RUN apt install -y python3-pip

RUN pip install pycryptodome --break-system-packages

RUN mkdir /app

WORKDIR /app
COPY task.py .

CMD ["/bin/sh", "-c", "socat TCP-LISTEN:12345,fork,reuseaddr EXEC:'python3 task.py',stderr"]