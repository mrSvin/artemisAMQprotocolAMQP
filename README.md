<h1 align="center">Golang group listeners and one sender for Artemis and protocol AMQP</h1>

## Description
This program implements 4 listeners that synchronously process messages from one sender. If one listener disconnects, other listeners exchange its messages among themselves.

## Deploy Artemis in Docker
```
docker run -d --name artemis -p 8161:8161 -p 61616:61616 vromero/activemq-artemis
```

## Install library github.com/Azure/go-amqp@v0.15.0
```
go get github.com/Azure/go-amqp@v0.15.0  
```

## Execution result
![image](https://user-images.githubusercontent.com/59051004/219303145-be89bfcc-8d14-4598-bfcd-062c27a73644.png)
