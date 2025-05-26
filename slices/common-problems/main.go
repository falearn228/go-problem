package main

import (
	"fmt"
	"time"
)

// Задача: Система логирования с динамическими слайсами

// задача для изучения всех тонкостей работы с append и слайсами
// Условие:

// Создайте систему логирования, которая эффективно собирает и обрабатывает лог-сообщения. Система должна демонстрировать различные аспекты работы со слайсами и append.

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

type LogEntry struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
}

type Logger struct {
	entries []LogEntry
	buffer  []LogEntry // для демонстрации разных техник
}

func (l *Logger) Log(level LogLevel, message string) {
	// TODO: Добавьте новую запись в entries используя append
	// Создайте LogEntry с текущим временем
	// Измерьте, как изменяется capacity при добавлении записей

	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
	}
	l.entries = append(l.entries, entry)

	// Ваша реализация здесь
}

func (l *Logger) PrintCapacityInfo() {
	// TODO: Выведите len и cap слайса entries
	// Покажите, как capacity растет при append
	fmt.Println("cap ->", cap(l.entries), "len ->", len(l.entries))
}

func (l *Logger) LogBatch(messages []string, level LogLevel) {
	// TODO: Эффективно добавьте множество сообщений
	// Сравните производительность с обычным append

	for _, message := range messages {
		entry := LogEntry{
			Timestamp: time.Now(),
			Level:     level,
			Message:   message,
		}

		l.entries = append(l.entries, entry)
	}
}

func (l *Logger) LogBatchOptimized(messages []string, level LogLevel) {
	// TODO: Предварительно выделите память для всех сообщений
	// Используйте make([]LogEntry, 0, expectedSize)
	// Затем используйте append как обычно

	newEntries := make([]LogEntry, 0, len(messages))

	for _, message := range messages {
		entry := LogEntry{
			Timestamp: time.Now(),
			Level:     level,
			Message:   message,
		}

		newEntries = append(newEntries, entry)
	}

	l.entries = append(l.entries, newEntries...)
}

func (l *Logger) GetRecentLogs(count int) []LogEntry {
	// TODO: Верните последние count записей
	// ВНИМАНИЕ: покажите проблему с shared underlying array
	// Затем исправьте это с помощью copy
	if count > len(l.entries) {
		count = len(l.entries)
	}
	// ❌ ОПАСНО: возвращает slice, который делит underlying array
	return l.entries[len(l.entries)-count:]
}

func (l *Logger) GetRecentLogsSafe(count int) []LogEntry {
	// TODO: Безопасная версия, которая не делит underlying array
	// Используйте copy для создания независимой копии
	safecopy := make([]LogEntry, count)

	copy(safecopy, l.entries)
	return safecopy
}

func MergeLogs(loggers ...*Logger) []LogEntry {
	// TODO: Объедините логи из нескольких логгеров
	// Используйте append с ... для слияния слайсов
	// Оптимизируйте память, предварительно посчитав общий размер
	totalLength := 0
	for _, logger := range loggers {
		totalLength += len(logger.entries)
	}
	merged := make([]LogEntry, 0, totalLength)
	for i := range loggers {
		merged = append(merged, loggers[i].entries...)
	}
	return merged
}

func MergeLogsCopy(loggers ...*Logger) []LogEntry {
	// TODO: версия с copy
	// Сначала посчитайте общее количество записей
	// Затем выделите память одним вызовом make
	// Используйте copy для эффективного копирования
	totalEntries := 0
	for _, logger := range loggers {
		totalEntries += len(logger.entries)
	}
	merged := make([]LogEntry, totalEntries)

	for _, logger := range loggers {
		copy(merged, logger.entries)
	}

	return merged
}

func (l *Logger) FilterByLevel(minLevel LogLevel) []LogEntry {
	// TODO: Верните все записи с уровнем >= minLevel
	// Используйте append для построения результата
	// Сравните два подхода: с предаллокацией и без
	entries := make([]LogEntry, 0, len(l.entries))
	for _, v := range l.entries {
		if v.Level >= minLevel {
			entries = append(entries, v)
		}
	}
	return entries
}

func (l *Logger) FilterByLevelInPlace(minLevel LogLevel) {
	// TODO: Удалите записи с уровнем < minLevel из текущего слайса
	// Используйте технику "append in place"
	// Покажите, как это влияет на capacity

	originalLen := len(l.entries)
	originalCap := cap(l.entries)

	// "Append in place" - перезаписываем тот же слайс
	writeIndex := 0 // Позиция для записи "хороших" элементов

	for _, v := range l.entries {
		if v.Level >= minLevel {
			l.entries[writeIndex] = v
			writeIndex++
		}
	}

	l.entries = l.entries[:writeIndex]

	fmt.Printf("Фильтрация in-place: %d -> %d элементов, capacity: %d -> %d\n",
		originalLen, len(l.entries), originalCap, cap(l.entries))

}

