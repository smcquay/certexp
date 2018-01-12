# certexp

A tool that reports certificate expiry for a collection of servers. It yields
equivalent information to:

```bash
$ echo | openssl s_client -connect $hostname:$port 2> /dev/null | openssl x509 -noout -dates | grep notAfter
```

but fetches the information for a colleciton of hosts and does it concurrently.

## example usage

```bash
$ certexp google.com amazon.com imap.gmail.com:993
google.com:443           2018-03-07 13:01:00 +0000 UTC
imap.gmail.com:993       2018-03-07 13:02:00 +0000 UTC
amazon.com:443           2018-09-21 23:59:59 +0000 UTC
```

or to stdin:

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
