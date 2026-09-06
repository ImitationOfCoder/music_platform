package grpc_app

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func loggingInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	start := time.Now()

	// Логируем начало запроса
	log.Printf("Метод: %s, Запрос: %v", info.FullMethod, req)

	// Вызываем следующий обработчик (или следующий интерцептор)
	resp, err := handler(ctx, req)

	// Логируем результат
	duration := time.Since(start)
	if err != nil {
		log.Printf("Ошибка: %v, Время: %v", err, duration)
	} else {
		log.Printf("Успешно, Время: %v", duration)
	}

	return resp, err
}
