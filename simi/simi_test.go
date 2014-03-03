package simi

import "testing"

func Test_StringSimilarity(t *testing.T) {

	// similarly related (1538.2096034883082)
	str1 := "golang: read text file into string array (and write)"
	str2 := "writing to a text file gets cut short using java"

	// strongly similar (468.2115223342196)
	str3 := "Programming Interviews: What are some good resources to brush up on (or get better at) system designing interview questions? Specifically when you need to design a system all the way from the data center (and the machines in it) to picking the client-side architecture and framework."
	str4 := "I want to work on client side frameworks. For programming interview, I need to design a system from a data center or machines. What would be the good resource to prepare this interview?"
	if StringSimilarity(str1, str2) < StringSimilarity(str3, str4) {
		t.Errorf("Error: str1, str2 should be more similar than str3, str4")
	}

	// not similar at all
	str5 := "What's going on has to do with the pragmatics of conversation. In particular, you're flouting a conversational rule, called the Maxim of Quantity.This rule states that what a person contributes to a conversation should be neither more nor less than what is required in the context."
	str6 := "I think Go has a great future in Web development. For other applicative areas, I'm less certain as the development of the runtime is entirely focused on the problems associated with running stuff for the web: large static binaries, etc."
	if StringSimilarity(str5, str6) != 321.95180522788064 {
		t.Errorf("Error: str5, str6, StringSimilarity")
	}
}
