package main

import (
        "fmt"
        "time"
        "os"
        "github.com/hekmon/transmissionrpc"
        "github.com/k0kubun/pp"
)

const (
        ip = "192.168.1.15"
        user = "wmw"
        pass = "M9H4gntH1f56"
)


func main() {
        if len(os.Args) < 2 {
                println("No input file. Example: \"attt.exe 123.torrent\"")
                time.Sleep(time.Second * 2)
                panic("No input file. Example: \"attt.exe 123.torrent\"")
        }
        println("Trying to add torrent: ", os.Args[1])
        transmissionbt, _ := transmissionrpc.New(ip, user, pass,
            &transmissionrpc.AdvancedConfig{
                HTTPS: false,
                Port:  9091,
        })
        ok, serverVersion, serverMinimumVersion, err := transmissionbt.RPCVersion()
        if err != nil {
            panic(err)
        }
        if !ok {
            panic(fmt.Sprintf("Remote transmission RPC version (v%d) is incompatible with the transmission library (v%d): remote needs at least v%d",
                serverVersion, transmissionrpc.RPCVersion, serverMinimumVersion))
        }
        fmt.Printf("Remote transmission RPC version (v%d) is compatible with our transmissionrpc library (v%d)\n",
            serverVersion, transmissionrpc.RPCVersion)      
        
        torrents, err := transmissionbt.TorrentGetAll()
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        } else {
                for _, v := range torrents {
                        pp.Println(v.Name) // the only instanciated field, as requested
                }
        }


        // filepath := "2.torrent"
        filepath := os.Args[1]
        torrent, err := transmissionbt.TorrentAddFile(filepath)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        fmt.Printf("\n\"%v\" sucessfully added. \n\nID: %v\nName: %v\nHash: %v\"\n", os.Args[1], *torrent.ID, *torrent.Name, *torrent.HashString)
        //pp.Println(torrent)
        time.Sleep(time.Second * 2)
}
