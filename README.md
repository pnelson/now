# now

Live in the present. Otherwise, try [when](https://github.com/pnelson/when).

```sh
$ now
Sat Aug 8 11:24 EDT
```

```sh
$ now -f '15:04'
11:24
```

```sh
$ now -s
1596900247
```

```sh
$ now -rfc-3339
2020-08-08T11:24:09-04:00
```

```sh
$ now -u -rfc-3339
2020-08-08T15:24:09Z
```

```sh
$ now -l 'America/Vancouver,America/Toronto,UTC,Europe/Moscow,Asia/Taipei'
Sat Aug 8 08:24 PDT
Sat Aug 8 11:24 EDT
Sat Aug 8 15:24 UTC
Sat Aug 8 18:24 MSK
Sat Aug 8 23:24 CST
```

```sh
$ now -f '15:04' -l 'America/Vancouver,America/Toronto,UTC,Europe/Moscow,Asia/Taipei'
08:24
11:24
15:24
18:24
23:24
```
