NOT MAINTAINED ANYMORE. MERGED TO [**goling**](https://github.com/gyuho/goling).


gobay [![Build Status](https://travis-ci.org/gyuho/gobay.svg?branch=master)](https://travis-ci.org/gyuho/gobay) [![GoDoc](https://godoc.org/github.com/gyuho/gobay?status.png)](http://godoc.org/github.com/gyuho/gobay) [![Project Stats](http://www.ohloh.net/p/714471/widgets/project_thin_badge.gif)](http://www.ohloh.net/p/714471)
==========

Package gobay implements Naive Bayesian Classifier for sentiment analysis. (<a href="http://en.wikipedia.org/wiki/Naive_Bayes_classifier" target="_blank">Naive Bayesian Classifier</a> - Wikipedia)

YouTube Clips by me
==========
<a href="http://www.youtube.com/watch?v=dctzCcYt4AM" target="_blank"><img src="http://img.youtube.com/vi/dctzCcYt4AM/0.jpg"></a>
<ul>
	<li><a href="https://www.youtube.com/watch?v=dctzCcYt4AM&list=PLT6aABhFfintOGKWVWz9qMxC3qZZdHQRD&index=1" target="_blank">Bayesian Classification</li>
	<li><a href="https://www.youtube.com/watch?v=927YDZH_MLo&list=PLT6aABhFfintOGKWVWz9qMxC3qZZdHQRD" target="_blank">String Similarity, Cosine Similarity, Levenshtein Distance</li>
	<li><a href="https://www.youtube.com/watch?v=3qHx1VCcobY&list=PLT6aABhFfintOGKWVWz9qMxC3qZZdHQRD" target="_blank">Spell Check</li>
</ul>

[↑ top](https://github.com/gyuho/gobay#gobay---)


Getting Started
==========
- [godoc.org](http://godoc.org/github.com/gyuho/gobay)
- [gowalker.org](http://gowalker.org/github.com/gyuho/gobay#_index)

```go
go get github.com/gyuho/gobay
```
[↑ top](https://github.com/gyuho/gobay#gobay---)


Package Hierarchy
==========
```go
bay/		# Naive Bayesian Classifier
data/		# Training Data
example/	# Example Code
slm/		# Slice, Map Operations
```
[↑ top](https://github.com/gyuho/gobay#gobay---)


Example
==========
```go
func Test_NBC_2(test *testing.T) {
	DATA_amazon := bay.GetStruct("../data/train - amazon.csv")
	include := bay.GetInclFt("../data/filter - include.csv")
	exclude := bay.GetExcFt("../data/filter - exclude.csv")

	// Totally unfamiliar sentence
	// (Correct Classification!)
	// Now this data is trained

	bay.Print(DATA_amazon, include, exclude, "High quality code samples. It must be said: Mark Summerfield is a REALLY good programmer. All of the code in this book gives the impression of being well thought out. The other books had a lot of cargo cult programming, meaning the authors were going through the motions without thinking about what they were doing.")

	// Output:
	// Positive: High quality code samples. It must be said: Mark Summerfield is a REALLY good programmer. All of the code in this book gives the impression of being well thought out. The other books had a lot of cargo cult programming, meaning the authors were going through the motions without thinking about what they were doing.
}
```
[↑ top](https://github.com/gyuho/gobay#gobay---)


Training Data
==========
Training and filter data are to be frequentyly updated, directly on GitHub and Google Docs.

- <a href="https://docs.google.com/spreadsheet/ccc?key=0AvwDSsSZw04HdF95Rzdubi0xdnJSZXVsYU1OTk9hZWc&usp=sharing" target="_blank">train data : Google Docs</a>
	- sample: amazon.com review, city review models...
	- range from 1 to 10; 10 is most positive
	- We can add any category(class) you want; sports, newspaper, ...

- <a href="https://github.com/gyuho/gobay/blob/master/data/train%20-%20amazon.csv" target="_blank">train - amazon.csv : GitHub</a>

- <a href="https://github.com/gyuho/gobay/blob/master/data/train%20-%20city.csv" target="_blank">train - city.csv : GitHub</a>

- <a href="https://docs.google.com/spreadsheet/ccc?key=0AvwDSsSZw04HdHY3OVNLN1pXb0VMOEFhLVZWb0RNRVE&usp=sharing" target="_blank">filter data : Google Docs</a>
	- feature candidate word selection
	- signal words

- <a href="https://github.com/gyuho/gobay/blob/master/data/filter%20-%20exclude.csv" target="_blank">filter - exclude.csv : GitHub</a>

- <a href="https://github.com/gyuho/gobay/blob/master/data/filter%20-%20include.csv" target="_blank">filter - include.csv : GitHub</a>


[↑ top](https://github.com/gyuho/gobay#gobay---)

To-Do-List
==========
- Update Bayesian algorithms for some exceptional cases

[↑ top](https://github.com/gyuho/gobay#gobay---)
