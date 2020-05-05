FROM python:3.8.2

WORKDIR /usr/app

RUN \
   python -m pip install grpcio \
   && python -m pip install grpcio-tools \
   && curl https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz --output go1.14.2.linux-amd64.tar.gz \
   && tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz \
   && rm -f go1.14.2.linux-amd64.tar.gz \
   && export PATH=$PATH:/usr/local/go/bin \
   && export GO111MODULE=on \
   && go get github.com/golang/protobuf/protoc-gen-go


#export GO111MODULE=on
#go get google.golang.org/grpc@v1.28.1

#go get github.com/golang/protobuf



COPY 'libs/protoc-3.11.4-linux-x86_64/bin' ./bin
COPY libs/protoc-plugins/grpc_python_plugin ./bin
COPY sh ./sh/


#CMD bash sh/generate-all.sh
