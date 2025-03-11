# Pathfinder API

Система поиска пути в лабиринте с учётом различных типов местности и особенностей персонажей.

## Использование

### Клонирование репозитория

```bash
git clone https://github.com/magnoliAHAH/tbolimpiada-2025.git
cd tbolimpiada-2025
```

## Запуск с использованием Docker

1. Создайте Docker-образ:

```bash
sudo docker build -t pathfinder .
```

2. Запустите контейнер:

```bash
sudo docker run -p 8080:8080 --name pathfinder pathfinder
```

REST API сервер будет доступен на `http://localhost:8080`.

### Примеры запросов

Для взаимодействия с API можно использовать `curl` или другие инструменты (например, Postman).

#### Пример POST-запроса с передачей карты:

```bash
curl -X POST http://localhost:8080/process \
     -H "Content-Type: application/json" \
     -d '{"maze": "#####\n#...#\n#.#.#\n#...#\n#####"}'
```

Пример успешного ответа:

```
MeetingPoint: X:2, Y:2
```

## Использование скриптов для запросов

В директории `tbolimpiada-2025/api` находятся вспомогательные скрипты:

- `curl-command.sh` — для Linux
- `windows-request.txt` — для Windows

### Для Linux

1. Настройте IP-адрес и карту в `curl-command.sh`, при необходимости.

2. Сделайте скрипт исполняемым:

```bash
chmod +x ./api/curl-command.sh
```

3. Выполните скрипт:

```bash
sudo ./api/curl-command.sh
```

### Для Windows

1. Настройте IP-адрес и карту в `windows-request.txt`.

2. Скопируйте команду и выполните её в командной строке (cmd) или PowerShell.

## Ручной запуск программы без Docker

Если у вас установлен Go, вы можете запустить приложение напрямую.

1. Соберите исполняемый файл:

```bash
go build -o pathfinder ./game
```

2. Запустите программу, указав путь до карты:

```bash
./pathfinder -m [путь_до_карты]
```

Или выполните напрямую:

```bash
go run ./game/main.go -m [путь_до_карты]
```

### Пример вывода:

```
Карта загружена. Размер: 20x20
Путь построен. Оптимальная точка сбора: X:10, Y:15
```

## Остановка и удаление контейнера

Чтобы остановить контейнер, выполните:

```bash
sudo docker stop pathfinder
```

Для удаления контейнера:

```bash
sudo docker rm pathfinder
```

## Лицензия

Этот проект распространяется под лицензией MIT.

