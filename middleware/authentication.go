package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

type Session struct {
	Uid       int
	Role      string
	SessionId string
}
type Backup struct {
	Sessions []Session
}
type JwtClaims struct {
	Uid  string
	Role string
	jwt.RegisteredClaims
}

var SessionInfo Backup

const dsn = "root:admin123@tcp(127.0.0.1:3306)/goLearn?multiStatements=true"

func CreateSession(ctx *gin.Context) {
	// return func(ctx *gin.Context) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	var credentials struct {
		UserId   int    `json:"userId"`
		Password string `json:"password"`
	}
	claim := &JwtClaims{}

	if err := ctx.Bind(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	res, err := db.Query("SELECT grNo, userRole FROM students WHERE grNo=? AND sPwd=? ", credentials.UserId, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	temptime := time.Now().Add(24 * time.Hour)
	fmt.Println("time added 1 day", temptime)
	if res.Next() {
		var grNo int
		var role string
		err = res.Scan(&grNo, &role)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		claim.Uid = strconv.Itoa(grNo)
		claim.Role = role
		claim.RegisteredClaims.IssuedAt = jwt.NewNumericDate(time.Now())
		claim.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	} else if !res.Next() {
		res, err = db.Query("SELECT tId, userRole FROM teachers WHERE tId='?' AND tPwd='?' ", credentials.UserId, credentials.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		if res.Next() {
			var tId string
			var role string
			err = res.Scan(&tId, &role)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			claim.Uid = tId
			claim.Role = role
			claim.RegisteredClaims.IssuedAt = jwt.NewNumericDate(time.Now())
			claim.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(temptime)
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid credentials"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte("9tvfPMwMVQHdksYp"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	if _, err = db.Exec("INSERT INTO activeSessions (sessiontoken, userRole, validtime) VALUES (?,?,?)", tokenString, claim.Role, temptime); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process session"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})

	fmt.Println(claim)
	fmt.Println(ctx.Cookie("usercookie"))
	// }
}

func ValidateSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if userCookie, err := ctx.Cookie("userCookie"); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "token not found"})
			return
		} else {
			fmt.Println(userCookie)
			claim := &JwtClaims{}

			token, err := jwt.ParseWithClaims(userCookie, claim, func(t *jwt.Token) (any, error) {
				return []byte("9tvfPMwMVQHdksYp"), nil
			})

			if err != nil || !token.Valid {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
				return
			}

			ctx.Set("userrole", claim.Role)
			ctx.Next()
		}
	}

}
