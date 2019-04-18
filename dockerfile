FROM scratch

MAINTAINER  "hcf"

WORKDIR .
ADD main .
ADD test.toml .

EXPOSE 8082
CMD ["./main","-config=./test.toml"]
