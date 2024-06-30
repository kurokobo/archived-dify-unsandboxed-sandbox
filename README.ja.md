# Dify Unsandboxed Sandbox

🌏 [**日本語**](./README.ja.md)
🌏 [**English**](./README.md)

> [!CAUTION]
>
> - **このコンテナイメージを使用すると、あなたの環境は悪意のあるコードからの攻撃に非常に脆弱になります。**
> - **自分が行おうとしていることを自分で充分に理解できていない場合は、このコンテナイメージは使用しないでください。**
> - **本番環境では使用しないでください。**
> - **複数人で共有している Dify 環境では使用しないでください。**
> - **自分だけが使用できる環境で、実験とテストの目的でのみ使用してください。**

## 概要

これは、Dify のサンドボックス環境のすべての制約を取り除くことを目的として開発された、`langgenius/dify-sandbox` を置き換えるコンテナイメージです。

Dify のサンドボックス環境は、悪意のある（または意図せず不正な動作をする）コードから環境を保護するために、さまざまな制限があります。

このコンテナイメージは、これらの制約を取り除き、サンドボックス環境であらゆる任意のコードを実行できるようにします。

| テスト済み言語 |
| --- |
| ✅ Python |
| ⛔ Node.js |

## 免責事項

**このコンテナイメージを使用すると、あなたの環境は悪意のあるコードに非常に脆弱になります。**

- 🚨 **ユーザが実行するコードは、コンテナ内の特権ユーザにより実行されます。**
- 🚨 **ユーザが実行するコードは、コンテナ内のすべてのファイルとプロセスにアクセスできます。**
- 🚨 **ユーザが実行するコードは、コンテナが接続できる任意のネットワークにアクセスできます。**
- 🚨 **コンテナ環境にコンテナエスケープの脆弱性がある場合、攻撃の影響範囲はコンテナだけでなく、ホストおよび接続されている環境全体に及びます。**

**自分が行おうとしていることを自分で充分に理解できていない場合は、このコンテナイメージは使用しないでください。**

**リスクを十分に理解したうえで、使用方法を確認するには、[使用方法ページ](./docs/usage.ja.md) を参照してください。**