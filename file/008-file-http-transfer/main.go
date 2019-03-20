
//多协程文件传输服务端
//作者：LvanNeo
//邮箱：lvan_software@foxmail.com
//版本：1.0
//日期：2013-09-26
//对每个请求由一个单独的协程进行处理，文件接收完成后由一个协负责将所有接收的数据合并成一个有效文件

package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		// host   = "192.168.1.5"	//如果写locahost或127.0.0.1则只能本地访问。
		port = "9090"
		// remote = host + ":" + port

		remote = ":" + port //此方式本地与非本地都可访问
	)

	fmt.Println("服务器初始化... (Ctrl-C 停止)")

	lis, err := net.Listen("tcp", remote)
	defer lis.Close()

	if err != nil {
		fmt.Println("监听端口发生错误: ", remote)
		os.Exit(-1)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("客户端连接错误: ", err.Error())
			// os.Exit(0)
			continue
		}

		//调用文件接收方法
		go receiveFile(conn)
	}
}

/*
*	文件接收方法
*	2013-09-26
*	LvanNeo
*
*	con 连接成功的客户端连接
 */
func receiveFile(con net.Conn) {
	var (
		res          string
		tempFileName string                    //保存临时文件名称
		data         = make([]byte, 1024*1024) //用于保存接收的数据的切片
		by           []byte
		databuf      = bytes.NewBuffer(by) //数据缓冲变量
		fileNum      int                   //当前协程接收的数据在原文件中的位置
	)
	defer con.Close()

	fmt.Println("新建立连接: ", con.RemoteAddr())
	j := 0 //标记接收数据的次数
	for {
		length, err := con.Read(data)
		if err != nil {

			// writeend(tempFileName, databuf.Bytes())
			da := databuf.Bytes()
			// fmt.Println("over", fileNum, len(da))
			fmt.Printf("客户端 %v 已断开. %2d %d \n", con.RemoteAddr(), fileNum, len(da))
			return
		}

		if 0 == j {

			res = string(data[0:8])
			if "fileover" == res { //判断是否为发送结束指令，且结束指令会在第一次接收的数据中
				xienum := int(data[8])
				mergeFileName := string(data[9:length])
				go mainMergeFile(xienum, mergeFileName) //合并临时文件，生成有效文件
				res = "文件接收完成: " + mergeFileName
				con.Write([]byte(res))
				fmt.Println(mergeFileName, "文件接收完成")
				return

			} else { //创建临时文件
				fileNum = int(data[0])
				tempFileName = string(data[1:length]) + strconv.Itoa(fileNum)
				fmt.Println("创建临时文件：", tempFileName)
				fout, err := os.Create(tempFileName)
				if err != nil {
					fmt.Println("创建临时文件错误", tempFileName)
					return
				}
				fout.Close()
			}
		} else {
			// databuf.Write(data[0:length])
			writeTempFileEnd(tempFileName, data[0:length])
		}

		res = strconv.Itoa(fileNum) + " 接收完成"
		con.Write([]byte(res))
		j++
	}

}

/*
*	把数据写入指定的临时文件中
*	2013-09-26
*	LvanNeo
*
*	fileName	临时文件名
*	data 		接收的数据
 */
func writeTempFileEnd(fileName string, data []byte) {
	// fmt.Println("追加：", name)
	tempFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		// panic(err)
		fmt.Println("打开临时文件错误", err)
		return
	}
	defer tempFile.Close()
	tempFile.Write(data)
}

/*
*	根据临时文件数量及有效文件名称生成文件合并规则进行文件合并
*	2013-09-26
*	LvanNeo
*
*	connumber	临时文件数量
*	filename 	有效文件名称
 */
func mainMergeFile(connumber int, filename string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建有效文件错误", err)
		return
	}
	defer file.Close()

	//依次对临时文件进行合并
	for i := 0; i < connumber; i++ {
		mergeFile(filename+strconv.Itoa(i), file)
	}

	//删除生成的临时文件
	for i := 0; i < connumber; i++ {
		os.Remove(filename + strconv.Itoa(i))
	}

}

/*
*	将指定临时文件合并到有效文件中
*	2013-09-26
*	LvanNeo
*
*	rfilename	临时文件名称
*	wfile	 	有效文件
 */
