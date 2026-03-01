# How to use

```bash
go build -o links links.go
ll
-rwxrwxr-x  1 go go 8793703 мар  1 12:26 links*

go build -ldflags="-s -w" -o links1 links.go
ll
-rwxrwxr-x  1 go go 5984548 мар  1 12:31 links*
 
./links [--threads=number of threads] [--timeout=timeout in seconds] ...[links to files list]
```

---

```bash
./links --help
Usage of ./links:
  -threads int
        number of threads (default 2)
  -timeout int
        general timeout in seconds (default 120)
```

---

## 1 thread / 1s timeout / 3 files -> error

```bash
./links --threads=1 --timeout=1 https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg https://upload.wikimedia.org/wikipedia/commons/7/74/A-Cat.jpg https://onlinejpgtools.com/images/examples-onlinejpgtools/mountains-near-water.jpg
Starting downloading 3 links in 1 threads
Started: https://onlinejpgtools.com/images/examples-onlinejpgtools/mountains-near-water.jpg
Completed: https://onlinejpgtools.com/images/examples-onlinejpgtools/mountains-near-water.jpg -> mountains-near-water.jpg
Started: https://upload.wikimedia.org/wikipedia/commons/7/74/A-Cat.jpg

Error: error while download https://upload.wikimedia.org/wikipedia/commons/7/74/A-Cat.jpg: can't save file: context deadline exceeded
```

## default settings / non-existent file -> error

```bash
./links https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg https://upload.wikimedia.org/wikipedia/commons/7/74/A-Cat.jpg https://onlinejpgtools.com/images/examples-onlinejpgtools/mountains-near-water.jpg?size=500x500 https://upload.wikimedia.org/wikipedia/commons/404.zip
Starting downloading 4 links in 2 threads
Started: https://upload.wikimedia.org/wikipedia/commons/404.zip
Started: https://upload.wikimedia.org/wikipedia/commons/7/74/A-Cat.jpg

Error: error while download https://upload.wikimedia.org/wikipedia/commons/404.zip: bad status: 404 Not Found
```

## 4 thread / no timeout / 10+ files -> error

It gets error without `SkipVerify`
```go
    tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}
```

```bash
./links --threads=4 https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg https://upload.wikimedia.org/wikipedia/commons/7/74/A-Cat.jpg https://onlinejpgtools.com/images/examples-onlinejpgtools/mountains-near-water.jpg?size=500x500 https://angrytribe.org/old/downloaded_files/tild6163-6465-4561-b864-356261626532/DSC05897.JPG https://angrytribe.org/old/downloaded_files/tild6134-3764-4734-b232-306134393064/30565402375_cfb00f30.jpg https://angrytribe.org/old/downloaded_files/tild6436-3535-4630-b438-356362646631/30180832214_298f969c.jpg https://angrytribe.org/old/downloaded_files/tild6636-6231-4433-b965-623330373034/DJI_0057.JPG https://angrytribe.org/old/downloaded_files/tild6435-3561-4566-b063-383461363533/on5tJvk40kQ.jpg https://angrytribe.org/old/downloaded_files/tild3930-3736-4361-a639-346466386334/2018-09-19_114015_1.jpg https://angrytribe.org/old/downloaded_files/tild3432-3462-4864-b865-346236653435/P9231020.JPG https://angrytribe.org/old/downloaded_files/tild6566-3563-4764-b262-376461366539/DSC_4904.jpg https://angrytribe.org/old/downloaded_files/tild3534-6633-4464-b435-326638386334/P9220984.JPG https://angrytribe.org/old/downloaded_files/tild6630-3435-4632-a232-333962303735/2016-11-20_061854_2.jpg
Starting downloading 9 links in 4 threads
Started: https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg
Started: https://angrytribe.org/old/downloaded_files/tild6630-3435-4632-a232-333962303735/2016-11-20_061854_2.jpg
Started: https://angrytribe.org/old/downloaded_files/tild6435-3561-4566-b063-383461363533/on5tJvk40kQ.jpg
Started: https://angrytribe.org/old/downloaded_files/tild3930-3736-4361-a639-346466386334/2018-09-19_114015_1.jpg

Error: error while download https://angrytribe.org/old/downloaded_files/tild6435-3561-4566-b063-383461363533/on5tJvk40kQ.jpg: download failed: Get "https://angrytribe.org/old/downloaded_files/tild6435-3561-4566-b063-383461363533/on5tJvk40kQ.jpg": tls: failed to verify certificate: x509: certificate has expired or is not yet valid: current time 2026-03-01T12:16:14+03:00 is after 2026-01-17T09:12:20Z
```

## 4 thread / no timeout / 10+ files -> ok

