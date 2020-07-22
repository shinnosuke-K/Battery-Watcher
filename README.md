# Battery-Watcher

Macの充電情報を取得して保存するプログラムになります。

## 充電量情報について

自分のPCの充電情報を知りたいなら、以下のコマンドを打ちます。

```bash
ioreg -l | grep -v "Apple" | grep  -v "BatteryData" | grep -e "MaxCapacity" -e "DesignCapacity" -e "CurrentCapacity"
```

すると、以下のような結果が表示されると思います。

``` bash
    | |           "MaxCapacity" = 3778
    | |           "CurrentCapacity" = 2515
    | |           "DesignCapacity" = 4315
```

各項目は、

#### MaxCapacity

&ensp;&ensp;現時点での充電可能な最大量を表しています。

#### CurrentCapacity

&ensp;&ensp;現在の充電量を表しています。

#### DesignCapacity

&ensp;&ensp;出荷時の最大充電量を表しています。

## 作成しているもの

コマンドで取得できる充電情報を定期的に習得して、ファイル等に保存するプログラムを作成しています。

### 動作概要

["充電量情報について"](#充電量情報について)で紹介したコマンドをGoのファイルで実行します。

実行後、文字列を加工して、必要な値を取得します。

最後に、取得時間と共にcsvファイルへ書き込みを行います。

### データ
cscファイルへ書き込んでいる情報について紹介します。

```csv
MaxCapacity, CurrentCapacity, DesignCapacity, Rate, Year, Month, Day, Hour
```

#### MaxCapacity, CurrentCapacity, DesignCapacity

&ensp;&ensp;この3つの情報については["充電量情報について"](#currentcapacity)で紹介したものｔ同じになります。


#### Rate

&ensp;&ensp;[DesignCapacity](#designcapacity)に対する[MaxCapacity](#maxcapacity)の割合を表しています。

&ensp;&ensp;式で表すと以下のようになります。

```math
Rate = MaxCapacity / DesignCapacity
```

#### Year, Month, Day, Hour

&ensp;&ensp;充電量情報を取得した時間情報（年、月、日、時）を表しています。

## 気になること

現状はUNIXコマンドをGo言語のパッケージを利用して実行している形になります。

なので、UNIXコマンドを使うなら、シェルスクリプトでもいい気がしています。

なので、シェルスクリプトでもやってみてどちらがいいか検討をする予定です。

## 今後

将来的には、このプログラムを1日2回実行しようと考えています。

なので、バックグランドで実行できるようなプログラムを作成をします。

これに加えて、保存したデータをグラフにプロットするプログラムを別のリポジトリの形で作成も考えています。

グラフを見ることでなにかおもしろいものが見えてくるかもしれません。

なので、まずは半年ぐらいを目処にグラフ化を行い、どこかのLTで発表できればと思います。
