# router-limitconcurrency
## 説明
このプログラムでは、ラウンドロビン・ルーターを使用してアクター・システムを作成し、アクター・プール間でワークアイテムを分配します。最大同時実行レベルは `maxConcurrency` に設定されます。各アクターは `workItem` メッセージを処理し、同時実行レベルが定義された制限を超えないようにします。この例は、アクターベースのシステムにおける同時実行制御を理解する上で重要です。
## メモ
- 詳しい説明は以下のリンクを参照
  - https://github.com/asynkron/protoactor-go/blob/dev/examples/router-limitconcurrency
