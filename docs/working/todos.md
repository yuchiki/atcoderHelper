
# 実装予定リスト

この項目は #開発者メモ に移動し、適宜issueに移し変えていく

チェックがついていない機能は未実装です。

以下ではURLの `atcoder.jp/contests/xxx` の `xxx` の部分をcontestIDと呼びます。
以下ではURLの `atcoder.jp/contest/:contestID/tasks/yyy` の `yyy` の部分をtaskIDと呼びます

- [ ] `ach login` : atcoderにloginする

- [ ] `ach config use-template <template>` 使用するtemplateを選択する
- [ ] `ach language add --name <name> --atcoderName <atcoderName> --build <buildCommand> --run <runCommand>` : languageを追加する
- [ ] `ach template add --name <name> --language <language> --dir <directory>`: repository を
- [ ] `ach config view` : configを表示する
- [ ] `ach config language list` : 登録された言語たちを表示する
- [ ] `ach config language describe <language>` : 指定のtemplateの詳細を表示する
- [ ] `ach config template list` : 登録されたtemplateたちを表示する
- [ ] `ach config template describe <template>` : 指定のtemplateの詳細を表示する
- [ ] `ach version` : versionを表示する
- [ ] `ach contest upcoming` : 予定されたコンテスト一覧を取得する
  - `<contestID>: <contest名>` の形式
- [ ] `ach contest list xxx` xxxと部分一致するcontest名のlistを取得する
  - `<contestID>: <contest名>` の形式
- [ ] `ach contest create (--template <template>) <contestID>` : contestID用のdirectoryを生成する
- [ ] `ach task create (--template <template>) <contestID> <taskID>`: contestID用のdirectoryを生成する

  - directory 構成は 以下のようになる

```
abc190
  - achContestConfig.yaml
  - abc190_a
    - achTaskConfig.yaml
    - program
      - (configで設定されたtemplateがここに展開される)
    - sampleCases
      - case1.input
      - case1.output
      - case2.input
      - case2.output
      - case3.input
      - case3.output
      - case4.input // 出力を確認したいときにinputファイルのみを手動で追加することもできる
  - bbc190_b
  ...
```

以下のコマンドは contest repo 以下でのみ有効

- [ ] `ach contest status` : 問題ごとにAC/CE などの状態を閲覧できる


以下のコマンドは task repo以下でのみ有効

- [ ] `ach test` : sampleCasesをテストする
- [ ] `ach test <n>` :  case(n) のみをテストする
- [ ] `ach test --submit` :  全てのテストに通ったら自動でsubmitする
  - configでdefaultの設定を変更可能
- [ ] `ach submit` : 問題をsubmitする


# configファイルについて

`$HOME/.ach/config.yaml` ファイルは例えば以下のようになるはずである

```yaml
username: "foo"
(なんらかのloginに必要な情報): "bar"
languages:
- name: "c"
  atcoderName: "C(GCC 9.2.1)"
  build: "gcc $SOURCE_CODE -o a.out"
  run: "./a.out"
- name: "cpp"
  atcoderName: "C++(GCC 9.2.1)"
  build: "g++ $SOURCE_CODE -o a.out"
  run: "./a.out"
- name: "csharp"
  atcoderName: "C#(.NET Core 3.1.201)"
  build: ""
  run: "dotnet run"
-　name: "cs"
  alias_to: "csharp"
templates:
-　name: "csharp_default"
  language: "csharp"
  source_code: "Program.cs"
- name: "csharp_fancy"
  language: "csharp"
  source_code: "HeavilyCustomizedCode.cs"
- name: "cpp_default"
  language: "cpp"
  source_code: "main.cpp"
current-template: "csharp_default"
```

`contest_foo/achContestConfig.yaml` は以下のようになるはずである

```yaml
contestID: "foo"
```

`task_foo/achTaskConfig.yaml` は以下のようになるはずである

```yaml
contestID: "foo"
taskID: "bar"
template: "csharp_default"
```