func mergeFile(rfilename string, wfile *os.File) {

	// fmt.Println(rfilename, wfilename)
	rfile, err := os.OpenFile(rfilename, os.O_RDWR, 0666)
	defer rfile.Close()
	if err != nil {
		fmt.Println("合并时打开临时文件错误:", rfilename)
		return
	}

	stat, err := rfile.Stat()
	if err != nil {
		panic(err)
	}

	num := stat.Size()

	buf := make([]byte, 1024*1024)
	for i := 0; int64(i) < num; {
		length, err := rfile.Read(buf)
		if err != nil {
			fmt.Println("读取文件错误")
		}
		i += length

		wfile.Write(buf[:length])
	}

}
//多协程文件传输客户端
//作者：LvanNeo
//邮箱：lvan_software@foxmail.com
//版本：1.0
//日期：2013-09-26
//对待发送文件进行拆分，由多个协程异步进行发送

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		host   = "192.168.1.8"     //服务端IP
		port   = "9090"            //服务端端口
		remote = host + ":" + port //构造连接串

		fileName      = "node.exe" //待发送文件名称
		mergeFileName = "mm.exe"   //待合并文件名称
		coroutine     = 10         //协程数量或拆分文件的数量
		bufsize       = 1024       //单次发送数据的大小
	)

	//获取参数信息。
	//参数顺序：
	// 1：待发送文件名
	// 2：待合并文件名
	// 3：单次发送数据大小
	// 4：协程数量或拆分文件数量
	for index, sargs := range os.Args {
		switch index {
		case 1:
			fileName = sargs
			mergeFileName = sargs
		case 2:
			mergeFileName = sargs
		case 3:
			bufsize, _ = strconv.Atoi(sargs)
		case 4:
			coroutine, _ = strconv.Atoi(sargs)
		}

	}

	fmt.Printf("请输入服务端IP: ")
	reader := bufio.NewReader(os.Stdin)
	ipdata, _, _ := reader.ReadLine()

	host = string(ipdata)
	remote = host + ":" + port

	fl, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("userFile", err)
		return
	}

	stat, err := fl.Stat() //获取文件状态
	if err != nil {
		panic(err)
	}
	var size int64
	size = stat.Size()
	fl.Close()

	littleSize := size / int64(coroutine)

	fmt.Printf("Size: %d  %d \n", size, littleSize)

	begintime := time.Now().Unix()
	//对待发送文件进行拆分计算并调用发送方法
	c := make(chan string)
	var begin int64 = 0
	for i := 0; i < coroutine; i++ {

		if i == coroutine-1 {
			go splitFile(remote, c, i, bufsize, fileName, mergeFileName, begin, size)
			fmt.Println(begin, size, bufsize)
		} else {
			go splitFile(remote, c, i, bufsize, fileName, mergeFileName, begin, begin+littleSize)
			fmt.Println(begin, begin+littleSize)
		}

		begin += littleSize
	}

	//同步等待发送文件的协程
	for j := 0; j < coroutine; j++ {
		fmt.Println(<-c)
	}

	midtime := time.Now().Unix()
	sendtime := midtime - begintime
	fmt.Printf("发送耗时：%d 分 %d 秒 \n", sendtime/60, sendtime%60)

	sendMergeCommand(remote, mergeFileName, coroutine) //发送文件合并指令及文件名
	endtime := time.Now().Unix()

	mergetime := endtime - midtime
	fmt.Printf("合并耗时：%d 分 %d 秒 \n", mergetime/60, mergetime%60)

	tot := endtime - begintime
	fmt.Printf("总计耗时：%d 分 %d 秒 \n", tot/60, tot%60)

}

/*
*	文件拆分发送方法
*	2013-09-26
*	LvanNeo
*
*	remote 服务端IP及端口号（如：192.168.1.8:9090）
*	c				channel,用于同步协程
*	coroutineNum	协程顺序或拆分文件的顺序
*	size			发送数据的大小
*	fileName		待发送的文件名
*	mergeFileName	待合并的文件名
*	begin			当前协程拆分待发送文件中的开始位置
*	end				当前协程拆分待发送文件中的结束位置
 */
