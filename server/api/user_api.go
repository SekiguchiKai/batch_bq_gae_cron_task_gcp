package api

import (
	"context"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/model"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/store"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/task"
	"github.com/SekiguchiKai/batch_bq_gae_cron_task_gcp/server/util"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"net/http"
	"strconv"
	"time"
)

const (
	_UserName    = "UserName"
	_MailAddress = "MailAddress"
	_Age         = "Age"
	_Gender      = "Gender"
	_From        = "From"
)

// UserAPIを初期化する。
func InitUserAPI(g *gin.RouterGroup) {
	g.POST("/user/new", createUser)
	g.POST("user/analysis", createUserAnalyzedResult)

}

// リクエストを元にBigQueryからcsvを作成する
func createUserAnalyzedResult(c *gin.Context) {
	util.InfoLog(c.Request, "createUserAnalyzedResult is called")

	var param model.UserForAnalysis
	if err := bindUserForAnalyzeFromJson(c, &param); err != nil {
		util.RespondAndLog(c, http.StatusBadRequest, err.Error())
		return
	}

	sql, err := createSQLFromUserForAnalysis(param)
	if err != nil {
		util.RespondAndLog(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := appengine.NewContext(c.Request)
	if err := task.StartCreateCsvFromBigQuery(ctx, sql); err != nil {
		util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)

}

// リクエストで受け取ったUserをDatastoreに新たに格納する。
func createUser(c *gin.Context) {
	util.InfoLog(c.Request, "createUser is called")

	var param model.User
	if err := bindUserFromJson(c, &param); err != nil {
		util.RespondAndLog(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateParamsForUser(param); err != nil {
		util.RespondAndLog(c, http.StatusBadRequest, err.Error())
		return
	}

	u := model.NewUser(param)
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()

	util.InfoLog(c.Request, "u :%+v", u)

	err := store.RunInTransaction(c.Request, func(ctx context.Context) error {
		s := store.NewUserStoreWithContext(ctx)

		if exists, err := s.ExistsUser(u.ID); err != nil {
			util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
			return err
		} else if exists {
			util.RespondAndLog(c, http.StatusBadRequest, util.CreateErrMessage(_NotUniqueErrMessage).Error())
			return util.CreateErrMessage(_NotUniqueErrMessage)
		}

		if err := s.PutUser(u); err != nil {
			util.RespondAndLog(c, http.StatusInternalServerError, err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, nil)

}

// HTTPのリクエストボディのjsonデータUserに変換する。
func bindUserFromJson(c *gin.Context, dst *model.User) error {
	if err := c.BindJSON(dst); err != nil {
		return err
	}

	dst.ID = getUserID(c)
	return nil
}

// HTTPのリクエストボディのjsonデータUserForAnalyzeに変換する。
func bindUserForAnalyzeFromJson(c *gin.Context, dst *model.UserForAnalysis) error {
	if err := c.BindJSON(dst); err != nil {
		return err
	}

	return nil
}

// URIのIDを取得する。
func getUserID(c *gin.Context) string {
	return c.Param("id")
}

// 送信されて来たUserに必要なデータが存在するかどうかのバリデーションを行う。
func validateParamsForUser(u model.User) error {
	if u.UserName == _Empty {
		return util.CreateErrMessage(_UserName, _RequiredErrMessage)
	}

	if u.MailAddress == _Empty {
		return util.CreateErrMessage(_MailAddress, _RequiredErrMessage)
	}

	if u.Age < 0 {
		return util.CreateErrMessage(_Age, _ShouldBeOver, strconv.Itoa(0))
	}

	if u.Gender == _Empty {
		return util.CreateErrMessage(_Gender, _RequiredErrMessage)
	}

	if u.From == _Empty {
		return util.CreateErrMessage(_From, _RequiredErrMessage)
	}

	return nil

}

// UserForAnalyzeからSQLを作成する
func createSQLFromUserForAnalysis(u model.UserForAnalysis) (string, error) {

	base := `SELECT *
FROM [sandbox-sekky0905:batch_bq_task_gcp.user]
WHERE 
`

	sql := base

	if u.UserNameField.Signal != "" && u.UserNameField.Value != "" {
		sql = sql + "UserName" + _Space + u.UserNameField.Signal + _Space + u.UserNameField.Value
	}
	if u.MailAddressField.Signal != "" && u.MailAddressField.Value != "" {
		sql = sql + "MailAddress" + _Space + u.MailAddressField.Signal + _Space + _DoubleQuotation + u.MailAddressField.Value + _DoubleQuotation
	}
	if u.AgeField.Signal != "" && strconv.Itoa(u.AgeField.Value) != "" {
		sql = sql + "Age" + _Space + u.AgeField.Signal + _Space + _DoubleQuotation + strconv.Itoa(u.AgeField.Value) + _DoubleQuotation
	}
	if u.GenderField.Signal != "" && u.GenderField.Value != "" {
		sql = sql + "Gender" + _Space + u.GenderField.Signal + _Space + _DoubleQuotation + u.GenderField.Value + _DoubleQuotation
	}
	if u.FromField.Signal != "" && u.FromField.Value != "" {
		sql = sql + "From" + _Space + u.FromField.Signal + _Space + _DoubleQuotation + u.FromField.Value + _DoubleQuotation
	}
	if u.CreatedAtField.Signal != "" && u.CreatedAtField.Value != "" {
		sql = sql + "CreatedAt" + _Space + u.CreatedAtField.Signal + _Space + _DoubleQuotation + u.CreatedAtField.Value + _DoubleQuotation
	}
	if u.UpdatedAtField.Signal != "" && u.UpdatedAtField.Value != "" {
		sql = sql + "UpdatedAt" + _Space + u.UpdatedAtField.Signal + _Space + _DoubleQuotation + u.UpdatedAtField.Value + _DoubleQuotation
	}

	if sql == base {
		return "", errors.New("There is not value and signal")
	}

	return sql, nil
}
