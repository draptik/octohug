---
layout: post
title: "Pygments and Python"
date: 2013-05-24 21:11
comments: true
categories: [blog]
---
Frameworks like [Octopress](http://octopress.org/) and [Jekyll](http://jekyllbootstrap.com/) use the Python library [Pygments](http://pygments.org/) for syntax highlighting.

Currenty (2013-05-23) Pygments does not work flawlessly with Python3 
(see 
[this post](http://nonsenseby.me/blog/2013/04/13/arch-linux/) and 
[this issue on pygment's issue tracker](https://github.com/tmm1/pygments.rb/issues/45)).

This is a problem because many operating systems already provide Python3 as the default python installation.

To circumvent this issue, I found 2 different solutions:

1. 	**Rewire your system to always use python version X (i.e. 2.7) when python is called**:
	See [this post](http://nonsenseby.me/blog/2013/04/13/arch-linux/) for an example solution.

2. 	**Sandbox Python installations**.

Personally I prefer the approach of having a sandbox system. This way I don't change the system's default setting. Whenever I require a different Python version I can just switch it (this is similar to the [Ruby Version Manager rvm](https://rvm.io/)).

[This blog post](http://www.wongdev.com/blog/2013/01/16/octopress-on-archlinux/) explains some further details on how to install Octopress in Arch Linux. 
It comes down to installing [python-virtualenvwrapper](https://wiki.archlinux.org/index.php/Python_VirtualEnv#Virtualenvwrapper) and configuring a new custom environment <code>blog_env</code> which uses Python 2.7:

{% gist 5638602 %}

Paraphrazing the [original post](http://www.wongdev.com/blog/2013/01/16/octopress-on-archlinux/):

To switch to the newly created <code>blog_env</code>, run <code>workon blog_env</code>. To exit a virtualenv, run <code>deactivate</code>.

