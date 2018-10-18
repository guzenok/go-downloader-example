# downloader

тестовое задание

## Задание

Программа - принимает файл с URLs через параметр командной строки, параллельно скачивает их и сохраняет в базу syndtr/goleveldb.
Ключ: URL, значение: содержимое страницы.

Пример запуска: ./go_downloader -urls=./urls.txt

Использовать библиотеки:
 - https://godoc.org/github.com/syndtr/goleveldb
   - https://godoc.org/github.com/syndtr/goleveldb/leveldb
 - https://golang.org/pkg/
   - https://golang.org/pkg/flag/
 - https://golang.org/pkg/errors/
 - https://golang.org/pkg/os/
 - https://golang.org/pkg/net/http/
 - https://golang.org/pkg/sync/#WaitGroup

И подходы:
 - https://blog.golang.org/share-memory-by-communicating
 - https://blog.golang.org/concurrency-is-not-parallelism


## Решение

```
for q in $(echo {10..100}); do echo "https://www.google.com/search?q=1${q}"; done > urls.txt
go run ./main.go -urls=./urls.txt
rm -f urls.txt
```
