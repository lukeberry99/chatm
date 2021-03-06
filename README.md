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
    4.2. requires IE >= 9, android >= 4.4
5. Stateless

## Notes
* 1 go routine listening for incoming
* 1 go routine replying

## Potential pitfalls
1. Redis > 1000 nodes has some concurrency / parallelism issues
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
