catn
====

Write each FILE to standard output N times.

```
$ seq 3 | catn 2
1
2
3
1
2
3
```

Usage
-----

```
$ catn N [FILE]...
```

`N` is a number of times to repeat.

If `N` is a negative number, it interpreted as an infinity.

License
-------

MIT License

Author
------

kusabashira <kusabashira227@gmail.com>
