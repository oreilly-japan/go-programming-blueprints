# Go言語によるWebアプリケーション開発

---

![表紙](1752_go_prg_blueprints_cvr_w_r.png)

---

本リポジトリはオライリー・ジャパン発行書籍『[Go言語によるWebアプリケーション開発](http://www.oreilly.co.jp/books/9784873117522/)』（原書名『[Go Programming Blueprints](https://www.packtpub.com/application-development/go-programming-blueprints)』） のサポートサイトです。

## サンプルコード

サンプルコードの解説は本書籍をご覧ください。

[version1ブランチ](../../tree/version1)の[変更履歴](../../commits/version1)で本書籍の途中のコードを参照することができます。

原書のサンプルコードは[github.com/matryer/goblueprints](https://github.com/matryer/goblueprints)でアクセスできます。

## 正誤表

下記のとおり、本書に誤りがありました。お詫びして訂正いたします。

誤植など間違いを見つけた方は、japan＠oreilly.co.jpまでお知らせください。

### 第3刷まで

#### ■p.106 12行目
**誤**

```
log.Fatalf("%qに類語はありませんでした\n")
```

**正**

```
log.Fatalf("%qに類語はありませんでした\n", word)
```

#### ■p.186 29行目
**誤**

```
package meander 
type Place struct { 
  *googleGeometry `json:"geometry"` 
  Name string `json:"name"` 
  Icon string `json:"icon"` 
  Photos []*googlePhoto `json:"photos"` 
  Vicinity string `json:"vicinity"` 
} 
type googleResponse struct { 
  Results []*Place `json:"results"` 
} 
type googleGeometry struct { 
  *googleLocation `json:"location"` 
} 
type googleLocation struct { 
  Lat float64 `json:"lat"` 
  Lng float64 `json:"lng"` 
} 
type googlePhoto struct { 
  PhotoRef string `json:"photo_reference"` 
  URL string `json:"url"` 
} 
```

**正**

```
package meander
type Place struct {
  Geometry *googleGeometry `json:"geometry"`
  Name     string          `json:"name"`
  Icon     string          `json:"icon"`
  Photos   []*googlePhoto  `json:"photos"`
  Vicinity string          `json:"vicinity"`
}
type googleResponse struct {
  Results []*Place `json:"results"`
}
type googleGeometry struct {
  Location *googleLocation `json:"location"`
}
type googleLocation struct {
  Lat float64 `json:"lat"`
  Lng float64 `json:"lng"`
}
type googlePhoto struct {
  PhotoRef string `json:"photo_reference"`
  URL      string `json:"url"`
}

```

#### ■p.187 31行目
**誤**

```
    "lat": p.Lat,
    "lng": p.Lng,
```

**正**

```
    "lat": p.Geometry.Location.Lat,
    "lng": p.Geometry.Location.Lng,
```

### 第2刷まで

#### ■p.36, 41, 46, 49, 50, 53, 61, 63, 77, 88のコマンド

**誤**

```
./chat -host=":8080"
```

**正**

```
./chat -addr=":8080"
```

#### ■p.24 3行目

**誤**

```
必要な最小限のコードをtrace.goに追加します。
```

**正**

```
必要な最小限のコードをtracer.goに追加します。
```

#### ■p.73 2行目

**誤**

```
io.WriteString(m, strings.ToLower(user.Name()))
```

**正**

```
io.WriteString(m, strings.ToLower(user.Email()))
```

#### ■p.86 13行目

**誤**

```
io.WriteString(m, strings.ToLower(user.Name()))
```

**正**

```
io.WriteString(m, strings.ToLower(user.Email()))
```

### 第1刷

#### ■p.158 ノート記事の1行目

**誤**

```
GowebやGorillzによるmuxパッケージなどは、
```

**正**

```
GowebやGorillaによるmuxパッケージなどは、
```
