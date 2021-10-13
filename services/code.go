package services

type statusCode int

const (
	UserCreateSucceed statusCode = iota
	UserUpdateSucceed
	UserDeleteSucceed
	UserNotFound
	UserHasExist
	UserCenterExist
	UserIPError
	UserIPDuplicate
	ServerCreateSucceed
	ServerUpdateSucceed
	ServerDeleteSucceed
	ServerNotFound
	ServerTitleRequired
	ServerHasExist
	ServerIPError
	RuleCreateSucceed
	RuleUpdateSucceed
	RuleDeleteSucceed
	RuleNotFound
	RuleHasExist
	RuleIPError
	RuleMapCreateSucceed
	RuleMapUpdateSucceed
	RuleMapDeleteSucceed
	RuleMapNotFound
	RuleMapHasExist
)

func (code statusCode) String() string {
	switch code {
	case UserCreateSucceed:
		return "Create User Succeed"
	case UserUpdateSucceed:
		return "Update User Succeed"
	case UserDeleteSucceed:
		return "Delete User Succeed"
	case UserNotFound:
		return "User Not Found"
	case UserHasExist:
		return "User Has Exist"
	case UserCenterExist:
		return "User Center Has Exist"
	case UserIPError:
		return "User IP Error"
	case UserIPDuplicate:
		return "User IP Duplicate"
	case ServerCreateSucceed:
		return "Create Server Succeed"
	case ServerUpdateSucceed:
		return "Update Server Succeed"
	case ServerDeleteSucceed:
		return "Delete Server Succeed"
	case ServerNotFound:
		return "Server Not Found"
	case ServerTitleRequired:
		return "Server Title Required"
	case ServerHasExist:
		return "Server Has Exist"
	case ServerIPError:
		return "Server IP Error"
	case RuleCreateSucceed:
		return "Create Rule Succeed"
	case RuleUpdateSucceed:
		return "Update Rule Succeed"
	case RuleDeleteSucceed:
		return "Delete Rule Succeed"
	case RuleNotFound:
		return "Rule Not Found"
	case RuleHasExist:
		return "Rule Has Exist"
	case RuleIPError:
		return "Rule IP Error: IP/Netmask"
	case RuleMapCreateSucceed:
		return "Create RuleMap Succeed"
	case RuleMapUpdateSucceed:
		return "Update RuleMap Succeed"
	case RuleMapDeleteSucceed:
		return "Delete RuleMap Succeed"
	case RuleMapNotFound:
		return "RuleMap Not Found"
	case RuleMapHasExist:
		return "RuleMap Has Exist"
	}
	return "codeError"
}
