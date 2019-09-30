package controllers

import (
	"go-iris/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

/**
 *	/user/register 用户注册
 */
func (u *UserController) Register(ctx iris.Context){

	data := new(models.UserJson)
	log.Println(data)
	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(err)
		return
	}

	user, err := models.CreateUser(data)
	ctx.StatusCode(iris.StatusOK)

	if err != nil{
		ctx.JSON(ApiResult(false, err.Error(), user))
		return
	}
	ctx.JSON(ApiResult(true, "注册成功", user))
}

/**
 *	/user/login 用户登录
 */
func (u *UserController) Login(ctx iris.Context) {
	data := new(models.UserJson)

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(err)
		return
	}

	if data.Username == "" || data.Password == "" {
		ctx.JSON(ApiResult(false, "用户名或名不能为空", ""))
		return
	}

	user := new(models.User)
	if user = models.GetUserByUsername(data.Username); user.ID == 0 {
		ctx.JSON(ApiResult(false, "用户名不存在", ""))
		return
	}

	if user.Password != data.Password {
		ctx.JSON(ApiResult(false, "密码错误", ""))
		return
	}

	ctx.JSON(ApiResult(true, "登录成功", user))
}

/**
 *	/user/all 获取用户列表
 */
func (u *UserController) GetAll(ctx iris.Context){
	users := models.GetAllUsers()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(ApiResult(true, "", users))
	return
}

/**
 *	/user/username/{username:string} 根据用户名获取用户信息
 */
func (u *UserController) GetUserByUsername(username string, ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	var user *models.User
	if user = models.GetUserByUsername(username); user.ID == 0 {
		ctx.JSON(ApiResult(false, "用户不存在", ""))
		return
	}
	ctx.JSON(ApiResult(true, "", user))
}

/**
 *	/user/id/{id:int} 根据用户ID获取用户信息
 */
func (u *UserController) GetUserById(id int, ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	var user *models.User
	if user = models.GetUserById(uint(id)); user.ID == 0 {
		ctx.JSON(ApiResult(false, "用户不存在", ""))
		return
	}
	ctx.JSON(ApiResult(false, "", user))
}

/**
 *	/user/id/{id:int} 根据用户ID删除用户信息
 */
func (u *UserController) DelUserById(id int, ctx iris.Context) {
	models.DelUserById(uint(id))
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(ApiResult(false, "删除成功", ""))
}

type UserController struct {}

func (u *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/id/{id:int}", "GetUserById")
	b.Handle("POST", "/register", "Register")
	b.Handle("POST", "/login", "Login")
	b.Handle("DELETE", "/id/{id:int}", "DelUserById")
	b.Handle("GET", "/username/{username:string}", "GetUserByUsername")
	b.Handle("GET", "/all", "GetAll")
}
