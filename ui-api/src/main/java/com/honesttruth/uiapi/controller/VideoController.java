package com.honesttruth.uiapi.controller;

import com.honesttruth.uiapi.entity.Video;
import com.honesttruth.uiapi.service.VideoService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class VideoController {

    private final VideoService videoService;

    public VideoController(VideoService videoService) {
        this.videoService = videoService;
    }

    @GetMapping("/video")
    public List<Video> index() {
        // Need to add a VideoDTO to return and map to it -- entities to front end === bad
        return videoService.allVideos();
    }
}
