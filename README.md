<div align="center">
<img src="template/assets/images/logo.png" width="50%" />

# WHisperHArbor
_✨ 一个让大家自由倾诉的APP
</div>

### 简介
#### 程序介绍
- 是移动应用开发课程的期末大作业
- 前端是基于Android的简单的页面
- 后端本项目基于[Gin](https://github.com/gin-gonic/gin)

#### 程序理想
- 用户注册登录，每个用户都自己的uid
- 发布帖子的时候可以选择是否公开
  - 公开的话直接明文存储到`mysql`数据库中
  - 非公开则使用非对称加密进行密文存储，保证了用户的隐私。
- 首页可以浏览大家的帖子，可以进行收藏、点赞等操作
  - 因为是树洞，本人的理念是不可评论！
- 用户可以管理自己发的帖子包括修改、删除
- 用户可以查看自己的收藏等。

#### 程序进度

1. 实现了注册登录逻辑，包括
   2. 密码使用`bcrypt` 加密
   3. 使用`jwt`做验证，包括续签等操作



### 部分实现介绍

#### JWT实现
在`jwt`实现中，直接使用了[golang-jwt/jwt](https://github.com/golang-jwt/jwt)
##### 生成`Token`
可以参考`utils/CreateJWT.go`，生成标准`claim`的过期时间我设置了六小时过期，签名算法采用`HS256`
```go
const TokenExpireDuration = time.Hour * 6

var Secret = []byte("iamshitiloveeatshithhhhasdasdszcarsakjchduiashdi")

func GenerateToken(user model.LoginUser) (string, error) {
	expireDuration := time.Now().Add(TokenExpireDuration)
	claim := &model.MyClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireDuration.Unix(),
			Issuer:    "yrh",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	if tokenString, err := token.SignedString(Secret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

```

##### 校验`Token`
参考`utils/CheckJWT.go`，使用`ParseWithClaims`方法，把解析结果存到`claim`变量
```go
func ParseToken(token string) (*model.MyClaims, error) {
	claims := &model.MyClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	return claims, err
}

```

##### 续签`Token`
参考`middleware/JWTAuth.go`中间件。
```go
func RenweToken(claim *model.MyClaims) (string, error) {
	WithinLimit := func(s, l int64) bool {
		e := time.Now().Unix()
		return e-s < l
	}
	if WithinLimit(claim.ExpiresAt, 600) {
		return utils.GenerateToken(claim.User)
	}
	return "", errors.New("登录已过期")
}
```
简单来说就是，判断过期时间不超过10分钟直接续签