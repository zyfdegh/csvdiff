# csvdiff

diff 2 csv files and find out the additions of the latter.

It can be used to measure CPU single core, multi-core performance.


## Build
Go 1.9+


```sh
$ git clone https://github.com/zyfdegh/csvdiff
$ go build -mod=vendor
```

## Run
To figure out what will happen
```sh
# find and print lines in b.csv but not in a.csv
$ ./csvdiff testdata/a.csv testdata/b.csv
4
5
```

To use all the CPU cores by default
```sh
# a-full.csv 176543 lines
# b-full.csv 340012 lines
$ time ./csvdiff testdata/a-full.csv testdata/b-full.csv > testdata/c-full-multi.csv
```

To use single core, add `1` at last param (or any core you want)
```sh
$ time ./csvdiff testdata/a-full.csv testdata/b-full.csv > testdata/c-full-single.csv 1
```

For Windows
```sh
# PowerShell (Windows)
> Measure-Command {.\csvdiff.exe testdata\a-full.csv testdata\b-full.csv > .\testdata\c-full.csv}
```

## Result

| CPU | Cores, Threads | Base Freq | OS | Single core | Multicores |
|--|:--:|--|--|--|--|
| AMD R5 3600 | 6,12 | 3.6GHz | Windows 10 | | 27.2s |
| Intel i5-7500  | 4, 4 | 3.4GHz | macOS 10.14 | 3m19.5s | 1m2.29s |
| Intel i5-7267U | 2, 4 | 3.1GHz | macOS 10.14 | - | 1m36.56s |
| Intel i5-2520M | 2, 4 | 2.5GHz | Linux 5.5   | - | 2m16.95s |
| Intel Atom N270 | 1, 2 | 1.6GHz | Linux 5.5 | ~15m | - |

