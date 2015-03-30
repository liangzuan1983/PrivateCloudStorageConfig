package controllers

import (
    "net"
    "fmt"
	//"github.com/astaxie/beego"
)

func GetLocalIP(inter string) {
    ifi, err := net.InterfaceByName(inter)
    if err != nil {
        //beego.Critcal("GetLocalIP Failed")
    }
    addrs, err := ifi.Addrs()
    if err != nil {
        //beego.Critcal("GetLocalIP Failed")
    }
    for _, a := range addrs {
            fmt.Printf("Interface %q, address %v\n", ifi.Name, a)
    }
}