package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Validator interface {
	Validate(value interface{}) error
}

type StringValidator struct {
	name string
}

func (s StringValidator) Validate(value interface{}) error {
	return nil
}

type Converter interface {
	Convert(from interface{}) (interface{}, error)
}

func testNilInterfaces() {
	var validator Validator
	var stringValidator *StringValidator // nil pointer

	fmt.Printf("validator type -> %T, value -> %v\n", validator, validator)
	fmt.Printf("stringValidator type -> %T, value -> %v\n", stringValidator, stringValidator)
	fmt.Println(isReallyNil(validator)) // -> true

	validator = stringValidator
	fmt.Printf("AFTER validator type -> %T, value -> %v\n", validator, validator)
	fmt.Println(isReallyNil(validator)) // -> false, т.к interface теперь содержит тип, а значит не все его поля имеют nil
	// TODO: Реализуйте StringValidator
	// TODO: Продемонстрируйте разницу между nil interface и interface с nil pointer
	// TODO: Объясните почему validator == nil и stringValidator == nil дают разные результаты
	// TODO: Создайте функцию isReallyNil(interface{}) bool
}

func isReallyNil(i any) bool {
	if i == nil {
		return true
	}
	return false
}

func testTypeAssertions() {
	values := []interface{}{
		42,
		"hello",
		nil,
		(*string)(nil),
		[]int{1, 2, 3},
	}

	targetTypes := []string{"int", "string", "[]int", "*string"}

	for i, value := range values {
		fmt.Printf("\nЗначение %d: %v (тип: %T)\n", i, value, value)

		for _, targetType := range targetTypes {
			result, err := safeTypeAssertion(value, targetType)
			if err != nil {
				fmt.Printf("  -> %s: ERROR - %v\n", targetType, err)
			} else {
				fmt.Printf("  -> %s: SUCCESS - %v\n", targetType, result)
			}
		}
	}

	// TODO: Создайте функцию safeTypeAssertion(value interface{}, targetType string) (interface{}, error)
	// TODO: Обработайте все edge cases без panic
	// TODO: Используйте comma ok syntax правильно
	// TODO: Покажите разницу между value.(type) в switch и обычным assertion
}

func safeTypeAssertion(value interface{}, targetType string) (interface{}, error) {
	// Специальная обработка nil
	if value == nil {
		return nil, errors.New("value is nil")
	}

	switch targetType {
	case "int":
		if val, ok := value.(int); ok {
			return val, nil
		}
		return nil, fmt.Errorf("cannot convert %T to int", value)

	case "string":
		if val, ok := value.(string); ok {
			return val, nil
		}
		return nil, fmt.Errorf("cannot convert %T to string", value)

	case "float64":
		if val, ok := value.(float64); ok {
			return val, nil
		}
		return nil, fmt.Errorf("cannot convert %T to float64", value)

	case "[]int":
		if val, ok := value.([]int); ok {
			return val, nil
		}
		return nil, fmt.Errorf("cannot convert %T to []int", value)

	case "*string":
		if val, ok := value.(*string); ok {
			return val, nil
		}
		return nil, fmt.Errorf("cannot convert %T to *string", value)

	default:
		return nil, fmt.Errorf("unknown target type: %s", targetType)
	}
}
func testInterfaceComparison() {
	fmt.Println("=== Демонстрация panic при сравнении ===")
	demonstratePanicComparisons()

	fmt.Println("\n=== Тестирование safeEquals ===")
	testSafeEquals()

	fmt.Println("\n=== Edge cases с nil ===")
	testNilComparisons()

	fmt.Println("\n=== Сравнение одинаковых типов ===")
	testSameTypeComparisons()
}

// Демонстрация ситуаций где == вызывает panic
func demonstratePanicComparisons() {
	// Интерфейсы с несравнимыми типами
	panicCases := []struct {
		name string
		a, b interface{}
	}{
		{"slice comparison", []int{1, 2}, []int{1, 2}},
		{"map comparison", map[string]int{"a": 1}, map[string]int{"a": 1}},
		{"function comparison", func() {}, func() {}},
		{"channel comparison", make(chan int), make(chan int)},
		{"slice vs different slice", []int{1}, []string{"1"}},
	}

	for _, testCase := range panicCases {
		fmt.Printf("Тест: %s\n", testCase.name)

		// Небезопасное сравнение (будет panic)
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("  PANIC: %v\n", r)
				}
			}()

			result := testCase.a == testCase.b
			fmt.Printf("  Результат: %t (не должно выполниться)\n", result)
		}()

		// Безопасное сравнение
		result := safeEquals(testCase.a, testCase.b)
		fmt.Printf("  Safe result: %t\n", result)
		fmt.Println()
	}
}

