package CookieJar

import "net/http"
import "log"
import "sync"
import "os/exec"
import "strings"
/*
type cookieBox interface {
	CreateCookie() string
	AddCookie(UUID string, name string)
	DeleteCookie(value string)
	getValue(value string) (string, bool) 
}
*/
type CookieJar struct {
	Lock sync.RWMutex
	m map[string]string
}

func NewCookieJar() *CookieJar {
    return &CookieJar{m: make(map[string]string)}
}

//create the cookie and return UUID created
func CreateCookie(w http.ResponseWriter, name string) string {
	value := getUniqueValue() // generate UUID
	tempValue := string(value[:]) // turn into string
	//creating cookie
	cookie := &http.Cookie{
		Name: "UUID",
		Value: strings.Trim(tempValue, "\n"),
		Path: "/",
	}
	http.SetCookie(w,cookie)
	return tempValue
}

func (c *CookieJar) AddCookie(UUID string, name string){
	c.Lock.RLock()
	c.m[strings.Trim(UUID, "\n")] = name
	c.Lock.RUnlock()
}

func (c *CookieJar) GetValue(value string) (string, bool) {
	name, check := c.m[value]
	return name, check

}

func (c *CookieJar) DeleteCookie(value string){
	//cookie, _ := req.Cookie("UUID")
	_, ok := c.m[value]
	if ok {
		c.Lock.RLock()
		delete(c.m, value)
		c.Lock.RUnlock()
	}
	
}

func getUniqueValue() []byte{
	out, error := exec.Command("uuidgen").Output()
	if error != nil {
		log.Fatal(error)
	}
	return out
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
