package com.xiduoduo.controller;

import com.xiduoduo.exception.MyException;
import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

/**
 * Created by meiruihao on 2017/11/8.
 */
@RestController
@Configuration
@RequestMapping(value="test")
public class InvokerController {

    @Bean
    @LoadBalanced
    public RestTemplate getRestTemplate() {
        return new RestTemplate();
    }

    @RequestMapping(value = "/router")
    public String router() {
        RestTemplate restTpl = getRestTemplate();
        // 根据应用名称调用服务
        String json = restTpl.getForObject(
                "http://ljldata-server/test/add", String.class);
        return json;
    }

    @RequestMapping("/hello")
    public String hello() throws Exception {
        throw new Exception("发生错误");
    }

    @RequestMapping("/json")
    public String json() throws MyException {
        throw new MyException("发生错误2");
    }


}
