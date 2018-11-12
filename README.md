# ChatMi
###### Highly available chat server built in Go

## Key requirements

1. Chat app
    1.1. Uses redis for message exchange
    1.2. Horizontally scalable
2. Docker
3. Deployed to ECS / Fargate
4. WebSockets
    4.1. github.com/gorilla/websocket
6. Stateless

## Potential pitfalls
1. Redis > 1000 nodes has some pitfalls
	1.1. unlikely we will reach this 1000 node issue, switching to RabbitMQ would stop this and potentially 
    
<details>
<summary>Example architecture</summary>
<p>
```
                    ┌──────────────────────┐                  
                    │    load balancer     │                  
 ecs                └──────────────────────┘                  
┌────────────────────────────────────────────────────────────┐
│                                                            │
│            cluster1                     cluster2           │
│   ┌─────────────────────────┐  ┌─────────────────────────┐ │
│   │ ┌─────────┐┌─────────┐  │  │ ┌─────────┐┌─────────┐  │ │
│   │ │   app   ││   app   │  │  │ │   app   ││   app   │  │ │
│   │ └─────────┘└─────────┘  │  │ └─────────┘└─────────┘  │ │
│   │       ┌─────────┐       │  │       ┌─────────┐       │ │
│   │       │   app   │       │  │       │   app   │       │ │
│   │       └─────────┘       │  │       └─────────┘       │ │
│   └─────────────────────────┘  └─────────────────────────┘ │
│                                                            │
│                        redis cluster                       │
│                 ┌───────────────────────────┐              │
│                 │  ┌─────────┐ ┌─────────┐  │              │
│                 │  │  redis  │ │  redis  │  │              │
│                 │  └─────────┘ └─────────┘  │              │
│                 │        ┌─────────┐        │              │
│                 │        │  redis  │        │              │
│                 │        └─────────┘        │              │
│                 └───────────────────────────┘              │
└────────────────────────────────────────────────────────────┘
```
</p>
</details>
