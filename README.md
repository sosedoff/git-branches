# git-branches

A small tool to show current status of the git repository

## Install

```bash
go install github.com/sosedoff/git-branches@latest
```

## Example

```bash
$ git branches
```

Output:

```
+-----------------------------+--------+-------+-------------------------------+--------+
|            NAME             | BEHIND | AHEAD |          LAST COMMIT          | STATUS |
+-----------------------------+--------+-------+-------------------------------+--------+
| master                      |      1 |     1 | 2020-02-19 20:30:36 -0600 CST | ACTIVE |
| remote-double-click-on-cell |      5 |     1 | 2020-02-12 13:14:58 -0600 CST | ACTIVE |
| artifact-upload-action      |      8 |    12 | 2020-02-11 22:16:20 -0600 CST | ACTIVE |
| branch-binary-builds        |      8 |     3 | 2020-02-11 17:51:01 -0600 CST | ACTIVE |
| ssl-root-cert               |      9 |     1 | 2020-02-11 10:58:40 -0600 CST | ACTIVE |
| database-export-err-logging |     12 |     1 | 2020-01-02 15:35:06 -0600 CST | STALE  |
| remote-dbl-click            |     12 |     1 | 2019-12-26 15:54:54 -0600 CST | STALE  |
| deps-update-20191210        |     15 |     1 | 2019-12-10 20:33:01 -0600 CST | STALE  |
| oleggator-autocompletion    |     32 |     1 | 2019-11-02 13:22:31 -0500 CDT | DEAD   |
| fmt                         |     37 |     5 | 2019-09-29 09:54:03 -0500 CDT | DEAD   |
| discovery                   |     37 |    12 | 2019-09-25 20:54:10 -0500 CDT | DEAD   |
+-----------------------------+--------+-------+-------------------------------+--------+
```

Filter branches by name:

```
$ git branches foo
```

## License

MIT
