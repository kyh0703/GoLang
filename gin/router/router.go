package router

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func GetSomeJSON(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "Go",
		"tag":  "<br>",
	}
	c.AsciiJSON(http.StatusOK, data)
}

func GetBinding(c *gin.Context) {
	type StructD struct {
		NestedAnonyStruct struct {
			FieldX string `form:"field_x"`
		}
		FieldD string `form:"field_d"`
	}

	var b StructD
	c.Bind(&b)
	log.Println(b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func FromHandler(c *gin.Context) {
	type myForm struct {
		Colors []string `form:"colors[]"`
	}
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

// If `GET`, only `Form` binding engine (`query`) used.
// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
func BindQuery(c *gin.Context) {
	type Person struct {
		Name     string    `form:"name"`
		Address  string    `form:"address"`
		Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	}
	var person Person
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	c.String(200, "Success")
}

// https://gin-gonic.com/docs/examples/bind-uri/
// curl -v http://localhost:8080/test-uri/test/987fbc97-4bed-5078-9f07-9141ba07c9f3
func BindUri(c *gin.Context) {
	type Person struct {
		ID   string `uri:"id" binding:"required,uuid"`
		Name string `uri:"name" binding:"required"`
	}
	var person Person
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}
