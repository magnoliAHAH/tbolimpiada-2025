# Pathfinder API

Система поиска пути в лабиринте с учётом различных типов местности.

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

![image](https://github.com/user-attachments/assets/221ccc19-bbb3-4bcb-b463-0abf6d4b6c1e)


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
Пример вывода:
![image](https://github.com/user-attachments/assets/ed5b3cc7-e5ac-4f13-adb2-e06b176d307b)


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
![image](https://github.com/user-attachments/assets/20492d9e-7597-4abc-bcd8-3cfc73908718)  
Шапка вычислений с основными данными  

![image](https://github.com/user-attachments/assets/7d06381f-1bd2-42ab-aefd-7d28dcffde6a)  
Отмечены пути, раскрашена карта  


### Использованные технологии

Go (Golang) — https://golang.org

Docker — https://www.docker.com

cURL — https://curl.se

Linux/Windows — поддержка различных ОС

REST API — архитектурный стиль для взаимодействия сервисов
