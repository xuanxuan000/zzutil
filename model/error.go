package model

const (
	Sucess           = 0      //OK
	RetDBErr         = -99001 //数据库错误
	RetNoDataErr     = -99002 //无数据
	RetDataExistErr  = -99003 //数据已存在
	RetDataErr       = -99004 //数据错误
	RetJsonDecodeErr = -99005 //请求数据json解码失败
	RetSessionErr    = -99006 //用户未登录
	RetLoginErr      = -99007 //用户登录失败
	RetParamErr      = -99008 //参数错误
	RetUserErr       = -99009 //用户不存在或未激活
	RetRoleErr       = -99010 //用户身份错误
	RetPwdErr        = -99011 //密码错误
	RetReqErr        = -99012 //非法请求或请求次数受限
	RetIPErr         = -99013 //IP受限
	RetThirdErr      = -99014 //第三方系统错误
	RetIOErr         = -99015 //文件读写错误
	RetServerErr     = -99016 //内部错误
	RetUknowErr      = -99017 //未知错误

)

type BlError struct {
	Err error  `json:"err"`
	Msg string `json:"msg"`
	Ret int    `json:"status"`
}

var recodeText = map[int]string{
	Sucess:           "成功",
	RetDBErr:         "数据库查询错误",
	RetNoDataErr:     "无数据",
	RetDataExistErr:  "数据已存在",
	RetDataErr:       "数据错误",
	RetJsonDecodeErr: "请求数据json解码失败",
	RetSessionErr:    "用户未登录",
	RetLoginErr:      "用户登录失败",
	RetParamErr:      "参数错误",
	RetUserErr:       "用户不存在或未激活",
	RetRoleErr:       "用户身份错误",
	RetPwdErr:        "密码错误",
	RetReqErr:        "非法请求或请求次数受限",
	RetIPErr:         "IP受限",
	RetThirdErr:      "第三方系统错误",
	RetIOErr:         "文件读写错误",
	RetServerErr:     "内部错误",
	RetUknowErr:      "未知错误",
}

func RetText(code int) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RetUknowErr]
}

func NewBlError(ret int, msg string, err error) BlError {
	return BlError{Ret: ret, Msg: msg, Err: err}
}
