package main

import (
	"html/template"
	"io"
	"kys/services"
	"log"
	"reflect"
	"github.com/labstack/echo/v4"
)

func FirstElement(list []string) string {
  if len(list) > 0 {
    return list[0]
  }
  return ""
}

func GetHumanTags(s any) map[string]interface{} {
  mappedData := make(map[string]interface{})
  t := reflect.TypeOf(s)
  v := reflect.ValueOf(s)
  for i := 0; i < t.NumField(); i++ {
    fld := t.Field(i)
    humanTag := fld.Tag.Get("human")
    if humanTag != ""{
      mappedData[humanTag] = v.Field(i).Interface()
    }
  }
  return mappedData
}

func main(){
  log.Println("Running server")
	e := echo.New()
	e.Renderer = newTemplate()

	e.Static("/public", "public")

	e.GET("/", func(c echo.Context) error {
      result, err := db.QueryEntity("SELECT id, caption, schema, properties, image FROM test LIMIT 10;")
      if err != nil {
        log.Printf("ERROR")
        return c.JSON(200,err)
    }
		return c.Render(200, "index", result)
	})

  e.GET("/data", func(c echo.Context) error {
      result, err := db.QueryEntity("SELECT id, caption, schema, properties, image FROM test;")
      if err != nil {
        return c.JSON(200, err)
      }
			return c.JSON(200, result)

	})

	e.GET("/entities", func(c echo.Context) error {
    result, err := db.QueryEntity("SELECT id, caption, schema, properties, image FROM test;")
      if err != nil {
        return c.JSON(200, "DB down")
    }
			return c.Render(200, "entities", result)
	})


  e.GET("/entities/:entity", func(c echo.Context) error {
		query := c.Param("entity")
    result, err := db.QueryEntity("SELECT id, caption, schema, properties, image FROM test;")
      if err != nil {
        return c.JSON(200, "DB down")
    }
    for _, v:= range result {
      if v.Id == query {
        return c.Render(200, "profile", v)
      }
    }
	return c.Render(200, "", nil)
	})

  e.GET("/about", func(c echo.Context) error {
			return c.Render(200, "about", nil)

	})
	e.Start(":8080")

}


type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates{
  funcMap := template.FuncMap{
		"firstElement": FirstElement,
    "getHumanTags": GetHumanTags,
	}
	return &Templates{
			templates: template.Must(template.New("").Funcs(funcMap).ParseGlob("*views/*.html")),
	}
}