func splitFile(remote string, c chan string, coroutineNum int, size int, fileName, mergeFileName string, begin int64, end int64) {

	con, err := net.Dial("tcp", remote)
	defer con.Close()
	if err != nil {
		fmt.Println("服务器连接失败.")
		os.Exit(-1)
		return
	}
	fmt.Println(coroutineNum, "连接已建立.文件发送中...")

	var by [1]byte
	by[0] = byte(coroutineNum)
	var bys []byte
	databuf := bytes.NewBuffer(bys) //数据缓冲变量
	databuf.Write(by[:])
	databuf.WriteString(mergeFileName)
	bb := databuf.Bytes()
	// bb := by[:]
	// fmt.Println(bb)
	in, err := con.Write(bb) //向服务器发送当前协程的顺序，代表拆分文件的顺序
	if err != nil {
		fmt.Printf("向服务器发送数据错误: %d\n", in)
		os.Exit(0)
	}

	var msg = make([]byte, 1024)  //创建读取服务端信息的切片
	lengthh, err := con.Read(msg) //确认服务器已收到顺序数据
	if err != nil {
		fmt.Printf("读取服务器数据错误.\n", lengthh)
		os.Exit(0)
	}
	// str := string(msg[0:lengthh])
	// fmt.Println("服务端信息：",str)

	//打开待发送文件，准备发送文件数据
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(fileName, "-文件打开错误.")
		os.Exit(0)
	}

	file.Seek(begin, 0) //设定读取文件的位置

	buf := make([]byte, size) //创建用于保存读取文件数据的切片

	var sendDtaTolNum int = 0 //记录发送成功的数据量（Byte）
	//读取并发送数据
	for i := begin; int64(i) < end; i += int64(size) {
		length, err := file.Read(buf) //读取数据到切片中
		if err != nil {
			fmt.Println("读文件错误", i, coroutineNum, end)
		}

		//判断读取的数据长度与切片的长度是否相等，如果不相等，表明文件读取已到末尾
		if length == size {
			//判断此次读取的数据是否在当前协程读取的数据范围内，如果超出，则去除多余数据，否则全部发送
			if int64(i)+int64(size) >= end {
				sendDataNum, err := con.Write(buf[:size-int((int64(i)+int64(size)-end))])
				if err != nil {
					fmt.Printf("向服务器发送数据错误: %d\n", sendDataNum)
					os.Exit(0)
				}
				sendDtaTolNum += sendDataNum
			} else {
				sendDataNum, err := con.Write(buf)
				if err != nil {
					fmt.Printf("向服务器发送数据错误: %d\n", sendDataNum)
					os.Exit(0)
				}
				sendDtaTolNum += sendDataNum
			}

		} else {
			sendDataNum, err := con.Write(buf[:length])
			if err != nil {
				fmt.Printf("向服务器发送数据错误: %d\n", sendDataNum)
				os.Exit(0)
			}
			sendDtaTolNum += sendDataNum
		}

		//读取服务器端信息，确认服务端已接收数据
		lengths, err := con.Read(msg)
		if err != nil {
			fmt.Printf("读取服务器数据错误.\n", lengths)
			os.Exit(0)
		}
		// str := string(msg[0:lengths])
		// fmt.Println("服务端信息：",str)

	}

	fmt.Println(coroutineNum, "发送数据(Byte)：", sendDtaTolNum)

	c <- strconv.Itoa(coroutineNum) + " 协程退出"
}

/*
*	向服务端发送待合并文件的名称及合并指令
*	2013-09-26
*	LvanNeo
*
*	remote 			服务端IP及端口号（如：192.168.1.8:9090）
*	mergeFileName	待合并的文件名
*	coroutine		拆分文件的总个数
 */
func sendMergeCommand(remote, mergeFileName string, coroutine int) {

	con, err := net.Dial("tcp", remote)
	defer con.Close()
	if err != nil {
		fmt.Println("服务器连接失败.")
		os.Exit(-1)
		return
	}
	fmt.Println("连接已建立. 发送合并指令.\n文件合并中...")

	var by [1]byte
	by[0] = byte(coroutine)
	var bys []byte
	databuf := bytes.NewBuffer(bys) //数据缓冲变量
	databuf.WriteString("fileover")
	databuf.Write(by[:])
	databuf.WriteString(mergeFileName)
	cmm := databuf.Bytes()

	in, err := con.Write(cmm)
	if err != nil {
		fmt.Printf("向服务器发送数据错误: %d\n", in)
	}

	var msg = make([]byte, 1024)
	lengthh, err := con.Read(msg)
	if err != nil {
		fmt.Printf("读取服务器数据错误.\n", lengthh)
		os.Exit(0)
	}
	str := string(msg[0:lengthh])
	fmt.Println("传输完成（服务端信息）： ", str)
}
