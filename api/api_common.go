package api

import "github.com/gin-gonic/gin"

const (
	UrlAdminLogin         = "post:admin/login"
	UrlAdminGetSelf       = "get@admin:admin/self"
	UrlAdminGetList       = "get@admin:admin"
	UrlAdminGetMenu       = "get@admin:admin/menu"
	UrlAdminUpdate        = "put@admin:admin"
	UrlAdminCreate        = "post@admin:admin"
	UrlAdminDelete        = "delete@admin:admin"
	UrlAdminUpdatePwd     = "post@admin:admin/pwd"
	UrlAdminUpdatePwdSelf = "post@admin:admin/pwd_self"

	UrlWxLogin      = "get:wxLogin"
	UrlWxAdminLogin = "get:wxAdminLogin"
	UrlWxLocation   = "get:wxLocation"

	UrlUserGet               = "get:user/one"
	UrlUserGetByICCardClient = "get@client:user/one_by_iccard"
	UrlUserGetByICCard       = "get@admin:user/admin_one_by_iccard"
	UrlUserGetWxSelf         = "get@user:user/wx_self"
	UrlUserGetSelf           = "get@user:user/self"
	UrlUserGetList           = "get@admin:user"
	UrlUserDelete            = "delete@admin:user"
	UrlUserCreate            = "post@admin:user"
	UrlUserUpdateByAdmin     = "put@admin:user"
	UrlUserUpdateBySelf      = "put@user:user/info"
	UrlUserInBlack           = "post@admin:user/black"
	UrlUserBindPhone         = "post@user:user/bind_phone"

	UrlSmsSendBindPhone = "post@user:sms/bind_phone"

	UrlProductGetList = "get@admin:product"
	UrlProductCreate  = "post@admin:product"
	UrlProductUpdate  = "put@admin:product"
	UrlProductDel     = "delete@admin:product"

	UrlProductPayCreate       = "post@admin:product_pay"
	UrlProductPayUpdateStatus = "post@admin:product_pay/status"
	UrlProductPayGet          = "get@admin:product_pay"
	UrlProductPayGetByUser    = "get@user:product_pay_by_user"

	UrlBoxGetList       = "get@admin:box"
	UrlBoxGet           = "get@admin:box/one"
	UrlBoxGetByUser     = "get@user:box/one_by_user"
	UrlBoxGetOpenStatus = "get@user:box/open_status"
	UrlBoxOpen          = "post@admin:box/open"
	UrlBoxOpenByUser    = "post@user:box/open_by_user"
	UrlBoxGetListByUser = "get@user:box/by_user"
	UrlBoxUpdate        = "put@admin:box"
	UrlBoxDel           = "delete@admin:box"

	UrlAppointmentGetListByApp      = "get@admin:appointment/by_app"
	UrlAppointmentGetListByAdminGot = "get@admin:appointment/admin_got"
	UrlAppointmentGetListByUser     = "get@user:appointment/by_user"
	UrlAppointmentGet               = "get@admin:appointment/one"
	UrlAppointmentCreate            = "post@user:appointment"
	UrlAppointmentDel               = "delete@admin:appointment"
	UrlAppointmentCancel            = "post@user:appointment/cancel"
	UrlAppointmentClose             = "post@admin:appointment/close"
	UrlAppointmentGoing             = "post@admin:appointment/going"
	UrlAppointmentFinish            = "post@admin:appointment/finish"

	UrlWithdrawLogGetListByApp  = "get@admin:withdraw_log/app"
	UrlWithdrawLogGetListByUser = "get@user:withdraw_log/user"
	UrlWithdrawLogCreate        = "post@user:withdraw_log"
	UrlWithdrawLogDel           = "delete@admin:withdraw_log"
	UrlWithdrawLogAllow         = "post@admin:withdraw_log/allow"
	UrlWithdrawLogRefuse        = "post@admin:withdraw_log/refuse"

	UrlMsgBoardGetList = "get@admin:msg_board"
	UrlMsgBoardDel     = "delete@admin:msg_board"
	UrlMsgBoardUpdate  = "put@user:msg_board"
	UrlMsgBoardCreate  = "post@user:msg_board"

	UrlPointLogListByUser      = "get@user:point_log/by_user"
	UrlPointLogListByAdmin     = "get@admin:point_log/admin"
	UrlPointLogDonationByUser  = "post@user:point_log/donation_by_user"
	UrlPointLogDonationByAdmin = "post@admin:point_log/donation_by_admin"
	UrlPointLogExchangeByUser  = "post@user:point_log/exchange_by_user"
	UrlPointLogExchangeByAdmin = "post@admin:point_log/exchange_by_admin"
	UrlPointLogExchangeQrcode  = "post@user:point_log/exchange/qrcode"

	UrlPushLogListByUser = "get@user:push_log/by_user"
	UrlPushLogLastByUser = "get@user:push_log/last_by_user"
	UrlPushLogListByApp  = "get@admin:push_log/app"
	UrlPushLogCreate     = "post@user:push_log"
	UrlPushLogMark       = "post@admin:push_log/mark"

	UrlAppSetting            = "post@admin:app/setting"
	UrlAppResetDonationCount = "post@admin:app/reset_donationcount"
	UrlAppResetDonationPoint = "post@admin:app/reset_donationpoint"
	UrlAppGet                = "get@admin:app"
	UrlAppGetByUser          = "get@user:app/by_user"
	UrlAppGetByClient        = "get@client:app/by_client"

	UrlGetPointByMoney  = "get:app/point"
	UrlGetMoneyByPoint  = "get:app/money"
	UrlGetPointByWeight = "get:app/point/by_weight"

	UrlToolMqPublish = "get:tool/mq_publish"
	UrlUpload        = "post:upload"
)

func GetUid(ctx *gin.Context) int64 {
	return ctx.GetInt64("uid")
}

func DataAndTotal(data interface{}, total int64) interface{} {
	return map[string]interface{}{
		"total": total,
		"data":  data,
	}
}
