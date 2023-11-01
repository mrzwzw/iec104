package main

import (
	"fmt"
	"log"
	"time"

	"github.com/9d77v/iec104"
	_ "github.com/9d77v/iec104/example/client/worker"
	"github.com/sirupsen/logrus"
)

type myClient struct{}

// 总召
func zz() {
	address := "172.23.12.151:2404"
	// address := "127.0.0.1:6666"
	//subAddress := ""
	//if config.SubServerHost != "" && config.SubServerPort != 0 {
	//	subAddress = fmt.Sprintf("%s:%d", config.SubServerHost, config.ServerPort)
	//}
	mycli := &myClient{}
	logger := logrus.New()
	// logger.SetLevel(logrus.DebugLevel)
	// logger.Hooks.Add(utils.NewContextHook())
	client := iec104.NewClient(mycli, address, logger, nil, false)
	defer client.Close()

	client.Run()
	client.SendTotalCall()
	time.Sleep(30 * time.Second)
	// client.SendElectricityTotalCall()

	// time.Sleep(15 * time.Second)
}

// 电度总召
func dzzz() {
	address := "172.23.12.151:2404"
	// address := "127.0.0.1:6666"
	//subAddress := ""
	//if config.SubServerHost != "" && config.SubServerPort != 0 {
	//	subAddress = fmt.Sprintf("%s:%d", config.SubServerHost, config.ServerPort)
	//}
	mycli := &myClient{}
	logger := logrus.New()
	// logger.SetLevel(logrus.DebugLevel)
	// logger.Hooks.Add(utils.NewContextHook())
	client := iec104.NewClient(mycli, address, logger, nil, false)
	defer client.Close()
	client.Run()
	client.SendElectricityTotalCall()

	time.Sleep(30 * time.Second)
}

func main() {
	zz()
	dzzz()
	log.Println("数据量:", len(addrValue))

	for i := 1; i < 4000; i++ {
		_, ok := addrValue[uint32(i)]
		if ok {
			log.Println("addr", i)
			log.Println("value", addrValue[uint32(i)])
		}

	}
}

var (
	addrValue = make(map[uint32]float64)
	a         = 0
)

func (c *myClient) Datahandler(data *iec104.APDU) error {
	if data.Signals[0].TypeID == 21 {
		i := int64(data.Signals[0].Ts)
		tm := time.Unix(i, 0)
		fmt.Println(tm)
	}
	println("do task3")
	for _, signal := range data.Signals {
		addr := signal.Address
		log.Println("address", addr)
		log.Println("value", signal.Value)

		a = a + 1
		log.Println("aaaaaaaa:", a)
		addrValue[addr] = signal.Value
	}
	return nil
}
