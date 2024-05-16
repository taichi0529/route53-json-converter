# Route53 Json Converter

AWS Route 53のhosted zoneを別のアカウントに移すためのjson変換ツールです。

## ビルド方法
goがビルド出来る環境で下記を実行して下さい。

```bash
go build convert.go
```

## hosted zoneの移行手順

1. 既存のhosted zoneのレコードをjsonとして出力
2. 出力したjsonをインポート出来る形に変換
3. 新しいアカウントにhosted zoneを作成
4. 2で変換したjsonをインポート

この中で2以外はAWS CLIもしくはコンソールから行えますが、2の変換はこのツールを使って行います。
変換自体は大した作業ではないのですが数が多いと面倒なのでプログラムにしました。

変換内容の詳細詳細は下記を参照して下さい。
https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/hosted-zones-migrating.html#hosted-zones-migrating-edit-records


```bash
aws route53 list-resource-record-sets --hosted-zone-id ZXXXXXXXXXXXXXXXXXXX > xxxxxxxx.json
./convert xxxxxxxx.json > xxxxxxxx-converted.json
```

このあと新しいアカウントでhosted zoneを作成し、2で変換したjsonをインポートします。

```bash
aws route53 change-resource-record-sets --hosted-zone-id ZYYYYYYYYYYYYYYYYYYYY --change-batch file://xxxxxxxx-converted.json
```
