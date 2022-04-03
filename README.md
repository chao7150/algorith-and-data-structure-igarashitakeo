# igarashitakeo

『データ構造とアルゴリズム』（五十嵐健夫, 2007）を読みながらC++で自分で実装してみるリポジトリ。

それほど強い縛りプレイをする気はないのでvectorを使ってズルしたりします。

## ディレクトリ構造

- hoge.cpp
  - メインのロジックを記述する
- hoge.h
  - ヘッダーファイル
- hoge_test.cpp
  - テストファイル。基本的にテスト駆動で書いている。googletestを利用している。
- main.cpp
  - デバッグ実行のためにhoge.cppを利用した簡単なスクリプトを書く

## 実行

vscodeの場合。

- テスト
  - hoge.cppを開いた状態でCtrl+Shift+B
  - Tasks: Run Task
  - test
- デバッグ実行
  - hoge.cppを開いた状態でF5
