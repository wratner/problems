package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestProcessFile(t *testing.T) {
	wordBank := Words{}
	err := wordBank.processFiles()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetWordBank(t *testing.T) {
	expected := []string{"time", "year", "people", "way", "day", "man", "thing", "woman", "life", "child", "world", "school", "state", "family", "student", "group", "country", "problem", "hand", "part", "place", "case", "week", "company", "system", "program", "question", "work", "government", "number", "night", "point", "home", "water", "room", "mother", "area", "money", "story", "fact", "month", "lot", "right", "study", "book", "eye", "job", "word", "business", "issue", "side", "kind", "head", "house", "service", "friend", "father", "power", "hour", "game", "line", "end", "member", "law", "car", "city", "community", "Name", "president", "team", "minute", "idea", "kid", "body", "information", "back", "parent", "face", "others", "level", "office", "door", "health", "person", "art", "war", "history", "party", "result", "change", "morning", "reason", "research", "girl", "guy", "momtime", "year", "people", "way", "day", "man", "thing", "woman", "life", "child", "world", "school", "state", "family", "student", "group", "country", "problem", "hand", "part", "place", "case", "week", "company", "system", "program", "question", "work", "government", "number", "night", "point", "home", "water", "room", "mother", "area", "money", "story", "fact", "month", "lot", "right", "study", "book", "eye", "job", "word", "business", "issue", "side", "kind", "head", "house", "service", "friend", "father", "power", "hour", "game", "line", "end", "member", "law", "car", "city", "community", "Name", "president", "team", "minute", "idea", "kid", "body", "information", "back", "parent", "face", "others", "level", "office", "door", "health", "person", "art", "war", "history", "party", "result", "change", "morning", "reason", "research", "girl", "guy", "moment", "air", "teacher", "force", "educationent", "air", "teacher", "force", "education"}
	result, err := getWordBank("invalidFile.txt")
	if err == nil {
		t.Error("File does not exist and should fail")
	}
	result, err = getWordBank(NounsFile)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v Received: %v", expected, result)
	}
}

func TestScrubInput(t *testing.T) {
	expected := []string{"agree", "agreed", "do", "did", "know", "knew", "read", "read", "suggest", "suggested", "allow", "allowed", "eat", "ate", "learn", "learned", "remember", "remembered", "take", "took", "answer", "answered", "explain", "explained", "leave", "left", "run", "ran", "talk", "talked", "ask", "asked", "fall", "fell", "like", "liked", "say", "said", "tell", "told", "be", "was/were", "feel", "felt", "listen", "listened", "see", "saw", "think", "thought", "become", "became", "fill", "filled", "live", "lived", "sell", "sold", "travel", "travelled", "begin", "began", "find", "found", "look", "looked", "seem", "seemed", "try", "tried", "believe", "believed", "finish", "finished", "lose", "lost", "send", "sent", "turn", "turned", "borrow", "borrowed", "follow", "followed", "make", "made", "set", "set", "understand", "understood", "break", "broke", "fly", "flew", "may", "might", "shall", "use", "used (to)", "bring", "brought", "forget", "forgot", "mean", "meant", "should", "wait", "waited", "buy", "bought", "get", "got", "meet", "met", "show", "showed", "wake up", "woke up", "call", "called", "give", "gave", "move", "moved", "sit", "sat", "walk", "walked", "can", "could", "go", "went", "must", "sleep", "slept", "want", "wanted", "carry", "carried", "happen", "happened", "need", "needed", "speak", "spoke", "watch", "watched", "change", "changed", "have", "had", "open", "opened", "spend", "spent", "will", "would", "close", "closed", "hear", "heard", "pay", "paid", "stand", "stood", "win", "won", "come", "came", "help", "helped", "play", "played", "start", "started", "work", "worked", "cut", "cut", "hold", "held", "promise", "promised", "stop", "stopped", "worry", "worried", "decide", "decided", "keep", "kept", "put", "put", "study", "studied", "write", "wrote"}
	result := scrubInput("agree,agreed,do,did,know,knew,read,read,suggest,suggested,allow,allowed,eat,ate,learn,learned,remember,remembered,take,took,answer,answered,explain,explained,leave,left,run,ran,talk,talked,ask,asked,fall,fell,like,liked,say,said,tell,told,be,was/were,feel,felt,listen,listened,see,saw,think,thought,become,became,fill,filled,live,lived,sell,sold,travel,travelled,begin,began,find,found,look,looked,seem,seemed,try,tried,believe,believed,finish,finished,lose,lost,send,sent,turn,turned,borrow,borrowed,follow,followed,make,made,set,set,understand,understood,break,broke,fly,flew,may,might,shall,use,used (to),bring,brought,forget,forgot,mean,meant,should,wait,waited,buy,bought,get,got,meet,met,show,showed,wake up,woke up,call,called,give,gave,move,moved,sit,sat,walk,walked,can,could,go,went,must,sleep,slept,want,wanted,carry,carried,happen,happened,need,needed,speak,spoke,watch,watched,change,changed,have,had,open,opened,spend,spent,will,would,close,closed,hear,heard,pay,paid,stand,stood,win,won,come,came,help,helped,play,played,start,started,work,worked,cut,cut,hold,held,promise,promised,stop,stopped,worry,worried,decide	,decided,keep,kept,put,put,study,studied,write,wrote")
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v Received: %v", expected, result)
	}
}

func TestWriteParagraph(t *testing.T) {
	expected := "\tIs a test.\n"
	words := &Words{
		Nouns:      []string{"is"},
		Verbs:      []string{"a"},
		Adjectives: []string{"test"},
	}
	result := words.writeParagraph(1)
	if expected != result {
		t.Errorf("Expected: %s Received: %s", expected, result)
	}
	expected = "\tIs a test. Is a test is.\n"
	result = words.writeParagraph(2)
	if expected != result {
		t.Errorf("Expected: %s Received: %s", expected, result)
	}
}

func TestBuildSentence(t *testing.T) {
	expected := "Is a test."
	words := &Words{
		Nouns:      []string{"is"},
		Verbs:      []string{"a"},
		Adjectives: []string{"test"},
	}
	result := words.buildSentence(3)
	if expected != result {
		t.Errorf("Expected: %s Received: %s", expected, result)
	}
	words = &Words{
		Nouns:      []string{"is", "it", "test"},
		Verbs:      []string{"a", "is"},
		Adjectives: []string{"test", "another"},
	}
	expectedLength := 7
	result = words.buildSentence(7)
	if expectedLength != len(strings.Split(result, " ")) {
		t.Errorf("Expected: %d Received: %d", expectedLength, len(strings.Split(result, " ")))
	}
}
