# Whats This?

`docker-comopse`でやる、複数のgoアプリで互いに同じvolumeにファイルを読み書きするサンプル

`go_app1`と`go_app2`が同じフォルダ（それぞれのコンテナでは別の名前）を参照する。`go_app1`はそのフォルダにファイルを生成する。`go_app2`では、そのフォルダの変更を検知し、生成イベントがあった場合、それを別のフォルダ（`go_app2`の中では`saveFolderName`の値）にコピーし、コピー後に削除する。同じvolumeを共有しているので、`go_app2`で削除したファイルは当然`go_app1`からも消える。

ファイルの変更検知は[fsnotify](https://github.com/fsnotify/fsnotify)を使用している。

同じvolumeは`docker-compose.yml`では`volumes`で指定したもの。これを、それぞれのコンテナが起動するときにバインドしている。


# Note
`go_app2`が起動するよりも前のファイルは移動・コピーの対象にはならないので、その場合は手動で直接フォルダにいれちゃえばたぶん動く