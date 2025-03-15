package config

import (
	"flag"
	"os"
)

// getConfigValue - Получает значение конфигурации с приоритетом (env -> флаг -> config -> default)
func getConfigValue(envName string, flagName string, configValue *string, defaultValue string) (string, string) {
	// 1. Проверяем переменную окружения
	envVal, exists := os.LookupEnv(envName)
	if exists && envVal != "" {
		return envVal, "переменной окружения: " + envName
	}

	// 2. Проверяем флаг командной строки без повторного парсинга
	flagVal := flag.Lookup(flagName)
	if flagVal != nil && flagVal.Value.String() != "" {
		return flagVal.Value.String(), "флага командной строки"
	}

	// 3. Проверяем значение из конфигурационного файла
	if configValue != nil && *configValue != "" {
		return *configValue, "конфигурационного файла"
	}

	// 4. Если ничего не найдено, возвращаем значение по умолчанию
	return defaultValue, "значения по умолчанию"
}
