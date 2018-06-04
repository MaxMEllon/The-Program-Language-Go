### file{1,2,3}.md は CC-BY-SA 3.0 のライセンスに従い二次利用しています

![](https://upload.wikimedia.org/wikipedia/commons/thumb/d/d0/CC-BY-SA_icon.svg/88px-CC-BY-SA_icon.svg.png)

> https://ja.wikipedia.org/wiki/Wikipedia:ウィキペディアを二次利用する#ライセンスに従った二次利用方法

### 実行結果

```
    ch1/ex4  master   go run main.go file1.md file2.md
Vimはオランダ人のプログラマーBram MoolenaarによってAmiga向けに開発された。
        7       file1.md
        4       file2.md
Emacsは1970年代のMIT人工知能研究所（MIT AI研）で産声をあげた。
        2       file1.md
        6       file2.md
「EMACSは、共同参加を基として頒布される。つまり改良点は全て、組み入れて頒布するために、私のところへ戻ってこなければならない」
        2       file1.md
        19      file2.md
```


### 検証

```
    ch1/ex4  master   cat file1.md | grep 'Vimはオランダ人のプログラマーBram MoolenaarによってAmiga向けに開発された。' | wc -l
       7
    ch1/ex4  master   cat file2.md | grep 'Vimはオランダ人のプログラマーBram MoolenaarによってAmiga向けに開発された。' | wc -l
       4

    ch1/ex4  master   cat file1.md | grep 'Emacsは1970年代のMIT人工知能研究所（MIT AI研）で産声をあげた。' | wc -l
       2
    ch1/ex4  master   cat file2.md | grep 'Emacsは1970年代のMIT人工知能研究所（MIT AI研）で産声をあげた。' | wc -l
       6

    ch1/ex4  master   cat file1.md | grep '「EMACSは、共同参加を基として頒布される。つまり改良点は全て、組み入れて頒布するために、私のところへ戻ってこなければならない」' | wc -l
       2
    ch1/ex4  master   cat file2.md | grep '「EMACSは、共同参加を基として頒布される。つまり改良点は全て、組み入れて頒布するために、私のところへ戻ってこなければならない」' | wc -l
      19

```
