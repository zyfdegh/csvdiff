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
# it takes about 60 billion computes
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

## Benchmark

| CPU | Cores, Threads | Base Freq | TDP | OS | Single core | Multicores |
|:--:|:--:|--|--|--|--|--|
| Intel 10700k (Turbo) | 8, 16 | 3.6GHz | 125W | macOS 11.1 | - | 14.573s |
| Intel 10700k | 8, 16 | 3.6GHz | 95W[^1] |macOS 11.1 | - | 18s |
| Apple M1 | 8, 8 | 3.2GHz | 20-24W[^2] | macOS 11 | 1m47.39s | 21.388s |
| AMD R5 3600 | 6, 12 | 3.6GHz | 65W | Windows 10 | - | 27.2s |
| Intel i5-7500  | 4, 4 | 3.4GHz | 65W | macOS 10.14 | 3m01.35s | 48.856s |
| Intel i5-7267U | 2, 4 | 3.1GHz | 28W | macOS 10.14 | - | 1m36.56s |
| Intel i5-2520M | 2, 4 | 2.5GHz |  35W | Linux 5.5   | 4m41.82s | 2m16.95s |
| Intel Atom N270 | 1, 2 | 1.6GHz | 2.5W | Linux 5.5 | - | >15m |

[^1]: https://ark.intel.com/content/www/us/en/ark/products/199335/intel-core-i7-10700k-processor-16m-cache-up-to-5-10-ghz.html
[^2]: https://www.anandtech.com/show/16252/mac-mini-apple-m1-tested
