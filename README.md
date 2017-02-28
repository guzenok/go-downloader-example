# go_downloader

Тестовое задание.

## ТЗ

Программа - принимает файл с URLs через параметр командной строки, параллельно скачивает их и сохраняет в базу syndtr/goleveldb.
Ключ: URL, значение: содержимое страницы.
Пример запуска: ./downloader -urls=./urls.txt
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