// Безопасная функция сравнения интерфейсов
func safeEquals(a, b interface{}) bool {
	// Используем defer + recover для перехвата panic
	defer func() {
		recover() // Просто поглощаем panic
	}()

	// Специальная обработка nil
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Проверяем типы
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	// Проверяем на несравнимые типы заранее
	if !isComparable(a) {
		return deepEqual(a, b)
	}

	// Пытаемся обычное сравнение
	return a == b
}

// Проверка, является ли тип сравнимым
func isComparable(value interface{}) bool {
	if value == nil {
		return true
	}

	t := reflect.TypeOf(value)
	return t.Comparable()
}

// Глубокое сравнение для несравнимых типов
func deepEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// Тестирование safeEquals на различных случаях
func testSafeEquals() {
	testCases := []struct {
		name     string
		a, b     interface{}
		expected bool
	}{
		// Сравнимые типы
		{"equal ints", 42, 42, true},
		{"different ints", 42, 43, false},
		{"equal strings", "hello", "hello", true},
		{"different strings", "hello", "world", false},

		// Несравнимые типы
		{"equal slices", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"different slices", []int{1, 2}, []int{1, 3}, false},
		{"equal maps", map[string]int{"a": 1}, map[string]int{"a": 1}, true},
		{"different maps", map[string]int{"a": 1}, map[string]int{"a": 2}, false},

		// Разные типы
		{"different types", 42, "42", false},
		{"slice vs string", []byte("hello"), "hello", false},

		// Nil cases
		{"both nil", nil, nil, true},
		{"one nil", nil, 42, false},
		{"nil slice vs nil", ([]int)(nil), nil, false},
	}

	for _, tc := range testCases {
		result := safeEquals(tc.a, tc.b)
		status := "✅"
		if result != tc.expected {
			status = "❌"
		}
		fmt.Printf("%s %s: %t (expected %t)\n", status, tc.name, result, tc.expected)
	}
}

// Тестирование сравнений с nil
func testNilComparisons() {
	// Разные виды nil
	var nilInterface interface{} = nil
	var nilSlice []int = nil
	var nilMap map[string]int = nil
	var nilPointer *int = nil
	var nilFunc func() = nil

	nilCases := []struct {
		name string
		a, b interface{}
	}{
		{"nil interface vs nil interface", nilInterface, nilInterface},
		{"nil interface vs nil slice", nilInterface, nilSlice},
		{"nil slice vs nil slice", nilSlice, nilSlice},
		{"nil map vs nil map", nilMap, nilMap},
		{"nil pointer vs nil pointer", nilPointer, nilPointer},
		{"nil func vs nil func", nilFunc, nilFunc},
		{"nil slice vs nil map", nilSlice, nilMap},
	}

	for _, tc := range nilCases {
		fmt.Printf("%s:\n", tc.name)

		// Безопасное сравнение
		safeResult := safeEquals(tc.a, tc.b)
		fmt.Printf("  Safe equals: %t\n", safeResult)

		// Обычное сравнение (может panic-нуть)
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("  Normal == : PANIC - %v\n", r)
					return
				}
			}()

			normalResult := tc.a == tc.b
			fmt.Printf("  Normal == : %t\n", normalResult)
		}()

		// Анализ типов
		fmt.Printf("  Type A: %T, Type B: %T\n", tc.a, tc.b)
		fmt.Printf("  A == nil: %t, B == nil: %t\n", tc.a == nil, tc.b == nil)
		fmt.Println()
	}
}

