FROM python:3.10.4-alpine3.15 as dev

WORKDIR /work


FROM python:3.10.4-alpine3.15 as debugging

# add a debuggger

FROM dev as runtime

COPY . /app

ENTRYPOINT ["python","/app/intro/main1.py"]
