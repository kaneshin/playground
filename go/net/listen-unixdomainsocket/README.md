# Listen UNIX Domain Socket

## Usage

```shell
go run main.go &
echo -en 'GET / HTTP/1.0\r\n\r\n' | socat stdio UNIX-CONNECT:/var/run/gopher/go.sock
```

## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)

## Author

[Shintaro Kaneko](https://github.com/kaneshin) <kaneshin0120@gmail.com>

