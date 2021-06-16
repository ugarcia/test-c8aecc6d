# test-c8aecc6d

Command line tool to get md5 hashes from url responses in the format:
```bash
url1 hash1
url2 hash2
.....
```

### Usage

```bash
go build
./test-c8aecc6d [-parallel N] [url1] [url2] ...
```

Being `-parallel N` the number of parallel requests at any time
And `urlX` a url to request

### Test
```bash
go test ./...
```

### Help
```bash
go build
./test-c8aecc6d -h
```
