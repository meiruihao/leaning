package com.xiduoduo.controller;

import org.apache.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.client.discovery.DiscoveryClient;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

/**
 * Created by meiruihao on 2017/11/1.
 */
@Controller
@RequestMapping(value="test")
public class ComputerController {
    private final Logger logger = Logger.getLogger(getClass());

    @Autowired
    private DiscoveryClient client;

    @RequestMapping(value = "/add")
    @ResponseBody
    public Integer add() {
        System.out.println("111111111111111111");
        return 1;
    }
}
