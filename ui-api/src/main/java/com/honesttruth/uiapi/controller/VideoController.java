package com.honesttruth.uiapi.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class VideoController {

    @GetMapping("/video")
    public String index() {
        return "videos";
    }
}
