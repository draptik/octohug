---
layout: post
title: "Visual Studio's default path for new projects"
date: 2018-02-20 18:13:10 +0000
comments: true
categories: [ VisualStudio ]
---
Today I took the time to fix something very simple: Visual Studio's default path.

In the past decade there has never been a single project I wanted to save to:

`C:\Users\<username>\Documents\Visual Studio <vs version>\Projects`

{% img /images/posts/vs-default-path/vs-default-path1.png %}

I was expecting having to dive into the windows registry to fix this, but it turned out to be rather simple to change this.

In Visual Studio navigate to `Tools -> Options -> Projects and Solutions -> Locations`

{% img /images/posts/vs-default-path/vs-default-path2.png %}

Change `Project Location` to the desired folder:

{% img /images/posts/vs-default-path/vs-default-path3.png %}

Et voila:

{% img /images/posts/vs-default-path/vs-default-path4.png %}

TL/DR: In Visual Studio: `Tools -> Options -> Projects and Solutions -> Locations`

