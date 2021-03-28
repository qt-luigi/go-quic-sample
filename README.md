# go-quic-sample

## 概要

2021年3月28日(日)にGDG Shikoku主催で香川＆オンラインにて開催された

- [【e-とぴあ・かがわ春のオンラインフェス2021】QUIC スタート](https://gdgshikoku.connpass.com/event/207750/)
- [QUICスタート - YouTube](https://www.youtube.com/watch?v=-GgVO3-TN98)

での、Goハンズオンの成果物です。

講師は[Tam](https://twitter.com/tam_x)さんで、Goでサーバーとクライアントのコードを書いてQUICでechoさせました。

## 注意

使用しているパッケージ「[lucas-clemente/quic-go](https://github.com/lucas-clemente/quic-go)」は、更新頻度が高く破壊的変更が多いため、今回はバージョンを v0.19.3 に固定しているとのことです。

## ソースツリー

- server/main.go サーバー
- client/main.go クライアント

サーバーを起動した後にクライアントを実行すると、メッセージが表示されます。
