# Star Wars Calculator для macOS (Go)

Небольшое GUI-приложение калькулятора в стиле Star Wars для **macOS**, написанное на Go (Cocoa + WebKit через cgo).

## Что исправлено
Если раньше `go run .` падал с ошибками вида `syntax error ... main_darwin.go`, это исправлено: логика HTML/JS вынесена из cgo-комментария в Go-строку, поэтому парсинг Go больше не ломается.

## Возможности
- Окно **Galactic Calculator** в тёмной космической теме.
- Кнопки `0-9`, `00`, `.`, `+`, `-`, `*`, `/`, `=`.
- `C` — очистка, `⌫` — удалить последний символ.
- Защита от деления на ноль.
- Пасхалка по кнопке `◼`.

## Требования (macOS)
1. **macOS** (рекомендуется актуальная версия).
2. **Go 1.22+**.
3. **Xcode Command Line Tools**:

```bash
xcode-select --install
```

## Быстрый запуск локально (macOS)
В корне проекта:

```bash
go run .
```

После запуска откроется GUI-окно калькулятора.

---

## Сборка бинарника

```bash
go build -o galactic-calculator .
./galactic-calculator
```

---

## Сборка `.app` для запуска двойным кликом

В проект добавлен скрипт `scripts/build-macos-app.sh`, который создаёт полноценный app bundle.

### 1) Выполнить сборку

```bash
./scripts/build-macos-app.sh
```

После этого появится:

```text
build/Galactic Calculator.app
```

### 2) Запустить приложение

```bash
open "build/Galactic Calculator.app"
```

### 3) Если macOS блокирует запуск (Gatekeeper)
Для локальной неподписанной сборки можно снять quarantine-атрибут:

```bash
xattr -dr com.apple.quarantine "build/Galactic Calculator.app"
```

### 4) (Опционально) Подпись для локального запуска
Если хотите подписать своим Developer сертификатом:

```bash
codesign --force --deep --sign "Developer ID Application: YOUR_NAME" "build/Galactic Calculator.app"
```

Проверка подписи:

```bash
codesign --verify --deep --strict --verbose=2 "build/Galactic Calculator.app"
```

## Файлы проекта
- `main_darwin.go` — основной GUI для macOS.
- `main.go` — fallback для не-macOS окружений.
- `scripts/build-macos-app.sh` — сборка app bundle (`.app`).
- `go.mod` — модуль Go.
