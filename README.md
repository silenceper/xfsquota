# xfsquota
xfsquota is a tool for managing XFS quotas

support set quota and get quota

# build
```shell
make build
```
##  Usage
```shell
xfsquota is a tool for managing XFS quotas

Usage:
  xfsquota [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Get quota information
  help        Help about any command
  set         Set quota information

Flags:
  -b, --basePath string   base path for backing filesystem device (default "./")
      --debug             debug mode
  -h, --help              help for xfsquota
```

## Set Quota
set quota size 1MiB ,inodes 20 for path `/data/test/quota`
```shell
> xfsquota -b /data set /data/test/quota  -s 1MiB -i 20

set quota success, path: /data/test/quota, size:1MiB, inodes:20
```

## Get Quota
get quota for path `/data/test/quota`
```shell
> xfsquota -b /data get /data/test/quota

quota Size(bytes): 1048576
quota Inodes: 20
diskUsage Size(bytes): 0
diskUsage Inodes: 1
```

