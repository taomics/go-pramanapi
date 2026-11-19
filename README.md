# go-pramanapi

`github.com/taomics/pramanapi` の Protocol Buffers 定義から生成された Go モジュールです。

## バージョニング

### 正規リリース（CalVer）

正規リリースは日付ベースのカレンダーバージョニング（CalVer）を使用します。

```text
v0.YYYYMMDD.PATCH
```

- 例: `v0.20260914.0`
- 同日に複数回リリースする場合は `PATCH` をインクリメントします（例: `v0.20260914.1`）。

### ベータ版（プレリリース）

開発中の機能をベータ版として配布する場合は、`v0.0.0-` プレフィックス付きのプレリリースバージョンを使用します。

```text
v0.0.0-<feature>.<version>.beta.<seq>
```

- 例: `v0.0.0-healthcheck.0.2.0.beta.1` → `v0.0.0-healthcheck.0.2.0.beta.2`
- `v0.0.0-` を使うことで正規リリース（`v0.YYYYMMDD.0`）より常に小さく評価され、`@latest` や自動アップデートで誤って取り込まれるのを防ぎます。
- 機能名でグルーピングされるため、正規リリースの日付を気にせず複数機能のベータ版を並行して刻めます（必要に応じて `beta` 部分を `alpha` や `rc` にすることもできます）。

#### タグ命名時のルール（Go Modules / SemVer 制約）

- 使用できる文字は英数字（`0-9`, `A-Z`, `a-z`）、ハイフン（`-`）、区切りのドット（`.`）のみです（アンダースコア `_` は不可）。
- ドット区切りのうち数値のみの要素には先頭ゼロを付けないでください（`.01` は不可、`.1` は可）。
- 一度発行したタグは打ち直さず、更新時は必ず新しいタグを作成してください。

## 利用方法

利用側のサービスでは `go.mod` の `replace` ディレクティブでバージョンを指定します（[examples/go.mod](./examples/go.mod) も参照）。

### 正規リリースを利用する

```sh
go mod edit -replace=github.com/taomics/pramanapi=github.com/taomics/go-pramanapi@v0.20260914.0
go mod tidy
```

### ベータ版を導入・更新する

```sh
# 導入
go mod edit -replace=github.com/taomics/pramanapi=github.com/taomics/go-pramanapi@v0.0.0-healthcheck.0.2.0.beta.1
go mod tidy

# 更新
go mod edit -replace=github.com/taomics/pramanapi=github.com/taomics/go-pramanapi@v0.0.0-healthcheck.0.2.0.beta.2
go mod tidy
```

正規リリースへ切り替える際は、同様に `replace` のバージョンを正規リリースのタグへ書き換えて `go mod tidy` を実行してください。
