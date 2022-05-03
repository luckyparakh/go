package main
import (
	"sync"
	"os"
	"io"
	"encoding/gob"
)

type UrlStore struct{
	urls map[string]string
	mu sync.RWMutex
	ch chan Record
}

type Record struct{
	Key string
	URL string
}

func NewUrlStore(fileName string) *UrlStore{
	s:=&UrlStore{urls:make(map[string]string),ch:make(chan Record,100)}
	f,err:=os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err!=nil{
		panic(nil)
	}
	if err=s.load(f);err!=nil{
		// To save from reloading empty file
		if err != io.EOF{
			panic(err)
		}
	}
	wg.Add(1)
	go s.save(f)
	return s
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
			s.ch <- Record{key,url}
			// if err:=s.save(key, url);err!=nil{
			// 	panic(err)
			// }
			return key
		}
	}
	return ""
}

//func (s *UrlStore) save(key, url string) error{
func (s *UrlStore) save(f *os.File){
	for{
		r := <- s.ch
		enc := gob.NewEncoder(f)
		if err:=enc.Encode(r);err!=nil{
			panic(err)
		}
	}
	wg.Done()
}

func (s *UrlStore) load(file *os.File) error{
	dec := gob.NewDecoder(file)
	var err error
	for {
		var r Record
		if err = dec.Decode(&r); err == nil{
			s.Set(r.Key,r.URL)
		}else if err == io.EOF{
			return nil
		}else{
			return err
		}
	}
}