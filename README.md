# Star Wars Calculator для macOS (Go)

Небольшое GUI-приложение калькулятора в стиле Star Wars для **macOS**, написанное на Go.

## Что внутри
- Окно приложения **Galactic Calculator** в тёмной космической теме.
- Кнопки `0-9`, `00`, `.`, `+`, `-`, `*`, `/`, `=`.
- `C` — очистка, `⌫` — удалить последний символ.
- Пасхалка по кнопке `◼`.

## Требования (macOS)
1. **macOS** (рекомендуется актуальная версия).
2. **Go 1.22+**.
3. **Xcode Command Line Tools**:

```bash
xcode-select --install
```

## Запуск локально на macOS
В корне проекта выполните:

```bash
go run .
```

После запуска откроется GUI-окно калькулятора.

## Сборка бинарника
```bash
go build -o galactic-calculator .
```

Затем запуск:

```bash
./galactic-calculator
```

## Файлы проекта
- `main_darwin.go` — основной GUI для macOS (Cocoa + WebKit через cgo).
- `main.go` — fallback для не-macOS окружений.
- `go.mod` — модуль Go.
