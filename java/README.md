# Before you start:
- Make sure you have basic understanding about ``` Event-Driven Applications ```, you can start from here https://www.upsolver.com/blog/apache-kafka-event-driven-architecture-streaming
# Get Started: 
1 - Make sure docker is up and running

2 - spin up docker file to setup kafka server and zookeeper by running this command:
```bash
$ docker-compose up -d
```

3 - run the application , and send message to consumer by using postman or from command line : 
```bash
$ curl -X GET \
  'http://localhost:8080/kafka/publish?message=test' 
```

4 - consumer will auto pick up the message because both consumer and producer listening to the same kafka server
```bash
[ntainer#0-0-C-1] com.demo.kafka.demo.consumers.Consumer   : #### -> Consumed message -> test
```



## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://choosealicense.com/licenses/mit/)
