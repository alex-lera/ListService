FROM ubuntu
ENV TZ=Europe/Madrid
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update
RUN apt-get -y install curl
RUN apt-get -y install wget
RUN apt-get -y install git
RUN apt-get -y install golang
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "gopkg.in/yaml.v2"
RUN go get "os"
COPY config.yaml /opt/config.yaml
COPY getcar /opt/getcar
EXPOSE 8081
CMD ["/opt/getcar"]