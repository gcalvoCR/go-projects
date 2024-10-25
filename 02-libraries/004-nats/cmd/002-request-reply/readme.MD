# Queue Groups

In the request reply pattern we can have an issue.

Multiple clients can acknowledge the same message. 

The solutions are queues.

If you have multiple subscribers, both are receiving the message.

The solution is in the subscriber, the publisher should not need to know where to send the message.