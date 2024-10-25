# NATS

Connectivity Technology

## The problem
- Communication between microservices could get really complex
- Apps should be resilient against failure
- Apps should scale seamlessly
- New services should be added7removed without disruption

## Desired Solution
- Have a way to connect many decoupled application services
- Easy to implement
- Securely
- Provide flexibility in messaging patterns
- The solution should scale

## Why NATS
- We introduce a Messaging fabric or a bus to solve solution

### NATS features
- Location Independent Addressing
- Pub/sub 
- Request/reply
- Streams
- Persistence
- At most once/at least once/ exactly once
- Secure by default
- Multi Tenancy
- Global scale
- Small bincary
- Many client librarues
- High speed
- Durable subscribers
- Encryption at rest
- Subject limits
- Scalable clustering
- Edge & IoT
- Easy configuration
- OSS

## How to run NATS
Install it via:
- Binary
```sh
brew install nats-server
nats-server
```
- Docker
```sh
docker run -d --rm -p 4222:4222 nats:latest
```
- K8s
[https://nats.io/](https://nats.io/)

## CLI commands

```sh
# install it
brew install nats-io/nats-tools/nats

# run a publisher
nats pub test --count 10 "Message {{Count}}: {{ Random 10 100 }}"

# regular subscription
nats sub "test"

# queue subscription
nats sub "test" --queue=QUEUE

# request reply
nats request subject "is there somebody?" --count=100

nats reply subject "got your request"
```

## Subjects

**What is a subject?**
"At its simplest", a subject is a string of characters that form a name which the publisher and subscriber can user to find each other. It helps scope messages into streams or topics.

Subject Names
- a to z, A to Z and 0 to 9 (names are case sensitive, and cannot contain whitespace).
- ., >, and * (always used for subscribers, never on publishers)
  - (.) character is used to create subject hierarchy. It creates tokens
  - (*) to match a single token --> `weather.*.east` anything that is in the east.
  - (>) to match 1 or more tokens --> `weather.>` (the > should always comes last)
- Reserved subject names:
  - By convention subject names starting with $ are reserved for system use
