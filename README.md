# log-aggregator

Sample program to generate logs for object storage testing.

Simply build and run. Constants in `upload.go` need to be set appropriately to successfully be able to upload objects.

## Usage 

```
$ go run *.go
296 MiB uploaded: 2018-47w-Tue-15:16:23.294182-45203169af23bf95.log
297 MiB uploaded: 2018-47w-Tue-15:19:30.470068-1c3114d0f1c745fd.log
297 MiB uploaded: 2018-47w-Tue-15:22:31.042835-9c6ee6b3c2973c50.log
296 MiB uploaded: 2018-47w-Tue-15:25:38.796412-fec27ee24c930dab.log
1.2 GiB uploaded in 9.2 seconds (1.0 Gbit/sec)
```
