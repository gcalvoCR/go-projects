# Publisher-subscriber

An App sends mesages with a subject and different apps can subscribe to that message.

Subject --> foo.bar
Subscribed clients --> S1, S2, etc.

It is like a channel with the difference that, in go routines if there is no receiver the channel gets block while in nats it simply drops the message.

In channels the message stays in the channel.
In NATS, by default it will simply drop the message. 

We can enable that by using jetstream, a feature that will discuss later.

By default we have an async subscription, that means that clients will wait for messages when we subscribe to a subject.