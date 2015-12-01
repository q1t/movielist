package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/q1t/movielist/db"
	"github.com/q1t/movielist/middleware"

	"html/template"

	valid "github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/lib/pq"
)

const Bearer = "Bearer"

type App struct {
	IndexTempl *template.Template
	SigningKey []byte
	E          *echo.Echo
	DB         *sqlx.DB
}

func NewApp(key []byte, dbConn *sqlx.DB) *App {
	app := &App{
		SigningKey: key,
		E:          echo.New(),
		DB:         dbConn,
	}

	// default middleware
	app.E.Use(mw.Logger())
	app.E.Use(mw.Recover())

	// serve the static assets
	app.E.Static("/public", "./public/")

	// API routes
	api := app.E.Group("/api")

	validUserInp := BindAndValidate(userData{})
	login := api.Group("/login")
	login.Use(validUserInp)
	register := api.Group("/register")
	register.Use(validUserInp)

	login.Post("", app.userLogin)
	register.Post("", app.registerNewAcc)

	// Private API routes
	auth := middleware.JWTAuth(string(app.SigningKey))
	p_api := api.Group("/private")
	p_api.Use(auth)

	user := p_api.Group("/user")
	user.Get("", app.UserEverything)
	user.Get("/lists", app.UserLists)
	user.Post("/lists/new", app.NewList)
	user.Delete("/lists/:id", app.DeleteList)

	app.E.Get("/proxy/*", app.proxyImg)

	// user routing is on the client side
	app.E.ServeFile("/*", "./public/index.html")
	return app
}

// FIXME reuse user
func (app *App) DeleteList(c *echo.Context) error {
	claims := app.claims(c)
	username := claims["username"].(string)
	user, err := FindUserByUsername(app.DB, username)
	if err != nil {
		return err
	}
	lid, _ := strconv.Atoi(c.Param("id"))
	log.Println(lid)
	if err = user.DeleteList(app.DB, int64(lid)); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

// FIXME reuse user
func (app *App) UserLists(c *echo.Context) error {
	claims := app.claims(c)
	username := claims["username"].(string)
	user, err := FindUserByUsername(app.DB, username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user.ListsAll())
}

// FIXME reuse user
func (app *App) NewList(c *echo.Context) error {
	nl := List{}
	err := c.Bind(&nl)
	if err != nil {
		return err
	}
	claims := app.claims(c)
	username := claims["username"].(string)
	user, err := FindUserByUsername(app.DB, username)
	if err != nil {
		return err
	}
	err = user.NewList(app.DB, nl)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) claims(c *echo.Context) map[string]interface{} {
	return c.Get("claims").(map[string]interface{})
}

// TODO reduce complexity
func BindAndValidate(data interface{}) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			// FIXME
			input := data.(userData)

			err := c.Bind(&input)
			if err != nil {
				err := fmt.Errorf("Struct binding error %s", err)
				return err
			}

			c.Set("userInput", input)

			var response struct {
				Errors []string `json:"errors,omitempty"`
			}

			validated, err := valid.ValidateStruct(input)

			if err != nil {
				errs := err.(valid.Errors)
				for _, v := range errs.Errors() {
					// TODO format error
					e := strings.Join(strings.Split(v.Error(), ":")[:1], "")
					errorMsg := fmt.Sprintf("%s is not valid, min length 4", e)
					response.Errors = append(response.Errors, errorMsg)
				}
				// log.Printf("Validation error: %#v", err)
			}
			if validated {
				if err := h(c); err != nil {
					c.Error(err)
				}
				return nil
			}
			return c.JSON(http.StatusOK, response)
		}
	}
}
func hashPwd(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
}

// no spaces among the validatons params
// "email, required" is not correct
// "email,required" is correct
// this struct is used in register and login handlers
type userData struct {
	Username string `valid:"stringlength(4|64),required"`
	Password string `valid:"stringlength(4|64),required"`
	Email    string `valid:"email"`
}

