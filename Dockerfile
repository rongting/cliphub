FROM scratch

WORKDIR /home/golab/cliphub
COPY . /home/golab/cliphub

EXPOSE 9000
ENTRYPOINT ["./cliphub"]