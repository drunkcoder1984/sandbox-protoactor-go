# actor-setbehavior
## 説明
このプログラムでは2つのビヘイビアを持つ `SetBehaviorActor` を定義している： One` と `Other` である。最初は、アクタは `Hello` メッセージに対して `One` ビヘイビアで応答する。Hello` メッセージを受信すると、それ以降のメッセージは `Other` ビヘイビアに切り替わります。この例では、Proto.Actor のアクターが状態や受信したメッセージに基づいて応答を適応させる、動的な動作機能を説明しています。
## メモ
- 詳しい説明は以下のリンクを参照
  - https://github.com/asynkron/protoactor-go/blob/dev/examples/actor-setbehavior
