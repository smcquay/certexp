# certexp

report certificate expiry for a collection of servers, which yields equivalent
information to:

```bash
$ echo | openssl s_client -connect $hostname:$port 2> /dev/null | openssl x509 -noout -dates | grep notAfter
```

## example usage

```bash
$ cat sites.txt
apple.com
google.com
amazon.com
imap.gmail.com:993
$ cat sites.txt | certexp
apple.com                2018-10-31 23:59:59 +0000 UTC
google.com               2018-02-13 15:19:00 +0000 UTC
amazon.com               2018-09-21 23:59:59 +0000 UTC
imap.gmail.com           2018-02-27 09:29:00 +0000 UTC
```
