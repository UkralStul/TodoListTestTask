# Todo List App

## Описание

Приложение для управления задачами, развернутое в контейнерах с использованием Docker Compose.

## Запуск проекта

### 1. Клонирование репозитория

```sh
git clone https://github.com/your-repo/todolist.git
cd todolist
```
### 2. Создание переменной среды в файле .env

```sh
DATABASE_URL=postgres://postgres:7243@localhost:5432/todoList?sslmode=disable
```

### 2. Запуск контейнеров

```sh
docker-compose up --build
```


