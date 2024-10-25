# Key Value

In NAts the key is a string and the Value is an arrayf bytes, that means basically anything.

We just need a bucket to store the data.

```sh
# to list buckets
nats kv ls

# to add a kv
nats kv add music

# to remove a kv
nats kv rm music

# to add or update data into a bucket
nats put music 1 'Three little birds'

# to delete a key
nats kv del music 1
```