package main

import (
	"fmt"

	"github.com/gyuho/gobay/bay"
)

func main() {
	// 1. Read/Import training data (DATA) , from my GitHub / Google Docs
	DATAAmazon1 := bay.GetStruct("../data/train - amazon.csv")
	include1 := bay.GetInclFt("../data/filter - include.csv")
	exclude1 := bay.GetExcFt("../data/filter - exclude.csv")

	// Pass unfamiliar sentences and see how accurate its sentiment analysis is.
	str1 := "I highly recommend here. Great Weather!"
	str2 := "I hate the movie."
	str3 := "I enjoy it and want to do this again."
	str4 := "Quite disappointed. Never ever again!"
	str5 := "And all of this at a great price."

	fmt.Println("Bayesian Sentiment Analysis in Amazon.com model.")
	bay.Print(DATAAmazon1, include1, exclude1, str1)
	bay.Print(DATAAmazon1, include1, exclude1, str2)
	bay.Print(DATAAmazon1, include1, exclude1, str3)
	bay.Print(DATAAmazon1, include1, exclude1, str4)
	bay.Print(DATAAmazon1, include1, exclude1, str5)

	// Output:
	// Bayesian Sentiment Analysis in Amazon.com model.
	// Strongly Positive: I highly recommend here. Great Weather!
	// Strongly Negative: I hate this movie.
	// Strongly Positive: I enjoy it and want to do this again.
	// Negative: Quite disappointed. Never ever again!
	// Strongly Positive: And all of this at a great price.

	DATACity2 := bay.GetStruct("../data/train - city.csv")
	include2 := bay.GetInclFt("../data/filter - include.csv")
	exclude2 := bay.GetExcFt("../data/filter - exclude.csv")

	fmt.Println("Bayesian Sentiment Analysis in city review model.")
	bay.Print(DATACity2, include2, exclude2, str1)
	bay.Print(DATACity2, include2, exclude2, str2)
	bay.Print(DATACity2, include2, exclude2, str3)
	bay.Print(DATACity2, include2, exclude2, str4)
	bay.Print(DATACity2, include2, exclude2, str5)

	// Output:
	// Bayesian Sentiment Analysis in city review model.
	// Positive: I highly recommend here. Great Weather!
	// Strongly Negative: I hate this movie.
	// Strongly Positive: I enjoy it and want to do this again.
	// Strongly Negative: Quite disappointed. Never ever again!
	// Strongly Positive: And all of this at a great price.

	DATAAmazon3 := bay.GetStruct("../data/train - amazon.csv")
	include3 := bay.GetInclFt("../data/filter - include.csv")
	exclude3 := bay.GetExcFt("../data/filter - exclude.csv")

	// Totally unfamiliar sentence
	// (Correct Classification!)
	// Now this data is trained

	bay.Print(DATAAmazon3, include3, exclude3, "High quality code samples. It must be said: Mark Summerfield is a REALLY good programmer. All of the code in this book gives the impression of being well thought out. The other books had a lot of cargo cult programming, meaning the authors were going through the motions without thinking about what they were doing.")

	// Output:
	// Positive: High quality code samples. It must be said: Mark Summerfield is a REALLY good programmer. All of the code in this book gives the impression of being well thought out. The other books had a lot of cargo cult programming, meaning the authors were going through the motions without thinking about what they were doing.

	DATAAmazon4 := bay.GetStruct("../data/train - amazon.csv")
	include4 := bay.GetInclFt("../data/filter - include.csv")
	exclude4 := bay.GetExcFt("../data/filter - exclude.csv")

	bay.Print(DATAAmazon4, include4, exclude4, "I just paid good money for this book and went to the web site to download the code examples and exercises. The web page is almost totally unreadable since it has a black background with dark gray text. But the worst thing is that there are no links to download the code examples and exercises. While I can read the 3 chapters of updated text in my browser (chrome), the save button doesn't work, I can only print them. There is no link to tell the author any of this, so I have to do it here. I am new to Kindle, but I don't see how I can put this new material back into my Kindle book.")
	// Output:
	// Negative: I just paid good money for this book and went to the web site to download the code examples and exercises. The web page is almost totally unreadable since it has a black background with dark gray text. But the worst thing is that there are no links to download the code examples and exercises. While I can read the 3 chapters of updated text in my browser (chrome), the save button doesn't work, I can only print them. There is no link to tell the author any of this, so I have to do it here. I am new to Kindle, but I don't see how I can put this new material back into my Kindle book.

	// 1. Read/Import training data (DATA) , from my GitHub / Google Docs
	DATAAmazon5 := bay.GetStruct("../data/train - amazon.csv")
	include5 := bay.GetInclFt("../data/filter - include.csv")
	exclude5 := bay.GetExcFt("../data/filter - exclude.csv")

	bay.Print(DATAAmazon5, include5, exclude5, "The book is full of great content and contains the most comprehensive treatment on Go channels and goroutines that I've seen yet. The author also does a phenomenal job with coverage; the table of contents shows this. Virtually every nook and cranny of the Go language is covered. Including the initial founding of Go and its context. (I found this last bit very interesting.) The book is also full of code snippets exemplifying Go idioms and as examples to accomplish common tasks, which is really great considering the infancy of Go.")
	// Output:
	// Positive: The book is full of great content and contains the most comprehensive treatment on Go channels and goroutines that I've seen yet. The author also does a phenomenal job with coverage; the table of contents shows this. Virtually every nook and cranny of the Go language is covered. Including the initial founding of Go and its context. (I found this last bit very interesting.) The book is also full of code snippets exemplifying Go idioms and as examples to accomplish common tasks, which is really great considering the infancy of Go.
}
