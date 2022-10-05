package com.honesttruth.uiapi.entity;

import javax.persistence.*;

@Entity
@Table(name = "videos")
public class Video {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;

    @Column(name = "url")
    private String url;

    @Column(name = "title")
    private String title;
}
