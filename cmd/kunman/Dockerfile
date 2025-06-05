FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY billing-middleware /app

ENTRYPOINT ["/app/billing-middleware"]

EXPOSE 50890 50891 50892 50893

CMD ["run"]
