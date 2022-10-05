package com.honesttruth.uiapi.service;

import com.honesttruth.uiapi.entity.Video;
import com.honesttruth.uiapi.repo.VideoRepository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class VideoService {

    private final VideoRepository videoRepository;

    public VideoService(VideoRepository videoRepository) {
        this.videoRepository = videoRepository;
    }

    public List<Video> allVideos() {
        return videoRepository.findAll();
    }
}
