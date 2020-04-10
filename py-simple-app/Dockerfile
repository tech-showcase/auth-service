FROM python:3.7-alpine AS builder

RUN apk add --no-cache python3-dev libstdc++ && \
    apk add --no-cache g++ && \
    ln -s /usr/include/locale.h /usr/include/xlocale.h

COPY requirements.txt /
RUN pip install -r requirements.txt

COPY src/ /app
WORKDIR /app
CMD ["python", "-u", "main.py"]
