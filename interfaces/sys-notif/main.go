package main

import (
	"fmt"
	"time"
)

// Задача: Система уведомлений с множественными каналами доставки
// Условие:
// Создайте гибкую систему уведомлений для веб-приложения. Система должна уметь отправлять уведомления через разные каналы (email, SMS, push-уведомления, Slack, Telegram) и поддерживать различные типы сообщений (информационные, предупреждения, критические ошибки).
// Требования:
//     Архитектура через интерфейсы - система должна легко расширяться новыми каналами доставки
//     Множественная доставка - одно уведомление может отправляться через несколько каналов одновременно
//     Приоритеты и фильтрация - разные каналы для разных типов сообщений
//     Статистика и мониторинг - подсчет успешных/неудачных отправок
//     Graceful degradation - если один канал недоступен, остальные продолжают работать

type NotificationType int

const (
	InfoNotification NotificationType = iota
	WarningNotification
	ErrorNotification
	CriticalNotification
)

type Notification struct {
	ID        string
	Type      NotificationType
	Title     string
	Message   string
	Recipient string
	Timestamp time.Time
	Metadata  map[string]interface{}
}

// Ваши интерфейсы и реализации здесь...
// Интерфейс для отправителей
type Sender interface {
	Send(not Notification) error
	GetName() string
	CanHandle(notType NotificationType) bool
}

type NotificationSystem interface {
	RegisterChannel(string, sender Sender)
	Send(not Notification) SendResult
	SetRoutingRules(rules RoutingRules)
	GetStatistics() Statistics
}

type SendResult struct {
	SuccessCount int
	ErrorCount   int
	Errors       map[string]error
}

type RoutingRules struct {
	Rules map[NotificationType][]string // Тип уведомления -> список каналов
}

type Statistics struct {
	TotalSent    int
	TotalError   int
	ChannelStats map[string]int
}

// Основная реализация NotificationService
type NotificationService struct {
	channels map[string]Sender // Карта зарегистрированных каналов
	rules    RoutingRules      // Правила маршрутизации
	stats    Statistics        // Статистика
}

func (ns *NotificationService) SetRoutingRules(rules RoutingRules) {
	ns.rules = rules
	fmt.Println("Правила маршрутизации установлены")
}

func (ns *NotificationService) GetStatistics() Statistics {
	return ns.stats
}

// 5. Реализации конкретных отправителей
type EmailSender struct {
	smtpServer string
}

func NewEmailSender(smtpServer string) *EmailSender {
	return &EmailSender{smtpServer: smtpServer}
}

func (e *EmailSender) Send(notification Notification) error {
	// Имитация отправки email
	fmt.Printf("📧 EMAIL [%s]: %s - %s\n", notification.Recipient, notification.Title, notification.Message)
	time.Sleep(100 * time.Millisecond) // Имитация сетевой задержки
	return nil
}

func (e *EmailSender) GetName() string {
	return fmt.Sprintf("Email Server (%s)", e.smtpServer)
}

func (e *EmailSender) CanHandle(notificationType NotificationType) bool {
	// Email может обрабатывать все типы уведомлений
	return true
}

type SMSSender struct {
	apiURL string
}

func NewSMSSender(apiURL string) *SMSSender {
	return &SMSSender{apiURL: apiURL}
}

func (s *SMSSender) Send(notification Notification) error {
	// SMS только для критических уведомлений (чтобы не спамить)
	fmt.Printf("📱 SMS [%s]: %s\n", notification.Recipient, notification.Message)
	time.Sleep(200 * time.Millisecond)
	return nil
}

func (s *SMSSender) GetName() string {
	return fmt.Sprintf("SMS Service (%s)", s.apiURL)
}

func (s *SMSSender) CanHandle(notificationType NotificationType) bool {
	// SMS только для предупреждений и критических ошибок
	return notificationType >= WarningNotification
}

type SlackSender struct {
	webhookURL string
}

func NewSlackSender(webhookURL string) *SlackSender {
	return &SlackSender{webhookURL: webhookURL}
}

func (sl *SlackSender) Send(notification Notification) error {
	emoji := map[NotificationType]string{
		InfoNotification:     "ℹ️",
		WarningNotification:  "⚠️",
		ErrorNotification:    "❌",
		CriticalNotification: "🚨",
	}

	fmt.Printf("💬 SLACK: %s %s - %s\n", emoji[notification.Type], notification.Title, notification.Message)
	time.Sleep(150 * time.Millisecond)
	return nil
}

func (sl *SlackSender) GetName() string {
	return "Slack Webhook"
}

func (sl *SlackSender) CanHandle(notificationType NotificationType) bool {
	// Slack для всех типов, кроме обычной информации
	return notificationType > InfoNotification
}

