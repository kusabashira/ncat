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

Other Specification
-------------------

`N` is a number of times to repeat.

If `N` is a negative number, it interpreted as an infinity.

```
$ echo Hello | catn -2
Hello
Hello
Hello
Hello
Hello
Hello
Hello
Hello
Hello
...
```

License
-------

MIT License

Author
------

nil2 <nil2@nil2.org>
