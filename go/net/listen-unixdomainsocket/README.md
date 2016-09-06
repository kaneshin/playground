# Listen UNIX Domain Socket

## Usage

```shell
echo d /var/run/gopher 0755 `whoami` `whoami` - > /etc/tmpfiles.d/gopher.conf
systemd-tmpfiles --create /etc/tmpfiles.d/gopher.conf
systemctl daemon-reload

go run main.go &
echo -en 'GET / HTTP/1.0\r\n\r\n' | socat stdio UNIX-CONNECT:/var/run/gopher/go.sock

HTTP/1.0 200 OK
Date: Tue, 06 Sep 2016 16:45:16 GMT
Content-Length: 9
Content-Type: text/plain; charset=utf-8

It works!
```

## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)

## Author

[Shintaro Kaneko](https://github.com/kaneshin) <kaneshin0120@gmail.com>

