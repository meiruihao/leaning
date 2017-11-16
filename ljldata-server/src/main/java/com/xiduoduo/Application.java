package com.xiduoduo;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;

//启动一个服务注册中心
@EnableEurekaClient
@SpringBootApplication
public class Application {

    public static Logger logger = LoggerFactory.getLogger(Application.class);

    public static void main(String[] args) throws Exception {
        //启动Spring Boot项目的唯一入口
        new SpringApplicationBuilder(Application.class).web(true).run(args);
        logger.info("程序启动成功");
    }


}
