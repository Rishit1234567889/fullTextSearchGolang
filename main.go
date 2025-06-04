package main

import (
	"flag"
	utils "github.com/Rishit1234567889/fullSearch/utils"
	"log"
	"time"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-pages-logging10.xml", "wiki dump path")
	flag.StringVar(&query, "q", "small wild cats", "query to search")
	flag.Parse()
	log.Println("Full text search in progress ")
	start := time.Now()
	docs, err := utils.LoadDocument(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

}
