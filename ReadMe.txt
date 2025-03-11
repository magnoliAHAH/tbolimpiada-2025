Использование

Для использования необходимо клонировать репозиторий:

git clone https://github.com/magnoliAHAH/tbolimpiada-2025.git

Перейдите в директорию проекта:

cd tbolimpiada-2025

Запуск с использованием Docker

Создайте Docker-образ:

sudo docker build -t pathfinder .

Запустите контейнер:

sudo docker run -p 8080:8080 --name pathfinder pathfinder

Теперь REST API сервер работает на порту 8080. Вы можете взаимодействовать с ним с помощью curl, терминала Windows или специализированных приложений для работы с API (например, Postman).

Использование скриптов для запросов

В директории tbolimpiada-2025/api находятся скрипты:

curl-command.sh — для Linux

windows-request.txt — для Windows

Для Linux

Измените IP-адрес и при необходимости карту в файле curl-command.sh.

Сделайте файл исполняемым:

chmod +x ./api/curl-command.sh

Выполните скрипт:

sudo ./api/curl-command.sh

Пример вывода команды curl:

MeetingPoint: X:10, Y:15

Для Windows

Откройте файл windows-request.txt, измените IP-адрес и карту, если требуется.

Скопируйте команду и выполните её в командной строке (cmd) или PowerShell.

Пример вывода для Windows:

MeetingPoint: X:10, Y:15

Ручной запуск программы без Docker

Вы можете запустить основную программу без Docker, если установлен Go.

Соберите исполняемый файл:

go build -o pathfinder ./game

Выполните команду с указанием пути к карте:

./pathfinder -m [путь до карты]

Либо используйте go run:

go run ./game/main.go -m [путь до карты]

Пример вывода команды go run с информацией о карте:

Карта загружена. Размер: 20x20

Пример вывода с отмеченными путями: