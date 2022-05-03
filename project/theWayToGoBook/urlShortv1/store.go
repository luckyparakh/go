package main
import "sync"

type UrlStore struct{
	urls map[string]string
	mu sync.RWMutex
}

func NewUrlStore() *UrlStore{
	return &UrlStore{urls:make(map[string]string)}
}
func (s *UrlStore) Count() int{
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *UrlStore) Get(key string) string{
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key] // Will return "" if key is not present
}

func (s *UrlStore) Set(key string,url string) bool{
	s.mu.Lock()
	defer s.mu.Unlock()
	if _,present := s.urls[key];present{
		return false
	}
	s.urls[key] = url
	return true
}

func (s *UrlStore) Put(url string) string{
	key := ""
	for {
		key = generateKey()
		if ok:=s.Set(key,url);ok{
			return key
		}
	}
	return ""
}