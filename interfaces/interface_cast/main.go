package main

type BaseStorageImpl struct{}

func (s BaseStorageImpl) Close() {}

type SyncStorageImpl struct{}

func (s SyncStorageImpl) Close() {}
func (s SyncStorageImpl) Sync()  {}

type BaseStorage interface {
	Close()
}

type SyncStorage interface {
	Close()
	Sync()
}

func main() {
	var baseStorage BaseStorage = BaseStorageImpl{}
	var syncStorage SyncStorage = SyncStorageImpl{}

	println("baseStorage:", baseStorage)
	// Здесь присвоение интерфейса syncStorage в baseStorage
	// отработает корректно, поскольку у syncStorage реализован
	// метод Close.
	baseStorage = syncStorage
	println("baseStorage:", baseStorage)

	println("syncStorage:", syncStorage)
	// Аналогично здесь можно присваивать интерфейсу syncStorageCasted
	// интерфейс syncStorage только с одним методом Close().
	syncStorageCasted := syncStorage.(interface{ Close() })
	println("syncStorageCasted:", syncStorageCasted)

	// Это работать не будет, поскольку syncStorage требует
	// методов Close() и Sync().
	//syncStorage = baseStorage
}
