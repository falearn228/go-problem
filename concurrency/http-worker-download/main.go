package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type Result struct {
	URL      string
	Filename string
	Success  bool
	Error    error
}

// Задача: Параллельная загрузка и обработка файлов

// Условие:
// Создайте программу, которая параллельно загружает HTML-страницы с нескольких URL и сохраняет их в файлы
// . Программа должна использовать goroutines и каналы для эффективной обработки.
// Требования:
//     Входные данные: Список из 5-10 URL сайтов
//     Ограничения: Максимум 3 одновременных загрузки (worker pool pattern)
//     Функциональность:
//         Скачать HTML-контент с каждого URL

//         Сохранить каждую страницу в отдельный файл

//         Вести подсчет успешных и неудачных загрузок

//         Измерить общее время выполнения

func main() {
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://stackoverflow.com",
		"https://google.com",
		"https://yandex.ru",
		"https://reddit.com",
		"https://news.ycombinator.com",
		"https://www.wikipedia.org",
		"https://www.medium.com",
		"https://www.quora.com",
	}

	wg := new(sync.WaitGroup)

	numJobs := len(urls)
	numWorkers := 3

	jobs := make(chan string, numJobs)
	results := make(chan error, numJobs)

	for i := range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i, jobs, results, download)
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for _, v := range urls {
		jobs <- v
	}
	close(jobs)

	for v := range results {
		fmt.Println("result ->", v)
	}

	fmt.Println("done!")

}

func download(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		// TODO: implement get error
		return fmt.Errorf("http get error %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status %s", resp.Status)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	filename := fmt.Sprintf("%v.txt", url)
	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("Error writing in file %v", err)
	}
	fmt.Println("done file -> %v", url)
	return nil
}

func worker(id int, jobs <-chan string, results chan<- error, method func(string) error) {
	for v := range jobs {
		fmt.Println("worker", id, "started  job", v)
		err := method(v)
		fmt.Println("worker", id, "finished  job", v)
		results <- err
	}
}
