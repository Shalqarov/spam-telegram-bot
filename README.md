# spam-telegram-bot

<h2>Как запустить сервис:<h2>

1.Скачиваем зависимости:

```bash
go mod tidy
```
2.Создаем бд->например:
```bash
example.db
```

3.Настраиваем конфиг в example.toml


4.Запускаем сервис:

```bash
go run ./cmd [flags: --config, --addr]
```