type ConsoleSender struct{}

func NewConsoleSender() *ConsoleSender {
	return &ConsoleSender{}
}

func (c *ConsoleSender) Send(notification Notification) error {
	fmt.Printf("🖥️  CONSOLE: [%v] %s - %s (to: %s)\n",
		notification.Type, notification.Title, notification.Message, notification.Recipient)
	return nil
}

func (c *ConsoleSender) GetName() string {
	return "Console Output"
}

func (c *ConsoleSender) CanHandle(notificationType NotificationType) bool {
	return true // Console может все
}

// Реализация системы
func NewNotificationService() *NotificationService {
	return &NotificationService{
		channels: make(map[string]Sender),
		rules:    RoutingRules{make(map[NotificationType][]string)},
		stats: Statistics{
			ChannelStats: make(map[string]int),
		},
	}
}

func (n *NotificationService) RegisterChannel(name string, sender Sender) {
	n.channels[name] = sender
}

func (ns *NotificationService) Send(notification Notification) SendResult {
	result := SendResult{
		Errors: make(map[string]error),
	}

	// Определяем каналы для отправки на основе правил
	channelNames := ns.getChannelsForNotification(notification.Type)

	for _, channelName := range channelNames {
		sender, exists := ns.channels[channelName]
		if !exists {
			result.ErrorCount++
			result.Errors[channelName] = fmt.Errorf("канал %s не найден", channelName)
			continue
		}

		// Проверяем, может ли канал обработать этот тип уведомления
		if !sender.CanHandle(notification.Type) {
			result.ErrorCount++
			result.Errors[channelName] = fmt.Errorf("канал %s не поддерживает тип %v", channelName, notification.Type)
			continue
		}

		// Отправляем уведомление
		err := sender.Send(notification)
		if err != nil {
			result.ErrorCount++
			result.Errors[channelName] = err
		} else {
			result.SuccessCount++
			ns.stats.ChannelStats[channelName]++
		}
	}

	ns.stats.TotalSent++
	if result.ErrorCount > 0 {
		ns.stats.TotalError++
	}

	return result
}

func (ns *NotificationService) getChannelsForNotification(notType NotificationType) []string {
	if channels, exists := ns.rules.Rules[notType]; exists {
		return channels
	}

	// Если правила не заданы, отправляем через все каналы
	var allChannels []string
	for name := range ns.channels {
		allChannels = append(allChannels, name)
	}
	return allChannels
}

func main() {
	// Создание системы
	notifier := NewNotificationService()

	// Регистрация каналов (второй аргумент - объект отправителя!)
	notifier.RegisterChannel("email", NewEmailSender("smtp.company.com"))
	notifier.RegisterChannel("sms", NewSMSSender("api.sms.com"))
	notifier.RegisterChannel("slack", NewSlackSender("webhook-url"))
	notifier.RegisterChannel("console", NewConsoleSender())

	// Настройка правил маршрутизации
	rules := RoutingRules{
		Rules: map[NotificationType][]string{
			InfoNotification:     {"email", "console"},                 // Обычная инфа
			WarningNotification:  {"email", "slack", "console"},        // Предупреждения
			ErrorNotification:    {"email", "slack", "sms", "console"}, // Ошибки
			CriticalNotification: {"email", "slack", "sms", "console"}, // Критичные
		},
	}
	notifier.SetRoutingRules(rules)

	// Тестовые уведомления
	notifications := []Notification{
		{ID: "1", Type: InfoNotification, Title: "Добро пожаловать", Message: "Регистрация завершена", Recipient: "user@example.com"},
		{ID: "2", Type: WarningNotification, Title: "Много попыток входа", Message: "Подозрительная активность", Recipient: "user@example.com"},
		{ID: "3", Type: ErrorNotification, Title: "Ошибка оплаты", Message: "Не удалось списать средства", Recipient: "user@example.com"},
		{ID: "4", Type: CriticalNotification, Title: "Система недоступна", Message: "Сервер базы данных не отвечает", Recipient: "admin@company.com"},
	}

	fmt.Println("\n=== Отправка уведомлений ===")
	for _, notification := range notifications {
		result := notifier.Send(notification)
		fmt.Printf("Уведомление %s: успешно %d, ошибок %d\n",
			notification.ID, result.SuccessCount, result.ErrorCount)
	}

	// Статистика
	fmt.Println("\n=== Статистика ===")
	stats := notifier.GetStatistics()
	fmt.Printf("Всего отправлено: %d, ошибок: %d\n", stats.TotalSent, stats.TotalError)
	for channel, count := range stats.ChannelStats {
		fmt.Printf("  %s: %d отправок\n", channel, count)
	}
}
