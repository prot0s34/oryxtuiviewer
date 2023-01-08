# oryxtuiviewer

View Russia's equipment losses (approved by Oryx) in Terminal

![oryxlogo](https://1.bp.blogspot.com/-HLfgaUHRLnc/XwNGwWwCQYI/AAAAAAAAN5g/Uh4v-kIdiWoEuTZgIs6TUabEYBEzGswDgCK4BGAYYCw/s1600/23.png)

### Executing

```
go run oryxtuiviewer.go
```

or

```
docker build --tag oryxtuiviewer .
docker run -ti oryxtuiviewer
```

or

```
docker pull prot0s/oryxtuiviewer:v0.0.1
```

## Acknowledgments

Inspiration, code snippets, etc.
* [termui](https://github.com/gizak/termui) for TUI
* [goquery](https://github.com/PuerkitoBio/goquery) for scrape https://www.oryxspioenkop.com/

# Preview:
![preview](https://user-images.githubusercontent.com/79843027/211197151-29432ba8-3673-49e5-bce7-129e2e39898c.png)