func BenchmarkAppendVsAssign() {
	// TODO: Сравните производительность:
	// 1. Создание слайса с известным размером и присваивание по индексам
	// 2. Создание пустого слайса и append
	// 3. Создание слайса с capacity и append

	length := 50000

	case1 := make([]int, length)
	time1 := time.Now()
	for i := range length {
		case1[i] = i
	}
	time1End := time.Since(time1)
	fmt.Println("time1End ->", time1End)

	case2 := make([]int, 0)
	time2 := time.Now()
	for i := range length {
		case2 = append(case2, i)
	}
	time2End := time.Since(time2)
	fmt.Println("time2End ->", time2End)

	case3 := make([]int, 0, length)
	time3 := time.Now()
	for i := range length {
		case3 = append(case3, i)
	}
	time3End := time.Since(time3)
	fmt.Println("time3End ->", time3End)
}

func DemonstrateCapacityGrowth() {
	// TODO: Покажите, как растет capacity при последовательных append
	// Выведите capacity после каждого добавления
	// Объясните паттерн роста (обычно удваивается)
	length := 50000
	prevCap := 1.0

	case1 := make([]int, 0)
	for i := range length {
		case1 = append(case1, i)
		if float64(cap(case1)) != prevCap {
			if float64(cap(case1))/float64(prevCap) < 2 {
				fmt.Println(len(case1))
			}
			fmt.Println("cap ->", float64(cap(case1))/float64(prevCap))
			prevCap = float64(cap(case1))
		}
	}
}

func main() {
	logger := &Logger{}

	fmt.Println("=== 1. Базовое логирование ===")
	logger.Log(INFO, "Система запущена")
	logger.Log(WARNING, "Предупреждение")
	logger.Log(ERROR, "Ошибка")
	logger.PrintCapacityInfo()

	// Добавим много записей чтобы увидеть рост capacity
	for i := 0; i < 20; i++ {
		logger.Log(DEBUG, fmt.Sprintf("Debug сообщение %d", i))
		if i%5 == 0 {
			fmt.Printf("После %d записей: ", len(logger.entries))
			logger.PrintCapacityInfo()
		}
	}
	fmt.Println("\n=== 2. Пакетное логирование ===")
	messages := []string{"Msg1", "Msg2", "Msg3", "Msg4", "Msg5"}

	// Тест обычного подхода
	start := time.Now()
	logger.LogBatch(messages, INFO)
	fmt.Printf("Обычный batch: %v\n", time.Since(start))

	// Тест оптимизированного подхода
	start = time.Now()
	logger.LogBatchOptimized(messages, INFO)
	fmt.Printf("Оптимизированный batch: %v\n", time.Since(start))

	fmt.Println("\n=== 3. Проблема shared array ===")
	recent := logger.GetRecentLogs(3)
	fmt.Printf("Получили %d записей и  cap(%d)\n", len(recent), cap(recent))
	fmt.Printf("Before recent logs: %v\n", recent)

	// Изменим оригинальный logger
	logger.Log(ERROR, "Новая ошибка")
	logger.Log(ERROR, "Новая ошибка")
	logger.Log(ERROR, "Новая ошибка")
	logger.Log(ERROR, "Новая ошибка")
	logger.Log(ERROR, "Новая ошибка")
	logger.Log(ERROR, "Новая ошибка")
	logger.Log(ERROR, "Новая ошибка")
	logger.Log(ERROR, "Новая ошибка")
	logger.Log(ERROR, "Новая ошибка")
	logger.Log(ERROR, "Новая ошибка")

	fmt.Println("После добавления новой записи:")
	fmt.Printf("Recent logs: %v\n", recent) // Может быть изменен!

	// Безопасная версия
	recentSafe := logger.GetRecentLogsSafe(3)
	logger.Log(ERROR, "Еще одна ошибка")
	fmt.Printf("Safe recent logs: %v\n", recentSafe) // Должен остаться неизменным

	fmt.Println("\n=== 4. Слияние логов ===")
	logger2 := &Logger{}
	logger2.Log(INFO, "Logger2 сообщение")

	merged := MergeLogs(logger, logger2)
	fmt.Printf("Всего объединенных записей: %d\n", len(merged))

	mergedOpt := MergeLogsCopy(logger, logger2)
	fmt.Printf("Оптимизированное слияние: %d записей\n", len(mergedOpt))

	fmt.Println("\n=== 5. Фильтрация ===")
	errors := logger.FilterByLevel(ERROR)
	fmt.Printf("Найдено ошибок: %d\n", len(errors))

	originalCount := len(logger.entries)
	logger.FilterByLevelInPlace(WARNING) // Удаляем DEBUG и INFO
	fmt.Printf("После фильтрации на месте: %d -> %d записей\n",
		originalCount, len(logger.entries))

	fmt.Println("\n=== 6. Демонстрация производительности ===")
	BenchmarkAppendVsAssign()

	fmt.Println("\n=== 7. Рост capacity ===")
	DemonstrateCapacityGrowth()
}
