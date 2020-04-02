# csvdiff

diff text files and print the adds

```sh
$ ./diffcsv testdata/a.csv testdata/b.csv
4
5

$ time ./csvdiff testdata/a-full.csv testdata/b-full.csv > testdata/c-full-multi.csv
```

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
about 15min
```
