# csvdiff

diff csv files and find out the additions.

It can be used to measure CPU single core, multi-core performance.


## Build
Go 1.12+


```sh
$ git clone https://github.com/zyfdegh/csvdiff
$ go build -mod=vendor
```

## Run
```sh
# find and print lines in b.csv but not in a.csv
$ ./csvdiff testdata/a.csv testdata/b.csv
4
5

# a-full.csv 176543 lines
# b-full.csv 340012 lines
$ time ./csvdiff testdata/a-full.csv testdata/b-full.csv > testdata/c-full-multi.csv

# PowerShell (Windows)
> Measure-Command {.\csvdiff.exe testdata\a-full.csv testdata\b-full.csv > .\testdata\c-full.csv}
```

## Result

```sh
1. diff(), Intel i5-7500
183.91s user 2.06s system 93% cpu 3:19.30 total

2. diffmulti(), Intel i5-7500, 68.8% less time compared to [1]
173.17s user 1.60s system 280% cpu 1:02.29 total

3. diffmulti()，Intel i5-2520M
537.72s user 1.12s system 393% cpu 2:16.95 total

4. diffmulti()，Intel i5-7267U
348.91s user 1.44s system 362% cpu 1:36.56 total

5. diffmulti(), Intel Atom N270
~ 15min

6. diffmulti(), AMD R5 3600
~ 27s
```
