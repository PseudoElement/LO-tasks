# LO-tasks

**Мануал по использованию**

Запуск сервера производится командой `make start-prod`. После этого можно отправлять запросы с клиента или curl.

Основная сущность:
# type BasicTask struct {
#	Creator   string `json:"creator"`   // имя автора задачи
#	Info      string `json:"info"`      // описание
#	Completed bool   `json:"completed"` // статус
# }


1) создать задачу - отправить post запрос http://localhost:8080/tasks с телом запроса типа BasicTask(ключами использовать json значения). 
Пример запроса: `curl -X POST -H "Content-Type: application/json" -d '{"creator": "borrow", "info": "Apply for a job in LO.", "completed": false}' http://localhost:8080/tasks`
Пример респонса:
# {
#   "creator": "Golang dev",
#   "info": "Apply for a job in LO.",
#   "completed": true,
#   "created_at": "13-08-2025 22:37",
#   "id": "04025f99-c2eb-4dee-a4ef-5da617a66fb3"
# }

2) получить задачу -  отправить get запрос http://localhost:8080/tasks. С помощью id в дальнейшем можно получить конкретную задачу. 
Пример запроса: `curl http://localhost:8080/tasks/04025f99-c2eb-4dee-a4ef-5da617a66fb3`.
Пример респонса: 
# {
#   "creator": "Golang dev",
#   "info": "Apply for a job in LO.",
#   "completed": true,
#   "created_at": "13-08-2025 22:37",
#   "id": "04025f99-c2eb-4dee-a4ef-5da617a66fb3"
# }

3) получить все задачи - отправить get запрос http://localhost:8080/tasks. Можно передать опциональный параметр `?completedFirst=true` - тогда все запросы с "completed: true" будут первыми в списке, по дефолту `completedFirst=false`
Пример запроса: `http://localhost:8080/tasks?completedFirst=true`
Пример респонса: 
# [
# {
#    "creator": "basd1231orrow",
#    "info": "Apply for a job in LO.",
#    "completed": true,
#    "created_at": "13-08-2025 22:37",
#    "id": "04025f99-c2eb-4dee-a4ef-5da617a66fb3"
# },
# {
#   "creator": "borrow",
#    "info": "Apply for a job in LO.",
#    "completed": true,
#    "created_at": "13-08-2025 22:37",
#    "id": "dd8b958b-8e67-4acc-951e-1bf72d15d304"
# },
# {
#    "creator": "borrow",
#    "info": "Apply for a job in LO.",
#    "completed": false,
#    "created_at": "13-08-2025 22:36",
#    "id": "d66b51db-35eb-42de-a942-54710fd3dbd7"
# }
# ]