#container image
FROM amazoncorretto:17

#container working directory
WORKDIR /app

ENV PORT 8080

#expose our port to run application in container
EXPOSE 8080
#copy our build to container working directory
COPY target/*.jar /app/spring-boot-app.jar



#run the application
ENTRYPOINT exec java $JAVA_OPTS -jar spring-boot-app.jar