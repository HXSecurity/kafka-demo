FROM openjdk:8u322-jre
FROM maven:latest
WORKDIR /opt/
COPY . .
RUN mvn clean package -Dmaven.test.skip=true
CMD java -javaagent:agent.jar -Ddongtai.app.name=${dongtai_name} -Ddongtai.log.path=/tmp/dongtai.log -jar target/demo-0.0.1-SNAPSHOT.jar
