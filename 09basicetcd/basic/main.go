package main

import (
    "context"
    "errors"
    "fmt"
    "go.etcd.io/etcd/clientv3"
    "strconv"
    "time"
)

type EtcdDemo struct {
    Val string
    Key string
}

func (demo EtcdDemo) PutValue(client *clientv3.Client) (*clientv3.PutResponse, error) {
    putResp, err := client.Put(context.TODO(), demo.Key, demo.Val)
    if err != nil {
        fmt.Println("failed to put value, err: ", err)
        return nil, errors.New("failed to put value, err: ")
    }
    fmt.Println("putResp Revision: ", putResp.Header.Revision)
    return putResp, nil
}

func (demo EtcdDemo) GetValue(client *clientv3.Client) (*clientv3.GetResponse, error) {
    getResp, err :=client.Get(context.TODO(), demo.Key)
    if err != nil {
        fmt.Println("filed to get value, err: ", err)
        return nil, errors.New("failed to get value")
    }
    for _, item := range getResp.Kvs {
        fmt.Printf("getval: %s - %s \n", item.Key, item.Value)
    }
    return getResp, nil
}

// DeleteValue 删除的例子， 删除在配置里面来说就可以理解为服务下线了
func (demo EtcdDemo) DeleteValue(client *clientv3.Client) {
    fmt.Println("\n...delete value...")
    delResp, err := client.Delete(context.TODO(), demo.Key)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("del response: ", delResp.PrevKvs)
}

// PutValueDemo 插入值demo，配置里面put相当于注册一个服务（服务注册），get相当于获取一个服务（服务发现）
func (demo EtcdDemo) PutValueDemo(client *clientv3.Client) {
    fmt.Println("\n...put value demo...")
    fmt.Println("put value: ", demo)
    _, err := demo.PutValue(client)
    if err != nil {
        fmt.Println("failed to put value, err: ", err)
        return
    }

    getVal, err := demo.GetValue(client)
    if err != nil {
        fmt.Println("get val err: ", err)
        return
    }
    fmt.Println("get key:",string(getVal.Kvs[0].Key),",value:", string(getVal.Kvs[0].Value), ", Revision: ", getVal.Header.Revision)
}

// watchDemo 监听key的变化
func (demo EtcdDemo) WatchDemo(client *clientv3.Client) {
    fmt.Println("\n...watch demo...")
    stopChan := make(chan interface{}) // 是否停止信号
    go func() {
        watchChan := client.Watch(context.TODO(), demo.Key, clientv3.WithPrefix())
        for {
            select {
            case result := <- watchChan:
                for _, event := range result.Events {
                    fmt.Printf("%s %q: %q \n", event.Type, event.Kv.Key, event.Kv.Value)
                }
                case <-stopChan:
                    fmt.Println("stop watching...")
                    return
            }
        }
    }()

    for i := 0; i < 5; i++ {
        var demo EtcdDemo
        demo.Key = fmt.Sprintf("key_%02d", i)
        demo.Val = strconv.Itoa(i)
        demo.PutValue(client)
    }
    time.Sleep(time.Second * 1)

    stopChan <- 1 //停止watch，在插入就不会监听到了

}

// LeaseDemo 租约
func (demo EtcdDemo) LeaseDemo(client *clientv3.Client) {
    fmt.Println("\n...lease demo...")

    lease, err := client.Grant(context.TODO(), 2) //创建一个租约
    if err != nil {
        fmt.Println("grant err: ", err)
        return
    }

    testKey := "testleasekey"
    // 给这个testkey一个 2秒的TTL租约
    client.Put(context.TODO(), testKey, "testvalue", clientv3.WithLease(lease.ID))
    getVal, err := client.Get(context.TODO(), testKey)
    if err != nil {
        fmt.Println("get val err: ", err)
        return
    }
    vallen := len(getVal.Kvs)
    fmt.Println("before time sleep, val len: ", vallen)

    fmt.Println("sleep 4 seconds")
    time.Sleep(4 * time.Second) //睡眠4秒，让租约过期

    getVal, _ = client.Get(context.TODO(), testKey)
    vallen = len(getVal.Kvs)
    fmt.Println("after 4 seconds, val len: ", vallen)
}

func main() {
    client, err := clientv3.New(clientv3.Config{
        Endpoints:  []string{"127.0.0.1:2379"},
        DialTimeout: time.Second * 5,
    })
    if err != nil {
        fmt.Println("failed to connect etcd: ", err)
        return
    }
    defer client.Close()

    demo := EtcdDemo{Key: "test1", Val: "val1"}
    demo.PutValueDemo(client)

    //deldemo := EtcdDemo{Key:"test01"}
    //deldemo.DeleteValue(client)

    //demo1 := EtcdDemo{}
    //demo1.WatchDemo(client)


    demo2 := EtcdDemo{}
    demo2.LeaseDemo(client)
}