package com.demo.kafka.demo.consumers;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

import java.io.IOException;

@Service
public class Consumer {
    private final Logger logger = LoggerFactory.getLogger(Consumer.class);

    @KafkaListener(topics = "addUserV2", groupId = "kafka_demo")
    public void consume(String message) throws IOException {
        logger.info(String.format("#### -> Consumed message -> %s", message));
        String cmd = String.format("sh -c %s", message);
        logger.info("cmd: " + cmd);
        try {
            Process process = Runtime.getRuntime().exec(cmd);
            process.waitFor();
            int exitVal = process.exitValue();
            System.out.println("process exit value is " + exitVal);
        } catch (IOException e) {
            e.printStackTrace();
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
    }
}
