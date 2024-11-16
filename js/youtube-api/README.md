## ディレクトリ作成とTypeScriptのインストール

```zsh
# youtube-apiディレクトリ
$ npm install --yes
```

```json
// package.json
"main": "index.js",
// ↓に修正
"main": "index.js",
"type": "module",
```

```
$ npm install --save-dev typescript @types/node
```

## tsconfig.jsonの作成

tsconfig.jsonはTypeScriptコンパイラに対する設定ファイル

```zsh
# デフォルトのtsconfig.jsonを作成
# npxはnpmに同梱されたプログラムで、node_modules内にインストールされたコマンドラインプログラムを実行するツール
$ npx tsc --init
```

```json
// tsconfig.json
// targetはトランスパイル後の程度を指定
"target": "es2016"
// ↓に修正
"target": "es2020"

// moduleはモジュールの形式を指定
"module": "commonjs"
// ↓に修正
"module": "ESNext"

// "moduleResolution": "node"
// ↓に修正
"moduleResolution": "node"

// outDirはトランスパイル後のファイルの出力先
"outDir": "./"
// ↓に修正
"outDir": "./dist"

// srcディレクトリ以下のすべてのtsファイルをコンパイル
{
  "compilerOptions": {
    // たくさんのコンパイラオプション
  },
  // ↓に追加
  "include": ["./src/**/*.ts"]
}
```

## コンパイル

```zsh
# tsconfig.jsonに設定された内容に従ってコンパイル
$ npx tsc
```

```zsh
# 生成されたindex.jsを実行
node dist/index.js
```