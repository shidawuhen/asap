package design

import "fmt"

/////////////////////单一职责原则
type Game struct {
	x int64
	y int64
}

func (game *Game) Show() {
	fmt.Println(game.x, game.y)
}
func (game *Game) Move() {
	game.x--
	game.y++
}

/////////////////////开闭原则
//存储报警规则
type AlertRules struct {
}

func (alertRules *AlertRules) GetMaxTPS(api string) int64 {
	if api == "test" {
		return 10
	}
	return 100
}
func (alertRules *AlertRules) GetMaxError(api string) int64 {
	if api == "test" {
		return 10
	}
	return 100
}
func (alertRules *AlertRules) GetMaxTimeOut(api string) int64 {
	if api == "test" {
		return 10
	}
	return 100
}

const (
	SERVRE = "SERVRE"
	URGENT = "URGENT"
)

//通知类
type Notification struct {
}

func (notification *Notification) Notify(notifyLevel string) bool {
	if notifyLevel == SERVRE {
		fmt.Println("打电话")
	} else if notifyLevel == URGENT {
		fmt.Println("发短信")
	} else {
		fmt.Println("发邮件")
	}
	return true
}

//检查类
type Alert struct {
	alertRules   *AlertRules
	notification *Notification
	//版本2
	handlers []AlertHandler
}

func CreateAlert(a *AlertRules, n *Notification) *Alert {
	return &Alert{
		alertRules:   a,
		notification: n,
	}
}

//版本1
func (alert *Alert) Check(api string, tps int64, errCount int64) bool {
	if tps > alert.alertRules.GetMaxTPS(api) {
		alert.notification.Notify(URGENT)
	}
	if errCount > alert.alertRules.GetMaxError(api) {
		alert.notification.Notify(SERVRE)
	}
	return true
}

//版本2
func (alert *Alert) AddHanler(alertHandler AlertHandler) {
	alert.handlers = append(alert.handlers, alertHandler)
}
func (alert *Alert) CheckNew(apiStatInfo ApiStatInfo) bool {
	for _, h := range alert.handlers {
		h.Check(apiStatInfo)
	}
	return true
}

//优化
type ApiStatInfo struct {
	api          string
	tps          int64
	errCount     int64
	timeoutCount int64
}

type AlertHandler interface {
	Check(apiStatInfo ApiStatInfo) bool
}

type TPSAlertHandler struct {
	alertRules   *AlertRules
	notification *Notification
}

func CreateTPSAlertHandler(a *AlertRules, n *Notification) *TPSAlertHandler {
	return &TPSAlertHandler{
		alertRules:   a,
		notification: n,
	}
}

func (tPSAlertHandler *TPSAlertHandler) Check(apiStatInfo ApiStatInfo) bool {
	if apiStatInfo.tps > tPSAlertHandler.alertRules.GetMaxTPS(apiStatInfo.api) {
		tPSAlertHandler.notification.Notify(URGENT)
	}
	return true
}

type ErrAlertHandler struct {
	alertRules   *AlertRules
	notification *Notification
}

func CreateErrAlertHandler(a *AlertRules, n *Notification) *ErrAlertHandler {
	return &ErrAlertHandler{
		alertRules:   a,
		notification: n,
	}
}

func (errAlertHandler *ErrAlertHandler) Check(apiStatInfo ApiStatInfo) bool {
	if apiStatInfo.errCount > errAlertHandler.alertRules.GetMaxError(apiStatInfo.api) {
		errAlertHandler.notification.Notify(SERVRE)
	}
	return true
}

type TimeOutAlertHandler struct {
	alertRules   *AlertRules
	notification *Notification
}

func CreateTimeOutAlertHandler(a *AlertRules, n *Notification) *TimeOutAlertHandler {
	return &TimeOutAlertHandler{
		alertRules:   a,
		notification: n,
	}
}

