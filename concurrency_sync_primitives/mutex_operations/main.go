package main

import "sync"

// При выполнении данного кода будет deadlock,
// поскольку мы пытаемся захватить один и тот же
// mutex дважды
func lockAnyTimes() {
	mutex := sync.Mutex{}
	mutex.Lock()
	mutex.Lock()
}

// В этом коде пробуем разблокировать незаблокированный mutex.
// В этом случае будет ошибка: fatal error
func unlockWithoutLock() {
	mutex := sync.Mutex{}
	mutex.Unlock()
}

// Далее в коде сначала захватывается mutex, затем начинает
// работать WaitGroup, после чего запускается горутина,
// которая разблокирует mutex и завершает свое исполнение.
// Далее при помощи механизма WaitGroup дожидаемся завершения горутины
// и пробуем снова заблокировать, а затем разблокировать mutex.
// Программа отработает корректно, поскольку в Golang mutex-ы
// спроектированы так, что они не храянт идентификатор той
// горутины, которая его создает.
func unlockFromAnotherGoroutine() {
	mutex := sync.Mutex{}
	mutex.Lock()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		mutex.Unlock()
	}()

	wg.Wait()

	mutex.Lock()
	mutex.Unlock()
}

// Здесь происходит захват mutex-а на запись, а затем
// освобождение mutex-а на операцию чтения. Будет ошибка.
func RUnlockLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	m.RUnlock()
}

// Здесь наоборот пытаемся захватить mutex на операции чтения
// и освободить на операции записи. Будет ошибка.
func UnlockRLockedMutex() {
	m := sync.RWMutex{}
	m.RLock()
	m.Unlock()
}

// Здесь сначала пытаемся захватить mutex на операции записи,
// а затем еще раз захватить mutex на операции чтения. Будет ошибка deadlock.
func LockRLockedMutex() {
	m := sync.RWMutex{}
	m.Lock()
	m.RLock()
}

func main() {
}