```bash
./links --threads=4 https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg https://upload.wikimedia.org/wikipedia/commons/7/74/A-Cat.jpg https://onlinejpgtools.com/images/examples-onlinejpgtools/mountains-near-water.jpg?size=500x500 https://angrytribe.org/old/downloaded_files/tild6163-6465-4561-b864-356261626532/DSC05897.JPG https://angrytribe.org/old/downloaded_files/tild6134-3764-4734-b232-306134393064/30565402375_cfb00f30.jpg https://angrytribe.org/old/downloaded_files/tild6436-3535-4630-b438-356362646631/30180832214_298f969c.jpg https://angrytribe.org/old/downloaded_files/tild6636-6231-4433-b965-623330373034/DJI_0057.JPG https://angrytribe.org/old/downloaded_files/tild6435-3561-4566-b063-383461363533/on5tJvk40kQ.jpg https://angrytribe.org/old/downloaded_files/tild3930-3736-4361-a639-346466386334/2018-09-19_114015_1.jpg https://angrytribe.org/old/downloaded_files/tild3432-3462-4864-b865-346236653435/P9231020.JPG https://angrytribe.org/old/downloaded_files/tild6566-3563-4764-b262-376461366539/DSC_4904.jpg https://angrytribe.org/old/downloaded_files/tild3534-6633-4464-b435-326638386334/P9220984.JPG https://angrytribe.org/old/downloaded_files/tild6630-3435-4632-a232-333962303735/2016-11-20_061854_2.jpg
Starting downloading 13 links in 4 threads
Started: https://angrytribe.org/old/downloaded_files/tild6163-6465-4561-b864-356261626532/DSC05897.JPG
Started: https://angrytribe.org/old/downloaded_files/tild3534-6633-4464-b435-326638386334/P9220984.JPG
Started: https://angrytribe.org/old/downloaded_files/tild6134-3764-4734-b232-306134393064/30565402375_cfb00f30.jpg
Started: https://angrytribe.org/old/downloaded_files/tild6636-6231-4433-b965-623330373034/DJI_0057.JPG
Completed: https://angrytribe.org/old/downloaded_files/tild6163-6465-4561-b864-356261626532/DSC05897.JPG -> DSC05897_1.JPG
Started: https://angrytribe.org/old/downloaded_files/tild6435-3561-4566-b063-383461363533/on5tJvk40kQ.jpg
Completed: https://angrytribe.org/old/downloaded_files/tild6636-6231-4433-b965-623330373034/DJI_0057.JPG -> DJI_0057_1.JPG
Started: https://angrytribe.org/old/downloaded_files/tild6566-3563-4764-b262-376461366539/DSC_4904.jpg
Completed: https://angrytribe.org/old/downloaded_files/tild3534-6633-4464-b435-326638386334/P9220984.JPG -> P9220984_1.JPG
Started: https://angrytribe.org/old/downloaded_files/tild6436-3535-4630-b438-356362646631/30180832214_298f969c.jpg
Completed: https://angrytribe.org/old/downloaded_files/tild6134-3764-4734-b232-306134393064/30565402375_cfb00f30.jpg -> 30565402375_cfb00f30_1.jpg
Started: https://angrytribe.org/old/downloaded_files/tild3930-3736-4361-a639-346466386334/2018-09-19_114015_1.jpg
Completed: https://angrytribe.org/old/downloaded_files/tild6435-3561-4566-b063-383461363533/on5tJvk40kQ.jpg -> on5tJvk40kQ_1.jpg
Started: https://angrytribe.org/old/downloaded_files/tild3432-3462-4864-b865-346236653435/P9231020.JPG
Completed: https://angrytribe.org/old/downloaded_files/tild6566-3563-4764-b262-376461366539/DSC_4904.jpg -> DSC_4904_1.jpg
Started: https://angrytribe.org/old/downloaded_files/tild6630-3435-4632-a232-333962303735/2016-11-20_061854_2.jpg
Completed: https://angrytribe.org/old/downloaded_files/tild6436-3535-4630-b438-356362646631/30180832214_298f969c.jpg -> 30180832214_298f969c_1.jpg
Started: https://upload.wikimedia.org/wikipedia/commons/7/74/A-Cat.jpg
Completed: https://angrytribe.org/old/downloaded_files/tild3930-3736-4361-a639-346466386334/2018-09-19_114015_1.jpg -> 2018-09-19_114015_1_1.jpg
Started: https://onlinejpgtools.com/images/examples-onlinejpgtools/mountains-near-water.jpg?size=500x500
Completed: https://angrytribe.org/old/downloaded_files/tild3432-3462-4864-b865-346236653435/P9231020.JPG -> P9231020_1.JPG
Started: https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg
Completed: https://angrytribe.org/old/downloaded_files/tild6630-3435-4632-a232-333962303735/2016-11-20_061854_2.jpg -> 2016-11-20_061854_2_1.jpg
Completed: https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg -> Cat03_1.jpg
Completed: https://upload.wikimedia.org/wikipedia/commons/7/74/A-Cat.jpg -> A-Cat_1.jpg
Completed: https://onlinejpgtools.com/images/examples-onlinejpgtools/mountains-near-water.jpg?size=500x500 -> mountains-near-water_1.jpg

All downloads completed successfully!
```