// Тестирование сравнения одинаковых типов
func testSameTypeComparisons() {
	// Создаем одинаковые значения разными способами
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := slice1 // тот же underlying array

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}
	map3 := map1 // тот же underlying map

	fmt.Println("Slice comparisons:")
	fmt.Printf("  slice1 == slice2 (safe): %t\n", safeEquals(slice1, slice2))
	fmt.Printf("  slice1 == slice3 (safe): %t\n", safeEquals(slice1, slice3))
	fmt.Printf("  Same underlying array: %t\n", &slice1[0] == &slice3[0])

	fmt.Println("\nMap comparisons:")
	fmt.Printf("  map1 == map2 (safe): %t\n", safeEquals(map1, map2))
	fmt.Printf("  map1 == map3 (safe): %t\n", safeEquals(map1, map3))

}

// Продвинутая версия safeEquals с дополнительной информацией
func safeEqualsVerbose(a, b interface{}) (bool, string) {
	if a == nil && b == nil {
		return true, "both are nil interface"
	}
	if a == nil || b == nil {
		return false, "one is nil, other is not"
	}

	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)

	if aType != bType {
		return false, fmt.Sprintf("different types: %v vs %v", aType, bType)
	}

	if !aType.Comparable() {
		result := reflect.DeepEqual(a, b)
		return result, fmt.Sprintf("uncomparable type %v, used DeepEqual", aType)
	}

	// Пытаемся обычное сравнение
	defer func() {
		if r := recover(); r != nil {
			// Этого не должно случиться если Comparable() работает правильно
		}
	}()

	result := a == b
	return result, fmt.Sprintf("comparable type %v, used ==", aType)
}

func testEmptyInterface() {
	var data interface{} = []interface{}{
		map[string]interface{}{
			"nested": interface{}(42),
		},
	}

	path := []string{"0", "nested"}

	fmt.Println("=== Извлечение значения 42 ===")
	result, err := deepGet(data, path...)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Найдено значение: %v (тип: %T)\n", result, result)
	}

	// TODO: Извлеките значение 42 из этой структуры безопасно
	// TODO: Создайте функцию deepGet(data interface{}, path ...string) (interface{}, error)
	// TODO: Обработайте все возможные типы и nil значения
	// TODO: Используйте reflection только когда необходимо
}

func deepGet(data interface{}, path ...string) (interface{}, error) {
	var currentKey string
	var remainingPath []string
	if len(path) > 1 {
		currentKey = path[0]
		remainingPath = path[1:]
	}
	if len(path) == 1 {
		currentKey = path[0]
	}
	switch t := data.(type) {
	case int:
		fmt.Println(t)
		return t, nil
	case []any:
		index, err := parseIndex(currentKey)
		if err != nil {
			log.Fatal("error parse ", err)
		}
		return deepGet(t[index], remainingPath...)
	case map[string]any:
		return deepGet(t[currentKey], remainingPath...)
	default:
		return nil, errors.New("undefined type")
	}
}

func parseIndex(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("1")
	}

	index := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, fmt.Errorf("2")
		}
		index = index*10 + int(r-'0')
	}
	return index, nil
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string) error
}

type ReadWriter interface {
	Reader
	Writer
	Close() error
}

func testTypeSwitchEdgeCases() {
	cases := []interface{}{
		nil,
		(*int)(nil),
		(interface{})(nil),
		[]interface{}{nil},
		make(chan int),
		func() {},
	}

	for _, value := range cases {
		switch v := value.(type) {
		case nil:
			fmt.Println("Case: nil interface")
		case *int:
			if v == nil {
				fmt.Println("Case: *int (nil pointer)")
			} else {
				fmt.Printf("Case: *int with value %d\n", *v)
			}
		case []interface{}:
			fmt.Println("Case: []interface{}")
		case chan int:
			fmt.Println("Case: channel of int")
		case func():
			fmt.Println("Case: function type")
		case interface{}:
			fmt.Println("Case: interface{} containing interface")
		default:
			fmt.Printf("Case: unknown type %T\n", v)
		}
	}

}

// Тестовые данные - здесь скрыты ловушки!
func main() {
	fmt.Println("=== Тест 1: nil interface vs nil pointer ===")
	testNilInterfaces()

	fmt.Println("\n=== Тест 2: Type assertions и panic ===")
	testTypeAssertions()

	fmt.Println("\n=== Тест 3: Interface comparison ===")
	testInterfaceComparison()

	fmt.Println("\n=== Тест 4: Empty interface{} ловушки ===")
	testEmptyInterface()

	fmt.Println("\n=== Тест 5: Type switch edge cases ===")
	testTypeSwitchEdgeCases()
}
