# JetStream

NATS has a built-in distributed persistence system called Jetstream which enables new funcionalities and higher qualities of service on top of the base 'Core NATS' functionalities and qualities of service.

Persistency and distributy

What does JetStream enables?

Streaming
- Temporal decoupling between publishers and subscribers.
(Before jetstream if a publisher is sending messages, a subscriber needed to be online at the same time to get those messages)
- Other systems typically solved the issue with `durable consumers` or `queues`
- NATS solves this issues with `Stream`

Key/Value Store
- Similar to Redis

Document Store
- Similar to S3

## What is a Stream?
- A stream captures and stores messages published on one or more subject(s)
- NATS client can consume messages from the Stream, at any time, by creating 'subscribers'

on the publisher nothing changes, we now tell NATS to store them in a stream, NATS persists the messages.

When a client is subscribed to a stream we call it a `JetStream Consumer` or just a `Consumer`.

A `consumer` can replay or consume messages in the Stream.
- Consummers can replay all or some of the messages stored in the Stream.

## Feature of a stream
- Replay policies
  - How and when messages are delivered to Consumers
- Retention Policies
  - How long to keep a message and rules for discarding
  - Policies
    - Limits --> age, number of messages, stream size, etc.
    - Interest --> keep messages tht haven't been consumed yet
    - Work queue --> Stream is used as a shared queue, and messages are removed as they are consumed.
- Store
- ...

Note: A stream can store messages from multiple subjects.


### Consumer replay policies
-`all` It can ask for `all` of the messages currently in the Stream and NATS can try pushing:
  - `instant` deliver as fast as consumer can handle
  - `original` deliver at the same rate the messages were published.

- `last` message stored in the Stream or for all Subjects in the Stream.
- from a `sequence number` (give me from message 1000 forward, or message 100 only)
- from a specific `time`. 


## running example
```sh
# -js enables jetstream
docker run -p 4222:4222 nats:latest -js
# we create a stream
nats stream add ORDERS --subjects 'orders.*'
# publish a message
nats pub "orders.us" --count=-1 --sleep=2s  "US Order #  {{Count}} at {{TimeStamp}}"
# run the go application
...
## clean up
nats consumers ls ORDERS
nats consumer rm ORDERS my-consumer-1

nats stream ls
nats stream rm ORDERS

```
nats