// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"nusadua/rpc"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void multi_push( keys,  values)")
	fmt.Fprintln(os.Stderr, "   multi_pull( keys)")
	fmt.Fprintln(os.Stderr, "  void range_push(i64 start_key, i64 end_key,  values)")
	fmt.Fprintln(os.Stderr, "   range_pull(i64 start_key, i64 end_key)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = math.MinInt32 // will become unneeded eventually
	_ = strconv.Atoi
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := rpc.NewServerClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "multi_push":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "MultiPush requires 2 args")
			flag.Usage()
		}
		arg51 := flag.Arg(1)
		mbTrans52 := thrift.NewTMemoryBufferLen(len(arg51))
		defer mbTrans52.Close()
		_, err53 := mbTrans52.WriteString(arg51)
		if err53 != nil {
			Usage()
			return
		}
		factory54 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt55 := factory54.GetProtocol(mbTrans52)
		containerStruct0 := rpc.NewMultiPushArgs()
		err56 := containerStruct0.ReadField1(jsProt55)
		if err56 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Keys
		value0 := argvalue0
		arg57 := flag.Arg(2)
		mbTrans58 := thrift.NewTMemoryBufferLen(len(arg57))
		defer mbTrans58.Close()
		_, err59 := mbTrans58.WriteString(arg57)
		if err59 != nil {
			Usage()
			return
		}
		factory60 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt61 := factory60.GetProtocol(mbTrans58)
		containerStruct1 := rpc.NewMultiPushArgs()
		err62 := containerStruct1.ReadField2(jsProt61)
		if err62 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Values
		value1 := argvalue1
		fmt.Print(client.MultiPush(value0, value1))
		fmt.Print("\n")
		break
	case "multi_pull":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "MultiPull requires 1 args")
			flag.Usage()
		}
		arg63 := flag.Arg(1)
		mbTrans64 := thrift.NewTMemoryBufferLen(len(arg63))
		defer mbTrans64.Close()
		_, err65 := mbTrans64.WriteString(arg63)
		if err65 != nil {
			Usage()
			return
		}
		factory66 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt67 := factory66.GetProtocol(mbTrans64)
		containerStruct0 := rpc.NewMultiPullArgs()
		err68 := containerStruct0.ReadField1(jsProt67)
		if err68 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Keys
		value0 := argvalue0
		fmt.Print(client.MultiPull(value0))
		fmt.Print("\n")
		break
	case "range_push":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RangePush requires 3 args")
			flag.Usage()
		}
		argvalue0, err69 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err69 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err70 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err70 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		arg71 := flag.Arg(3)
		mbTrans72 := thrift.NewTMemoryBufferLen(len(arg71))
		defer mbTrans72.Close()
		_, err73 := mbTrans72.WriteString(arg71)
		if err73 != nil {
			Usage()
			return
		}
		factory74 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt75 := factory74.GetProtocol(mbTrans72)
		containerStruct2 := rpc.NewRangePushArgs()
		err76 := containerStruct2.ReadField3(jsProt75)
		if err76 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.Values
		value2 := argvalue2
		fmt.Print(client.RangePush(value0, value1, value2))
		fmt.Print("\n")
		break
	case "range_pull":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "RangePull requires 2 args")
			flag.Usage()
		}
		argvalue0, err77 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err77 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err78 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err78 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.RangePull(value0, value1))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
