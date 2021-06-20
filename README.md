# Runner API

## Основная функциональность
* API предоставляет методы для записи, изменения и листинга runner'ов
* Runner'ы необходимы для исполнения кода решения задачи

## Runner
```go
type Runner struct {
    guid string  // уникальный идентификатор
    os string    // операционная система
    arch string  // архитектура машины
}
```

## Контракт
| Метод              | Сигнатура                                   | Описание                                              |
|--------------------|---------------------------------------------|-------------------------------------------------------|
| CreateRunner       | os string, arch string                      | создает новый runner с уникальным guid'ом             |
| MultiCreateRunner  | {os string, arch string}[], batchSize int32 | создает N новых runner'ов с уникальными guid'ами      |
| UpdateRunner       | guid string, os string, arch string         | обновляет поля runner'а по guid'у                     |
| RemoveRunner       | guid string                                 | удаляет существующий раннер по guid'у                 |
| DescribeRunner     | guid string                                 | возвращает runner по guid'у                           |
| ListRunners        | {guid []string, os []string, arch []string} | возвращает все подходящие под фильтр runner'ы         |