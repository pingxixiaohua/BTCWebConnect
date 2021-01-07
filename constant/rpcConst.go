package constant

/**
 * 比特币节点的命令
 */
//关闭比特币节点
const STOP = "stop"
//获取比特币节点的区块数量
const GETBLOCKCOUNT = "getblockcount"
//获取节点的最新区块的hash值
const GETBESTBLOCKHASH  = "getbestblockhash"
//根据hash值获取区块的具体内容
const GETBLOCK = "getblock"
//生成新的比特币地址
const GETNEWADDRESS = "getnewaddress"
//根据高度获取某个区块的hash
const GETBLOCKHASH = "getblockhash"
//获取区块链信息
const GETBLOCKCHAININFO = "getblockchaininfo"
//获取当前比特币网络中的区块难度
const GETDIFFICULTY = "getdifficulty"
//获取钱包的信息
const GETWALLETINFO = "getwalletinfo"



//获取当前服务器运行的秒数
const GETBLOCKSERVERTIME = "uptime"
//设置交易比特币的费用每kb
const GETSETTXFEE = "settxfee"
//设置与地址关联的标签
const GETSETLABEL = "setlabel"
//停止由RPC调用触发的当前钱包重新扫描。
const GETABORTRESCAN = "abortrescan"
//验证已签名的消息
const GETVERIFYMESSAGE = "verifymessage"
//请求向所有其他节点发送一个ping，以测量ping时间
const GETPING = "ping"