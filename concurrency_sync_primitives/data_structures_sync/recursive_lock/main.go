package main

import (
	"sync"
)

/*
	Проблема кода заключается в том, что в методе Get повторно блокируется
	мьютекс, соответственно будет deadlock.

	Для этого можно реализовать отдельный метод, который будет вызываться
	только под блокировкой:

	func (c *Cache) sizeLocked() int {

		return len(c.data)
	}
*/

// Есть структура данных, которая обозначает кэш.
type Cache struct {
	mutex sync.Mutex
	data  map[string]string
}

// Конструктор для кэша.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Метод для установки ключа для кэша.
func (c *Cache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
}

// Есть метод, который получит значение из мапы по конкретному ключу.
func (c *Cache) Get(key string) string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// if c.sizeLocked() > 0 {
	// 		return c.data[key]
	// }
	if c.Size() > 0 {
		return c.data[key]
	}

	return ""
}

// Метод который возвращает размер мапы
func (c *Cache) Size() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return len(c.data)
	// return c.sizeLocked()
}
