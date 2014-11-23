Refactoring

Will be rewritten...

goling [![Build Status](https://travis-ci.org/gyuho/goling.svg?branch=master)](https://travis-ci.org/gyuho/goling) [![GoDoc](https://godoc.org/github.com/gyuho/goling?status.png)](http://godoc.org/github.com/gyuho/goling) [![Project Stats](http://www.ohloh.net/p/714473/widgets/project_thin_badge.gif)](http://www.ohloh.net/p/714473)
==========

Package goling implements spell-check, segmentation, and string similarity algorithms(Levenshtein Distance, Cosine Similarity).


Getting Started
==========
- [godoc.org](http://godoc.org/github.com/gyuho/goling)
- [gowalker.org](http://gowalker.org/github.com/gyuho/goling#_index)

```go
go get github.com/gyuho/goling
```


YouTube Clips by me
==========
<a href="http://www.youtube.com/watch?v=927YDZH_MLo" target="_blank"><img src="http://img.youtube.com/vi/927YDZH_MLo/0.jpg"></a>
<ul>
	<li><a href="https://www.youtube.com/watch?v=dctzCcYt4AM&list=PLT6aABhFfintOGKWVWz9qMxC3qZZdHQRD&index=1" target="_blank">Bayesian Classification</li>
	<li><a href="https://www.youtube.com/watch?v=927YDZH_MLo&list=PLT6aABhFfintOGKWVWz9qMxC3qZZdHQRD" target="_blank">String Similarity, Cosine Similarity, Levenshtein Distance</li>
	<li><a href="https://www.youtube.com/watch?v=3qHx1VCcobY&list=PLT6aABhFfintOGKWVWz9qMxC3qZZdHQRD" target="_blank">Spell Check</li>
</ul>

[↑ top](https://github.com/gyuho/goling#goling---)


Package Hierarchy
==========
```go
data/		# Training Data
sg/			# String Segmentation Algorithm
simi/		# String Similarity Algorithms
spell/		# Spell Check Algorithms
st/			# String Manipulating Functions
```
[↑ top](https://github.com/gyuho/goling#goling---)



Example
==========
```go
// similarly related
str1 := "golang: read text file into string array (and write)"
str2 := "writing to a text file gets cut short using java"
StringSimilarity(str1, str2)
// 1538.2096034883082 

str5 := "What's going on has to do with the pragmatics of conversation. In particular, you're flouting a conversational rule, called the Maxim of Quantity.This rule states that what a person contributes to a conversation should be neither more nor less than what is required in the context."
str6 := "I think Go has a great future in Web development. For other applicative areas, I'm less certain as the development of the runtime is entirely focused on the problems associated with running stuff for the web: large static binaries, etc."
StringSimilarity(str5, str6)
// 321.95180522788064
```

[↑ top](https://github.com/gyuho/goling#goling---)
