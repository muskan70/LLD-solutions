# Pub/Sub Queues

### Implement a persistent pub-sub queue mechanism with guaranteed delivery of every published message for all the subscribed consumers in that subscribed topic in the same order.

Topic
Topic A -> “{ a: b, c :d }”, “{ e: f, g :h }” 
Topic B -> “{ a: b, c :d }”


Subscriber A ->(subscribed to) Topic A
Subscriber B ->(subscribed to) Topic B
Subscriber C ->(subscribed to) Topic A, Topic B

Subscriber A -> 
Reading from topic A -> “{ a: b, c :d }”
Reading from topic B ->  

Subscriber C -> 
Reading from topic A -> “{ a: b, c :d }”
Reading from topic B -> “{ a: b, c :d }” 
Reading from topic A -> “{ e: f, g :h }”



# Features:
1. API to expose topics
2. API for publisher to push messages against a topic
3. API to subscribe and unsubscribe from topic
4. API for subscriber to consume from topic
5. Subcriber can again subscribe a topic and start consuming messages from where it left
6. API to check status of a message in a topic if consumed by subscribers or not