func (timeOutAlertHandler *TimeOutAlertHandler) Check(apiStatInfo ApiStatInfo) bool {
	if apiStatInfo.timeoutCount > timeOutAlertHandler.alertRules.GetMaxTimeOut(apiStatInfo.api) {
		timeOutAlertHandler.notification.Notify(SERVRE)
	}
	return true
}

//里氏替换原则
type Notify interface {
	Send()
}
type Message struct {
}

func (message *Message) Send() {
	fmt.Println("message send")
}

type SMS struct {
}

func (sms *SMS) Send() {
	fmt.Println("sms send")
}

func LetDo(notify *Message) {
	notify.Send()
}

//接口隔离原则
type Updater interface {
	Update() bool
}

type Shower interface {
	Show() string
}

type RedisConfig struct {
}

func (redisConfig *RedisConfig) Connect() {
	fmt.Println("I am Redis")
}

func (redisConfig *RedisConfig) Update() bool {
	fmt.Println("Redis Update")
	return true
}

func (redisConfig *RedisConfig) Show() string {
	fmt.Println("Redis Show")
	return "Redis Show"
}

type MySQLConfig struct {
}

func (mySQLConfig *MySQLConfig) Connect() {
	fmt.Println("I am MySQL")
}

func (mySQLConfig *MySQLConfig) Show() string {
	fmt.Println("MySQL Show")
	return "MySQL Show"
}

type KafkaConfig struct {
}

func (kafkaConfig *KafkaConfig) Connect() {
	fmt.Println("I am Kafka")
}

func (kafkaConfig *KafkaConfig) Update() bool {
	fmt.Println("Kafka Update")
	return true
}

func ScheduleUpdater(updater Updater) bool {
	return updater.Update()
}
func ServerShow(shower Shower) string {
	return shower.Show()
}

//迪米特法则
type Transporter interface {
	Send(address string, data string) bool
}
type NetworkTransporter struct {
}

func (networkTransporter *NetworkTransporter) Send(address string, data string) bool {
	fmt.Println("NetworkTransporter Send")
	return true
}

type HtmlDownloader struct {
	transPorter Transporter
}

func CreateHtmlDownloader(t Transporter) *HtmlDownloader {
	return &HtmlDownloader{transPorter: t}
}

func (htmlDownloader *HtmlDownloader) DownloadHtml() string {
	htmlDownloader.transPorter.Send("123", "test")
	return "htmDownloader"
}

type Document struct {
	html string
}

func (document *Document) SetHtml(html string) {
	document.html = html
}

func (document *Document) Analyse() {
	fmt.Println("document analyse " + document.html)
}

func mainpriciple() {
	//开闭原则
	fmt.Println("开闭原则")
	alert := CreateAlert(new(AlertRules), new(Notification))
	alert.Check("test", 20, 20)
	//版本2，alert其实已经不需要有成员变量AlertRules和Notification了
	a := new(AlertRules)
	n := new(Notification)
	alert.AddHanler(CreateTPSAlertHandler(a, n))
	alert.AddHanler(CreateErrAlertHandler(a, n))
	alert.AddHanler(CreateTimeOutAlertHandler(a, n))
	apiStatInfo := ApiStatInfo{
		api:          "test",
		timeoutCount: 20,
		errCount:     20,
		tps:          20,
	}
	alert.CheckNew(apiStatInfo)
	//里氏替换原则
	fmt.Println("里式替换原则")
	LetDo(new(Message))
	//接口隔离原则
	fmt.Println("接口隔离原则")
	ScheduleUpdater(new(RedisConfig))
	ScheduleUpdater(new(KafkaConfig))
	ServerShow(new(RedisConfig))
	ServerShow(new(MySQLConfig))
	//迪米特法则
	fmt.Println("迪米特法则")
	htmlDownloader := CreateHtmlDownloader(new(NetworkTransporter))
	html := htmlDownloader.DownloadHtml()
	doc := new(Document)
	doc.SetHtml(html)
	doc.Analyse()
}
