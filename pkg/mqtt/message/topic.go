package message

const (
	/// <summary>
	/// mqtt 大屏订单通知类
	/// </summary>
	PUSH_ORDER_TO_SCREEN = "BIG_SCREEN"

	/// <summary>
	/// MQTT状态上报
	/// </summary>
	UP_MORKD_HEARTBEAT = "UP/MORKD/HEARTBEAT"

	/// <summary>
	/// 煮面机错误处理主题
	/// </summary>
	UP_MORKD_ERROR = "UP/MORKD/ERROR"

	/// <summary>
	/// 允许下发订单通知
	/// </summary>
	UP_MORKD_ALLOWPUSHORDER = "UP/MORKD/ALLOWORDER"
	/// <summary>
	/// MQTT订单上报
	/// </summary>
	UP_MORKD_ORDER = "UP/MORKD/ORDER"
	/// <summary>
	/// MQTT订单下发
	/// </summary>
	DOWN_MORKD_ORDER = "DOWN/MORKD/ORDER"

	/// <summary>
	/// MQTT业务上报
	/// </summary>
	UP_MORKD_BUSINESS = "UP/MORKD/BUSINESS"
	/// <summary>
	/// MQTT业务下发
	/// </summary>
	DOWN_MORKD_BUSINESS = "DOWN/MORKD/BUSINESS"

	/// <summary>
	/// MQTT状态上报
	/// </summary>

	UP_BQL_HEARTBEAT = "UP/BQL/HEARTBEAT"

	/// <summary>
	/// 冰淇淋错误处理主题
	/// </summary>
	UP_BQL_ERROR = "UP/BQL/ERROR"

	/// <summary>
	/// MQTT订单上报
	/// </summary>
	UP_BQL_ORDER = "UP/BQL/ORDER"
	/// <summary>
	/// MQTT订单下发
	/// </summary>
	DOWN_BQL_ORDER = "DOWN/BQL/ORDER"

	/// <summary>
	/// MQTT业务上报
	/// </summary>
	UP_BQL_BUSINESS = "UP/BQL/BUSINESS"
	/// <summary>
	/// MQTT业务下发
	/// </summary>

	DOWN_BQL_BUSINESS = "DOWN/BQL/BUSINESS"

	/// <summary>
	/// 第三代煮面机器人辅料更新状态上报
	/// </summary>
	UP_MORKS_INGREDIENT = "UP/MORKS/INGREDIENT"

	/// <summary>
	/// 第三代煮面机器人 辅料消息下发
	/// </summary>
	DOWN_MORKS_INGREDIENT = "DOWN/MORKS/INGREDIENT"

	/// <summary>
	/// MQTT状态上报
	/// </summary>
	UP_MORKS_HEARTBEAT = "UP/MORKS/HEARTBEAT"

	/// <summary>
	/// 煮面机错误处理主题
	/// </summary>
	UP_MORKS_ERROR = "UP/MORKS/ERROR"

	/// <summary>
	/// 允许下发订单通知
	/// </summary>
	UP_MORKS_ALLOWPUSHORDER = "UP/MORKS/ALLOWORDER"

	/// <summary>
	/// MQTT订单上报
	/// </summary>
	UP_MORKS_ORDER = "UP/MORKS/ORDER"

	/// <summary>
	/// MQTT订单下发
	/// </summary>
	DOWN_MORKS_ORDER = "DOWN/MORKS/ORDER"

	/// <summary>
	/// MQTT业务上报
	/// </summary>
	UP_MORKS_BUSINESS = "UP/MORKS/BUSINESS"

	/// <summary>
	/// MQTT业务下发
	/// </summary>
	DOWN_MORKS_BUSINESS = "DOWN/MORKS/BUSINESS"
)
