/**
 * @author Paul Nguyen
 * @Date: 1/31/15
 * @Name: CookieJar.go
 * @Descrption: Cookie jar that keeps track of cookie using internal stored map
 * 				
 */
package CookieJar
//imports 
import "net/http"
import "log"
import "sync"
import "os/exec"
import "strings"
import "AuthClient"

//main structure contains a mutex lock for concurrent programming and a map to keep 
//track of cookies 
type CookieJar struct {
	Lock sync.RWMutex
	m map[string]string
}
//--------------------------------------------------------------------------------------
/**
 * NewCookieJar - functions as a constructor
 * Parameter Responsewriter and http.Request
 * returns Cookie jar object
 */
func NewCookieJar() *CookieJar {
    //return &CookieJar{m: make(map[string]string)}
    return &CookieJar{m: make(map[string]string)}
}

//--------------------------------------------------------------------------------------
/**
 * CreateCookie creates and sets the cookie 
 * Parameter Responsewriter and http.Request
 * After Creating the Cookie, we add cookie to internal map storage of cookie
 */
func (c *CookieJar) CreateCookie(w http.ResponseWriter, name string) {
	value := getUniqueValue() // generate UUID
	tempValue := string(value[:]) // turn into string
	//creating cookie
	cookie := &http.Cookie{
		Name: "UUID",
		Value: strings.Trim(tempValue, "\n"),
		Path: "/",
	}
	http.SetCookie(w,cookie)
	AuthClient.Set(strings.Trim(tempValue, "\n"), name, w)
	
	
	c.Lock.RLock()
	c.m[strings.Trim(tempValue, "\n")] = name
	c.Lock.RUnlock()
	
}

//--------------------------------------------------------------------------------------
/**
 * GetValue servers as a getter to grab info from map
 * Parameter value to be taken
 * returns 2 objects, the name and a boolean value if value exist
 */
func (c *CookieJar) GetValue(value string, w http.ResponseWriter) (string, bool) {
	_, check := c.m[value]
	//check := false
	name := AuthClient.Get(value, w)
	/*
	if name != "" {
		check = true
	}
	*/
	return name, check

}
//--------------------------------------------------------------------------------------
/**
 * DeleteCookie deletes cookie from map 
 * Parameter the value to be deleted 
 */
func (c *CookieJar) DeleteCookie(w http.ResponseWriter, value string){
	//cookie, _ := req.Cookie("UUID")
	_, ok := c.m[value]
	if ok {
		c.Lock.RLock()
		delete(c.m, value)
		c.Lock.RUnlock()
	}
	cookie := &http.Cookie{
		Name: "",
		Value: "",
		Path: "/",
	}
	http.SetCookie(w,cookie)
	
}
//--------------------------------------------------------------------------------------
/**
 * GetUniqueValue uses command line to create a UUID
 * Parameter None
 * Return a UUID value as a byte array
 */
func getUniqueValue() []byte{
	out, error := exec.Command("uuidgen").Output()
	if error != nil {
		log.Fatal(error)
	}
	return out
}

func (c *CookieJar) GetCookie(req *http.Request, ID string) *http.Cookie {
	cookie, _ := req.Cookie("UUID")
	return cookie
}