func (app *App) UserEverything(c *echo.Context) error {
	claims := c.Get("claims").(map[string]interface{})
	uid := int64(claims["uid"].(float64))
	user, _ := FindUserByID(app.DB, uid)
	return c.JSON(http.StatusOK, user)
}

// FIXME too many returns
func (app *App) registerNewAcc(c *echo.Context) error {
	var response struct {
		Error     string `json:"error,omitempty"`
		Timestamp int64  `json:"timestamp,omitempty"`
		Bearer    string `json:"bearer,omitempty"`
		// ID        int64  `json:"user_id,omitempty"`
	}

	// handle casting error
	userData := c.Get("userInput").(userData)
	hashedPwd, err := hashPwd(userData.Password)
	if err != nil {
		return err
	}

	user := User{
		Password: string(hashedPwd),
		Email:    userData.Email,
		Username: userData.Username,
	}

	id, err := app.newUser(user)
	user.ID = id
	if err != nil {
		if ExpectedError(err) {
			// return error to the user
			// FIXME
			response.Error = "UniqueViolation"
			return c.JSON(http.StatusOK, response)
		} else {
			// unexpected error, return code 500
			log.Println("Error while creating a user ", err)
			return err
		}
	}
	//	}
	response.Bearer = user.GenSignedString(app.SigningKey)
	response.Timestamp = time.Now().Unix()
	// response.ID = id
	return c.JSON(http.StatusOK, response)
}

func ExpectedError(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		// TODO make known errors as a map where you just can look
		// up for an existing key
		if database.UniqueViolationErr == string(err.Code) {
			return true
		}
	}
	return false
}

func (app *App) indexSinglePage(c *echo.Context) error {
	// add some info to the template
	tmpl := app.IndexTempl.Delims("[[", "]]")
	err := tmpl.Execute(c.Response(), nil)
	return err
}
func (app *App) userLogin(c *echo.Context) error {
	userData := c.Get("userInput").(userData)
	var response struct {
		Bearer    string `json:"bearer,omitempty"`
		Timestamp int64  `json:"timestamp,omitempty"`
		Error     string `json:"error,omitempty"`
	}

	// Handle credentials here
	// If everything is alright (pass is correct)
	// proceed further

	user, err := FindUserByUsername(app.DB, userData.Username)

	if err == nil {
		if err := bcrypt.CompareHashAndPassword(
			[]byte(user.Password),
			[]byte(userData.Password),
		); err == nil {
			response.Bearer = user.GenSignedString(app.SigningKey)
			response.Timestamp = time.Now().Unix()
			// user is found and password is correct
			return c.JSON(http.StatusOK, response)
		}
	}

	response.Error = "Incorrect credentials"
	return c.JSON(http.StatusOK, response)
}

func (app *App) newUser(u User) (int64, error) {
	return NewUser(app.DB, u)
}
func (app *App) GenTokenString(uid int64) string {
	token := jwt.New(jwt.SigningMethodHS256)
	duration := time.Now().Add(time.Hour * 96).Unix()

	// Set a header and a claim
	token.Header["typ"] = "JWT"
	token.Claims["exp"] = duration
	token.Claims["auth"] = true
	token.Claims["uid"] = uid

	// Generate encoded token
	t, err := token.SignedString([]byte(app.SigningKey))

	if err != nil {
		log.Printf("Error generating signed string out of token: %s", err)

		// return invalid token [temporarily]
		return ""
	}
	return t
}

// TODO: Enable cache
var cache map[string]*http.Response

func retriveFromCache(destUrl string) *http.Response {
	r, err := http.Get(destUrl)
	if err != nil {
		log.Printf("Error making req for img to proxy it: %s", err)
		return nil
	}
	return r
}

// proxy and cache images from imdb
func (app *App) proxyImg(c *echo.Context) error {
	resp := retriveFromCache(c.Param("_*"))
	data, _ := ioutil.ReadAll(resp.Body)
	for k, v := range resp.Header {
		if len(v) > 0 {
			c.Response().Header().Add(k, v[0])
		}
	}
	c.Response().Write(data)
	return nil
}
