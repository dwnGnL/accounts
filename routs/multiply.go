package routs

import(
	"github.com/gin-gonic/gin"
	"account/models"
	"github.com/jinzhu/gorm"
	"net/http"
	"crypto/md5"
	"time"
	"fmt"
	"io"
	eb64 "encoding/base64"
	"log"
)
var account *models.AccountServices
var Dbs *gorm.DB
var sql_account []models.Account
func Init(c *gin.Context){
	var check bool
	var number int
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Cannot decode " + err.Error()})
		log.Panic(err.Error())
		return
	}
	Dbs.Find(&sql_account)
	checks:=false
	for i,row:= range sql_account{
		if account.Login==row.Fname{
			hash:=md5.New()
			dates:=row.Fname+":"+row.ApiKey
			io.WriteString(hash,dates)
			if account.Password == eb64.RawStdEncoding.EncodeToString(hash.Sum(nil)){
				check=true
				number=i
			}else{
				c.JSON(http.StatusBadGateway,gin.H{"message":"error password"})
				return
			}
			checks=true
		}
	}
	if !checks{
		c.JSON(http.StatusBadGateway,gin.H{"message":"error login"})
			return
	}
	if check{
		switch account.Action.Type {
		case 0:multi(c,number)
		case 1:snyat(c,number)
		case 2:checkBalance(c, number)
		case 3:checkOperation(c,number)
		default:c.JSON(http.StatusBadGateway,gin.H{"message":"error type"})
		}
	}
}

func multi(c *gin.Context,i int){
	var id int
	var operation models.Operation
	for _,sum:=range account.Action.Services{
		sql_account[i].Balance+=sum.Sum
		operation.LastSum=sum.Sum
		fmt.Println(operation.LastSum)
		operation.Operations="Пополнение"
			now := time.Now()
			timef := now.Format("02.01.2006 15:04:05")
		operation.Date=timef
		operation.Name=sum.Name
		operation.Id_account=sql_account[i].ID
		Dbs.Create(&operation)
	}
	id=sql_account[i].ID
	log.Println(sql_account[i].Balance)
	Dbs.Model(&sql_account).Where("ID = ?",id).Update("Balance",sql_account[i].Balance)
	c.JSON(http.StatusOK,gin.H{"message":"All okey"})
}

func snyat(c *gin.Context,i int){
	var operation models.Operation
	var id int
	for _,sum:=range account.Action.Services{
		if sql_account[i].Balance>=sum.Sum{
			sql_account[i].Balance-=sum.Sum
			operation.LastSum=sum.Sum
			fmt.Println(operation.LastSum)
			operation.Operations="Cнятие"
				now := time.Now()
				timef := now.Format("02.01.2006 15:04:05")
			operation.Date=timef
			operation.Name=sum.Name
			operation.Id_account=sql_account[i].ID
			Dbs.Create(&operation)
		}else{
			c.JSON(http.StatusBadGateway,gin.H{"message":"you don`t have money"})
			return
		}
	}
	id=sql_account[i].ID
	log.Println(sql_account[i].Balance)
	Dbs.Model(&sql_account).Where("ID = ?",id).Update("Balance",sql_account[i].Balance)

	c.JSON(http.StatusOK,gin.H{"message":"All okey"})

}

func checkBalance(c *gin.Context,i int){
	c.JSON(http.StatusOK,sql_account[i])
}

func checkOperation(c *gin.Context,i int){
	var operation []models.Operation
	Dbs.Where("id_account = ?", sql_account[i].ID).Find(&operation)
	c.JSON(http.StatusOK,operation)
}