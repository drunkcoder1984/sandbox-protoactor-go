# request-response
## 説明
このプログラムは、アクターが `Hello` メッセージを受信してそれに応答するアクター システムをセットアップします。この `RequestFuture` メソッドは、アクターにメッセージを送信し、応答を待つために使用されます。このメソッドは、応答を処理する一時的なアクターを作成し、メイン関数が応答の到着またはタイムアウトを待機できるようにするため、非常に重要です。
## メモ
- 詳しい説明は以下のリンクを参照
  - https://github.com/asynkron/protoactor-go/blob/dev/examples/actor-request-response
