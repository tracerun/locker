package locker

import "sync"

type counter struct {
	count int
	lock  *sync.RWMutex
}

// Locker to lock certain resource
type Locker struct {
	info map[string]*counter
	lock *sync.Mutex
}

// New to create a locker instance
func New() *Locker {
	return &Locker{
		info: make(map[string]*counter),
		lock: new(sync.Mutex),
	}
}

// ReadLock to the resource
// return a release function that needs to be called after finishing
func (l *Locker) ReadLock(resource string) func() {
	l.lock.Lock()
	one := l.info[resource]
	if one == nil {
		one = &counter{count: 1, lock: new(sync.RWMutex)}
		l.info[resource] = one
	} else {
		one.count++
	}
	l.lock.Unlock()

	one.lock.RLock()

	return func() {
		one.lock.RUnlock()

		l.lock.Lock()
		one := l.info[resource]
		if one.count <= 1 {
			delete(l.info, resource)
		} else {
			one.count--
		}
		l.lock.Unlock()
	}
}

// WriteLock to the resource
// return a release function that needs to be called after finishing
func (l *Locker) WriteLock(resource string) func() {
	l.lock.Lock()
	one := l.info[resource]
	if one == nil {
		one = &counter{count: 1, lock: new(sync.RWMutex)}
		l.info[resource] = one
	} else {
		one.count++
	}
	l.lock.Unlock()

	one.lock.Lock()

	return func() {
		one.lock.Unlock()

		l.lock.Lock()
		one := l.info[resource]
		if one.count <= 1 {
			delete(l.info, resource)
		} else {
			one.count--
		}
		l.lock.Unlock()
	}
}
