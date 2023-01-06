基于Golang开发用于动态更新`godaddy`DNS记录的CLI程序

原理：获取运行设备的外网地址（IPV4/IPV6）并更新指定域名的解析结果

使用方法
```bash
❯ ./godaddy-ddns --help                                                                                                                                                                                                  godaddy-ddns
Usage of ./dist/godaddy-ddns_linux_amd64_v1/godaddy-ddns:
      --help               --help
      --domain string      your domain
      --type string        default 'A' (default "A")
      --name string        default nil
      --shopperid string   shopper id
      --key string         api key
      --secret string      api secret
```

举例如更新demo.chancel.me的IPV4地址为当前设备的外网IPV6地址
```bash
./godaddy-ddns --domain=chancel.me --type=AAAA --name=demo --shopperid=123456789 --key=32dyn3mcuabM3uxcnb1fdsaf1s --secret=In3bdbai3sn2
```

`type`参数A表示更新IPV4地址，AAAA表示更新为IPV6地址