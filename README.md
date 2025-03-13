## **Описание проекта**

Simple Service – это REST API-сервис, написанный на Go с использованием фреймворка Fiber и PostgreSQL. Сервис предоставляет базовый функционал для управления задачами.

Реализовано:

- Создание задач через API
  - запрос POST /v1/tasks
  - {"Title":"ExampleTitle", "Description":"ExampleDescription", "Status":"ExampleStatus"}
- Изменение статуса задачи по ID
  - запрос PUT /v1/tasks​/:id
  - {"Status":"ExampleNewStatus"}
- Получение всех задач
  - запрос GET /v1/tasks
- Получение задачи по ID
  - запрос GET /v1/task​/:id
- Удаление задачи по ID
  - запрос DELETE /v1/task​/:id
- Валидация входных данных
- Логирование с использованием `zap`
- Хранение данных в PostgreSQL
- Подключение через `pgxpool` для эффективного управления соединениями

---

## **1️⃣ Подготовка окружения**

### **1.1 Установка зависимостей**

Перед запуском убедитесь, что у вас установлены:

- Go
- DataGrip или аналогичное приложение
- Postman или Insomnia для тестирования API

### **1.2 Клонирование репозитория**

```
git clone https://github.com/yourusername/ToDo-crud.git
cd ToDo-crud
```

---

## **2 Настройка проекта**

### **3.1 Создание `.env` файла**

Создайте `.env` файл и пропишите параметры:

```
LOG_LEVEL=info
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin
POSTGRES_DB=simple_service
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
REST_LISTEN_ADDRESS=:8080
REST_TOKEN=your_secret_token

```

Также установите плагин в вашу IDLE.
Я использую ее: <https://github.com/Ashald/EnvFile>

### **3.2 Применение миграций**

Создайте таблицу `tasks` в базе данных:

```
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

```

---

## **4️⃣ Запуск сервиса**

### **4.1 Локальный запуск**

```
go run cmd/main.go

```

Сервис будет доступен по адресу `http://localhost:8080`

---

## **5️⃣ Тестирование API**

### **5.1 Создание задачи**

**Запрос:**

```
POST http://localhost:8080/v1/tasks
Content-Type: application/json
Authorization: Bearer your_secret_token

```

```
{
  "title": "New Feature",
  "description": "Develop new API endpoint"
  "status": "New"
}

```

**Ответ:**

```
{
  "status": "success",
  "data": {
    "task_id": 1
  }
}

```

---

### **5.2 Получение задачи по id**

**Запрос:**

```
GET http://localhost:8080/v1/tasks/1
Content-Type: application/json
Authorization: Bearer your_secret_token

```

**Ответ:**

```
{
  "id": "1"
  "title": "New Feature",
  "description": "Develop new API endpoint"
}

```

### **5.3 Получение списка всех задач**

**Запрос:**

```
GET http://localhost:8080/v1/tasks/
Content-Type: application/json
Authorization: Bearer your_secret_token

```

**Ответ:**

```
{
    "status": "Status OK",
    "data": {
        "1": {
            "id": 1,
            "title": "Task number one",
            "description": "There is description about task number one",
            "status": "New"
        },
        "2": {
            "id": 2,
            "title": "Task number two",
            "description": "There is description about task number two",
            "status": "New"
        }
    }
}

```

### **5.4 Изменение/обновление задачи по id**

**Запрос:**

```
PUT http://localhost:8080/v1/tasks/1
Content-Type: application/json
Authorization: Bearer your_secret_token

```

**Body:**

```
{
"id": 1,
"title": "Task number one exchange",
"description": "There is description about task number one exchange",
"status": "Done"
}
```

**Ответ:**

```
{
  "status": "Status OK",
    "data": {
    "task_id": 1
    }
}

```

### **5.5 Удаление задачи по id**

**Запрос:**

```
DELETE http://localhost:8080/v1/tasks/1
Content-Type: application/json
Authorization: Bearer your_secret_token

```

**Ответ:**

```
{
    "status": "Status OK",
    "data": {
        "task_id": 1
    }
}

```

---

## **Дополнительная информация**

- Файл `docs/openapi.yaml` содержит документацию API в формате OpenAPI 3.0
- Логирование ведётся через `zap.Logger`
- Переменные окружения загружаются через `envconfig`
- Соединение с PostgreSQL осуществляется через `pgxpool`

Сервис готов к работе.
