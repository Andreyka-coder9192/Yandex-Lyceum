package task110

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}

type OrderLogger struct {
	out *bytes.Buffer
}

func NewOrderLogger() *OrderLogger {
	return &OrderLogger{out: &bytes.Buffer{}}
}

func (logger *OrderLogger) AddOrder(order Order) {
	log.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)
}

func TestOrderLogger(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	logger := NewOrderLogger()

	order := Order{1, "Иванов", 100.50}
	logger.AddOrder(order)

	expectedLog := fmt.Sprintf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)
	actualLog := buf.String()

	if actualLog[20:] != expectedLog {
		t.Errorf("Лог не соответствует ожиданиям. Ожидалось:\n%s\nПолучено:\n%s\n", expectedLog, actualLog)
	}
}
