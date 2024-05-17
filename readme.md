
# Introduction
Application for massive logging user activities

## How to
Majoo repository **svc-activity** currently has several commands, to run those application simply run these commands below.
- Consumer/Listener (Kafka Broker)
```  
go run main.go consumer  
```  
- List item
```  
go run main.go http  
```  

## Create new commands

To create a new commands for application such as

- Cron / Scheduler
- Socket server
- gRPC
- Streaming
- Web RTC
- Etc.

Go to Folder named **cmd**

1. Add new file for new command such as **socket.go**
2. Register command that you created to **cmd.go**

## Folder Structure
This repository is using hexagonal architecture pattern, hexagonal architecture has an **flexibility** to adapt a new port or even changing port with minimum changes

- **`cmd`**: List of registered commands or entrypoints
- **`config`**: List of configurations and dependencies
- **`utils`**: List of **GLOBAL** helpers function
- **`docs`**: Stored documentations (add more docs in markdown format if necessary) and swagger
- **`internal`**: Source folder
    - **`adapter`**: List of external stacks of application (databases, cache, brokers, etc)
        - **`handler`**
        - **`libraries`**
        - **`repository`**
    - **`core`**: List of business logic here
        - **`domain`**: List of struct input and output for port
            - **`presenters`**: Request and response from handler
            - **`entities`**: Input and output from service
            - **`models`**: Input and output from repository
        - **`port`**: These are list of adapter interfaces to obey for interacting between services and adapter and for **Unit Test** mock purposes
            - **`handler`**
            - **`libraries`**
            - **`repository`**
        - **`services`**: These are list of business logic


## Structure Preview
This preview is **ONLY**  example, between `adapter` and `core` required port for flexibility:

![hexa architecture](docs/images/hexa.png)
![hexa diagram](docs/images/hexa-uml.png)

preview using: https://mermaid.js.org/intro/
```mermaid  
graph LR  
A(Client) -- payload --> B[Adapter - Handler] --> C{Port} --> D[Services]
D[Services] --> E{Port}
E{Port} --> F[Repository]
F[Repository] --> G1[Mongo]
F[Repository] --> G2[Postgres]
F[Repository] --> G3[Redis]
E{Port} --> F1[Library]
F1[Library] --> F2[Kafka]
F1[Library] --> F3[RabbitMQ]
```

## Swagger

This repository is using swagger for http documentation
|No|Link|
|--|--|
| echo swagger | https://github.com/swaggo/echo-swagger/tree/master |
| go swag |.https://github.com/swaggo/swag#declarative-comments-format |

> Install first swag command on go swag

How to add new documentation on handler

1. Add new declarative comment on handler
2. on root folder run this `swag init -g cmd/http.go`
3. Format comments `swag fm`
4. Those commands will update file on `docs/docs.go` `docs/swagger.json` `docs/swagger.yaml`
5. Run http application `go run main.go http`
6. Go to `base-url/docs/`, now see the update