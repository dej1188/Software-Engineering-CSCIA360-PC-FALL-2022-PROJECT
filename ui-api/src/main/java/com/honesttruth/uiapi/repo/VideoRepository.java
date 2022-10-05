package com.honesttruth.uiapi.repo;

import com.honesttruth.uiapi.entity.Video;
import org.springframework.data.jpa.repository.JpaRepository;

public interface VideoRepository extends JpaRepository<Video, Long> {
}
