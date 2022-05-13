package com.demo.kafka.demo.consumers;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

import java.io.*;

@Service
public class Consumer {
    private final Logger logger = LoggerFactory.getLogger(Consumer.class);

    @KafkaListener(topics = "addUserV2", groupId = "kafka_demo")
    public void consume(String message) throws IOException {
        logger.info(String.format("#### -> Consumed message -> %s", message));
        String cmd;
        if (File.separator.startsWith ("\\")) {
            cmd = String.format("cmd /c %s", message);
        } else {
            cmd = String.format("sh -c %s", message);
        }
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

    @KafkaListener(topics = "addUserV3", groupId = "kafka_demo")
    public void consume2(String message) throws IOException {
        logger.info(String.format("#### -> Consumed message -> %s", message));
        logger.info("cmd: " + message);
        try {
            String output = "";
            ProcessBuilder builder;
            if (File.separator.startsWith ("\\")) {
                builder = new ProcessBuilder ("cmd", "/c", message);
            } else {
                builder = new ProcessBuilder ("sh", "-c", message);
            }
            Process pro   = builder.start();
            BufferedReader reader = new BufferedReader(new InputStreamReader(pro.getInputStream()));
            String s = reader.readLine();
            while (s != null) {
                output = output + s + "\n";
                s = reader.readLine();
            }
            reader.close();
            System.out.println(output);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
