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
    return &CookieJar{m: make(map[string]string)}
}

//--------------------------------------------------------------------------------------
/**
 * CreateCookie creates and sets the cookie 
 * Parameter Responsewriter and http.Request
 * returns the UUID as a string
 */
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
//--------------------------------------------------------------------------------------
/**
 * AddCookie function adds a cookie to internally stored map
 * Parameter Responsewriter and http.Request
 * 
 */
func (c *CookieJar) AddCookie(UUID string, name string){
	c.Lock.RLock()
	c.m[strings.Trim(UUID, "\n")] = name
	c.Lock.RUnlock()
}
//--------------------------------------------------------------------------------------
/**
 * GetValue servers as a getter to grab info from map
 * Parameter value to be taken
 * returns 2 objects, the name and a boolean value if value exist
 */
func (c *CookieJar) GetValue(value string) (string, bool) {
	name, check := c.m[value]
	return name, check

}
//--------------------------------------------------------------------------------------
/**
 * DeleteCookie deletes cookie from map 
 * Parameter the value to be deleted 
 */
func (c *CookieJar) DeleteCookie(value string){
	//cookie, _ := req.Cookie("UUID")
	_, ok := c.m[value]
	if ok {
		c.Lock.RLock()
		delete(c.m, value)
		c.Lock.RUnlock()
	}
	
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

