gotrace
========
A syscall trace app in go. You can think of it as a very basic "strace" implementation in go.

![License](http://img.shields.io/badge/license-MIT-blue.svg)

Supported Platforms
===================

See [libtrace](https://github.com/jfrabaute/libtrace#supported-platforms)


Installation
============

```
go get github.com/jfrabaute/gotrace
```

Usage
============

```
gotrace <app_to_run> [args...]
```

ex:

```
gotrace sleep 2
```

gotrace is using the [libtrace](https://github.com/jfrabaute/libtrace) library.

Licensing
=========
gotrace is licensed under the MIT License. See LICENSE for full license text.
