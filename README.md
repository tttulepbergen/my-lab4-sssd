## Lab #4 – Демонстрация уязвимостей и их исправление (Go)

Минимальный HTTP-сервер на Go с двумя версиями:

- `vulnerable/` — демонстрация уязвимостей (для анализа в отчёте)
- `fix/` — исправленная версия (безопаснее, но с той же функциональностью)

### Структура

```
.
├─ vulnerable/
│  ├─ main.go
│  ├─ config.go
│  └─ deserialize.go
├─ fix/
│  ├─ main.go
│  ├─ config.go
│  └─ deserialize.go
├─ .gitignore
└─ env.example
```

### Как запустить

Требуется Go 1.20+.

Запуск уязвимой версии:

```
cd vulnerable
go run .
```

Запуск исправленной версии:

```
cd fix
cp ../env.example .env   # опционально: заполнить переменные окружения
APP_API_KEY=TEST_KEY_PLACEHOLDER go run .
```

Сервер слушает `:8080`.

### Эндпоинты

- `GET /` — демонстрация обработки ошибок
- `POST /create-user` — демонстрация (не)безопасной десериализации JSON
- `GET /config` — демонстрация работы с «секретами»

### Что демонстрируется в `vulnerable/`

1) Ошибки и стек-трейсы попадают пользователю (паника без recover)
2) Хардкод «секрета» в коде (`APIKey`)
3) Неконтролируемая десериализация JSON и доверие полю `is_admin`

### Что исправлено в `fix/`

- Добавлен middleware `recover` и возвращается обобщённый `500 Internal Server Error`
- Таймауты сервера и ограничение размера тела запроса
- Загрузка секретов из переменных окружения, `.env` не коммитится (см. `.gitignore` и `env.example`)
- Жёсткая схема JSON, `DisallowUnknownFields`, валидация, запрет клиентского `is_admin`
- Эндпоинт `/config` больше не раскрывает секрет

### Быстрые проверки (для скриншотов в отчёте)

1) Ошибки/stack trace:

```
# vulnerable: вернёт стек-трейс в ответ
curl -i http://localhost:8080/

# fix: вернёт 500 без деталей, стек уйдёт только в логи
curl -i http://localhost:8080/
```

2) Хардкод секретов:

```
# vulnerable: секрет доступен (плохо)
curl -s http://localhost:8080/config

# fix: секрет не показывается пользователю
curl -s http://localhost:8080/config
```

3) Десериализация:

```
# vulnerable: клиент может прислать is_admin=true, сервер «поверит»
curl -i -X POST http://localhost:8080/create-user \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice","is_admin":true}'

# fix: неизвестные/опасные поля отклоняются, требуется username
curl -i -X POST http://localhost:8080/create-user \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice","is_admin":true}'
```

### Риски 

- Утечка внутренних ошибок и стеков раскрывает структуру приложения и упрощает атаки
- Хранение секретов в коде и историях Git ведёт к компрометации ключей
- Неконтролируемая десериализация позволяет эскалацию привилегий и DoS через большие тела


## Instances' Structure
```
Users
• UserID (AUTO_INCREMENT PRIMARY KEY)
• User_Email(TEXT)
• Username (VARCHAR(30))
• Password (encrypted password)
• Number_of_phone_user(VARCHAR(50))
• Role(Default “User”)
Animal
• ID(AUTO_INCREMENT PRIMARY KEY)
• Kind_Of_Animal(VARCHAR(255))
• Breed_Of_Animal(VARCHAR(255))
• Name(VARCHAR(255))
• Age(INTEGER)
• Description(TEXT)
Admins
• AdminID(AUTO_INCREMENT PRIMARY KEY)
• Admin_Email(TEXT)
• Adminame (VARCHAR(30))
• Password (encrypted password)
• Number_of_phone_Admin(VARCHAR(50))
• Role(VARCHAR(255), Type:Back_end, Front_end, G_admin, etc)
Role
• Permissions(TEXT)
```
