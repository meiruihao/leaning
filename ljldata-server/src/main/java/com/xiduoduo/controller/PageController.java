package com.xiduoduo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

/**
 * Created by meiruihao on 2017/11/8.
 */
@Controller
@RequestMapping("page")
public class PageController {

    @RequestMapping(value = "/test")
    @ResponseBody
    public String test(){
        return "success";
    }
}
