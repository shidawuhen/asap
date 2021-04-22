package design

import "fmt"

func testProgram() {
	t := `
<?php
//工具类
class Utils {
    //返回当前的毫秒时间戳
    public function msectime() {
        list($msec, $sec) = explode(' ', microtime());
        $msectime         = (float) sprintf('%.0f', (floatval($msec) + floatval($sec)) * 1000);
        return $msectime;
    }
    //拼接URL
    public function buildQuery($url,$arr) {
        return $url."?".http_build_query($arr);
    }
}
//Http客户端
class HttpClient {
    //http Get请求
    function sendGetRequest($url) {
        $curl = curl_init();
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($curl, CURLOPT_TIMEOUT, 5);
        curl_setopt($curl, CURLOPT_URL, $url);
        $output    = curl_exec($curl);
        curl_close($curl);
        return $output;
    }
    //http Post请求
    function sendPostRequest($url,$postData) {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $postData);
        $output = curl_exec($ch);
        curl_close($ch);
        return $output;
    }
}
//和医院系统交互类
class HospitalOperation {
    public $httpClient;
    public $hospital;
    public $utils;
    function __construct($hospital) {
        $this->httpClient = new HttpClient();
        $this->hospital = $hospital;
        $this->utils = new Utils();
    }
    //1.获取个人id
    public function getUserId() {
        var_dump("开始执行getUserId");
        $url = $this->utils->buildQuery($this->hospital->getUinfoUrl(),$this->hospital->getUinfoParams());
        $uInfo   = $this->httpClient->sendGetRequest($url);
        if(empty($uInfo)) {
            var_dump("执行getUserId失败");
            return -1;
        }
        $res = $this->hospital->getUid($uInfo);
        if(empty($res)) {
            var_dump("获取getUserId失败");
            return -1;
        }
        var_dump("getUserId成功，结果为",$res);
        $this->hospital->setRecordData($res);
        return 0;
    }
    //2.获取病人id
    public function getPatientId() {
        var_dump("开始执行getPatientId");
        $url = $this->utils->buildQuery($this->hospital->getPatientUrl(),$this->hospital->getPatientParams());
        $listInfo   = $this->httpClient->sendGetRequest($url);
        if(empty($listInfo)) {
            var_dump("执行getPatientId失败");
            return -1;
        }
        $res = $this->hospital->getPatientId($listInfo);
        if(empty($res)) {
            var_dump("获取getPatientId失败");
            return -1;
        }
        var_dump("getPatientId成功，patientId为",$res);
        $this->hospital->setRecordData($res);
        return 0;
    }
    //3.获取医生可用时间段的id
    public function getBresPeakId() {
        var_dump("开始执行getBresPeakId");
        $url = $this->utils->buildQuery($this->hospital->getBresPeakUrl(),$this->hospital->getBresPeakParams());
        $bresInfo  = $this->httpClient->sendGetRequest($url);
        if(empty($bresInfo)) {
            var_dump("执行getBresPeakId失败");
            return -1;
        }
        $res = $this->hospital->getBresPeakId($bresInfo);
        if(empty($res)) {
            var_dump("获取getBresPeakId失败");
            return -1;
        }
        var_dump("获取getBresPeakId成功",$res);
        $this->hospital->setRecordData($res);
        return 0;
    }
    //4.注册
    public function registe() {
        var_dump("开始执行registe");
        $url = $this->utils->buildQuery($this->hospital->getRegistUrl(),$this->hospital->getRegistParams());
        $postData = $this->hospital->getRegistPostData();
        $regInfo = $this->httpClient->sendPostRequest($url,$postData);
        $res = $this->hospital->getRegistRes($regInfo);
        if(empty($res)) {
            var_dump('registe failed',$postData);
        } else {
            var_dump('registe success',$postData,$res);
        }
        return $res;
    }
    //5.执行整个流程
    function run() {
        $res = $this->getUserId();
        if($res == -1) {
            exit;
        }
        $res = $this->getPatientId();
        if($res == -1) {
            exit;
        }
        $res = $this->getBresPeakId();
        if($res == -1) {
            exit;
        }
        $this->registe();
    }
}
//返回数据解码
class Decode {
    public function decodeFunc($data) {
        $data = json_decode($data, true);
        return $data;
    }
}
//医院类
class Hospital {
    protected $wxId;
    //微信ID
    protected $regDepId;
    //预约诊室
    protected $docName;
    //医生姓名
    protected $patientName;
    //患者姓名
    protected $regDay;
    //看病日期
    protected $uInfoUrl;
    //根据微信号获取个人信息URL
    protected $patientListUrl;
    //获取患者列表URL
    protected $bresPeakUrl;
    //医生出诊时间URL
    protected $regUrl;
    //注册URL
    protected $uId;
    //微信ID对应的用户ID
    protected $recordData = array();
    //注册需要提交的数据
    public $utils;
    //工具类
    public $decode;
    //接口返回数据解析方式
    function __construct($config,$regDepId,$docName,$patientName,$regDay,$decode) {
        $this->wxId = $config['wxId'];
        $this->uInfoUrl = $config['uInfoUrl'];
        $this->patientListUrl = $config['patientListUrl'];
        $this->bresPeakUrl = $config['bresPeakUrl'];
        $this->regUrl = $config['regUrl'];
        $this->regDepId = $regDepId;
        $this->docName = $docName;
        $this->patientName = $patientName;
        $this->regDay = $regDay;
        $this->decode = $decode;
        $this->utils = new Utils();
    }
    /**
     * 中间结果记录
     */
    final public function setRecordData($data) {
        foreach ($data as $key => $value) {
            $this->recordData[$key] = $value;
        }
    }
    final public function getRecordData():array {
        return $this->recordData;
    }
    /**
     * url相关
     */
    final public function getUinfoUrl():string {
        return $this->uInfoUrl;
    }
    final public function getPatientUrl():string {
        return $this->patientListUrl;
    }
    final public function getBresPeakUrl():string {
        return $this->bresPeakUrl;
    }
    final public function getRegistUrl():string {
        return $this->regUrl;
    }
    public function getUinfoParams():array {
    }
    //子类需重写
    public function getPatientParams():array {
    }
    //子类需重写
    public function getBresPeakParams():array {
    }
    //子类需重写
    public function getRegistParams():array {
    }
    //子类需重写
    public function getRegistPostData():array {
    }
    //子类需重写
    /**
     * 获得微信对应的用户id，空表示获取失败
     * 子类需重写
     */
    public function getUid($res):array {
    }
    /**
     * 获得患者Id，空表示获取失败
     * 子类需重写
     */
    public function getPatientId($res):array {
    }
    /**
     * 获得医生问诊时段Id，空表示获取失败
     * 子类需重写
     */
    public function getBresPeakId($res):array {
    }
    /**
     * 获得注册结果，空表示注册失败
     * 子类需重写
     */
    public function getRegistRes($res):array {
    }
}
/**
 * Class HospitalA
 * 继承自Hospital类，主要用于
 * 1. 生成各个接口的URL
 * 2. 解析返回数据，获取想要的结果
 * 3. 最终生成注册的数据postData
 */
class HospitalA extends Hospital {
    /**
     * url请求参数相关
     */
    public function getUinfoParams():array {
        return array(
            'act'=>'userinfo_oid',
            'uid'=>$this->wxId,
            'tm'=>$this->utils->msectime(),
            'oid'=>$this->wxId,
        );
    }
    public function getPatientParams():array {
        $recordData = $this->getRecordData();
        return array(
            'act'=>'member',
            'uid'=>$recordData['uid'],
            'oid'=>$this->wxId,
        );
    }
    public function getBresPeakParams():array {
        return array(
            'act'=>'bespeak_v1',
            'deptid'=>$this->regDepId,
            'clsid' => 2, //专家号
            'tm'=>$this->utils->msectime(),
        );
    }
    public function getRegistParams():array {
        return array(
            'act'=>'bespeak',
        );
    }
    public function getRegistPostData():array {
        $recordData = $this->getRecordData();
        return array(
            'oid' => $this->wxId,
            'uid' => $recordData['uid'],
            'sickid' => $recordData['sickid'],
            'bespeakid' => $recordData['bespeakid'],
            'aorp' => $recordData['aorp'],
        );
    }
    /**
     * @param $res
     * @return array
     * 根据返回结果获取微信的uid
     */
    public function getUid($res):array {
        $res = $this->decode->decodeFunc($res);
        if($res['result'] != 'ok') {
            return array();
        }
        $uInfo = $this->decode->decodeFunc($res['data']);
        return array(
            'uid' => $uInfo['id'],
        );
    }
    /**
     * @param $res
     * @return array
     * 根据返回结果获取患者id
     */
    public function getPatientId($res):array {
        $res = $this->decode->decodeFunc($res);
        if($res['result'] != 'ok') {
            return array();
        }
        $list = $this->decode->decodeFunc($res['data']);
        foreach ($list as $item) {
            if ($item['name'] == $this->patientName) {
                return array(
                    'sickid' => $item['id'],
                );
            }
        }
        return array();
    }
    /**
     * @param $res
     * @return array
     * 根据返回结果获取可用的问诊时段
     */
    public function getBresPeakId($res):array {
        $res = $this->decode->decodeFunc($res);
        if($res['result'] != 'ok') {
            return array();
        }
        $docList = $this->decode->decodeFunc($res['data']);
        $bespeakid = -1;
        $aorp      = 0;
        //0上午 1下午
        $flag = 0;
        //var_dump($docList);
        foreach ($docList as $item) {
            if ($item['name'] == $this->docName && $item['bdate'] == $this->regDay) {
                if ($item['pm'] != 0 && $item['pm'] != '约满') {
                    //下午有号
                    $aorp      = 1;
                    $flag = 1;
                } else if ($item['am'] != 0 && $item['am'] != '约满') {
                    //上午有号
                    $aorp      = 0;
                    $flag = 1;
                }
                if($flag == 1) {
                    $bespeakid = (int)($item['id']);
                    var_dump('选择医生为',$item);
                    break;
                }
            }
        }
        if($bespeakid != -1) {
            return array(
                'bespeakid' => $bespeakid,
                'aorp' => $aorp,
            );
        }
        return array();
    }
    /**
     * @param $res
     * @return array
     * 根据返回结果判断注册是否成功
     */
    public function getRegistRes($res):array {
        $res = $this->decode->decodeFunc($res);
        var_dump($res);
        if($res['result'] != 'ok') {
            return array();
        }
        return array(
            'res' => 'success',
        );
    }
}
class HosiptalFactory {
    public function createHospital($hospitalName,$config,$regDepId,$docName,$patientName,$regDay,$decode) {
        $hospital = null;
        switch ($hospitalName) {
            case 'A': $hospital = new HospitalA($config['A'],$regDepId,$docName,$patientName,$regDay,$decode);
        }
        return $hospital;
    }
}
function main() {
    $config = array(
        'A' => array(//1号配置
            'wxId' => '**',
            'uInfoUrl' => '**',
            'patientListUrl' => '**',
            'bresPeakUrl' => '**',
            'regUrl' => '**',
        ),
    );
    $decode = new Decode();
    $docName     = '*';
    $patientName = '*';
    $regDepId    = *;
    $regDay      = '2021-03-26';
    $hospitalName = 'A';
    //使用工厂模式创建对应医院
    $factory = new HosiptalFactory();
    $hospital = $factory->createHospital($hospitalName,$config,$regDepId,$docName,$patientName,$regDay,$decode);
    $oper = new HospitalOperation($hospital);
    $oper->run();
}
main();
`
	fmt.Println(t)
	return
}
