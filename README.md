# Как работать с проектом
## 1. Установка
  1. Убедитесь, что у вас установлен Go версии 1.18+

  2. Клонируйте репозиторий:

**git clone https://github.com/yourusername/deadlock-tutorial.git**
**cd DeadLock**

  3. Скачайте зависимости:

**go mod download**
## 2. Задачи

**Задача 1: Банковские переводы**
**Проблема**: Реализация перевода денег между счетами содержит deadlock при одновременных взаимных переводах.

**Файлы:**

**taskBank/problem/bank.go** - код для исправления

**taskBank/solution/bank.go** - правильная реализация (как вариант исправления)

**cmd/bank/main.go** - демонстрация проблемы

**Как работать:**

  1. Запустите демо:

**go run cmd/bank/main.go**

  2. Увидите сообщение о deadlock

  3. Исправьте код в taskBank/problem/bank.go

  4. Проверьте исправление:

**go test ./taskBank/tests/...**

Если тесты проходят значит deadlock теперь не возникает

**Задача 2: Числа Фибоначчи**
**Проблема:** Рекурсивный расчет чисел Фибоначчи с кэшированием приводит к deadlock.

**Файлы:**

**taskFibonacci/problem/fibonacci.go** - код для исправления

**taskFibonacci/solution/fibonacci.go** - правильная реализация (как вариант исправления)

**cmd/fibonacci/main.go** - демонстрация проблемы

**Как работать:**

  1. Запустите демо:

**go run cmd/fibonacci/main.go**

 2. Увидите сообщение о deadlock

 3. Исправьте код в taskFibonacci/problem/fibonacci.go

 4. Проверьте исправление:

**go test ./taskFibonacci/tests/...**
