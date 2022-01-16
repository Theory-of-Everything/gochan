package main

import (
	"os"
	"flag"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"gochan/lib/apiHandler"
)

// command line flags
var (
	fBoard    string // the board being picked
	fPageNum  int    // the page number of threads in that board
	fThreadId int
	fPostId   int // the post id
	fPostOP   int // used with -i, the original thread to search for a comment via id

	fHelp     bool // print help dialouge

	MaxPages  int = 10 // the maxium of pages in a request, this is the hard limit
)

func getPage(request apiHandler.ReqBoardThreads, converter md.Converter) {
	// length of request
	reqLen := len(request[fPageNum].Threads)
	for i := 0; i < reqLen; i++ {
		// get the content of a thread, including
		// title, author, date, and content (in markdown)
		fmt.Print("-------------BEGIN CONTENT #")
		fmt.Println(i)
		thread := apiHandler.GetThread(fBoard, int(request[fPageNum].Threads[i].No))
		for j := 0; j < len(thread.Posts); j++ {

			// if the thread title is empty, just print
			// the post title
			threadTitle := thread.Posts[j].Sub
			if threadTitle == "" {
				fmt.Print("Post #")
				fmt.Println(j)
			} else {
				fmt.Println(threadTitle)
			}

			fmt.Println(thread.Posts[j].Name) // print author name (90% of the time anonymous)
			fmt.Println(thread.Posts[j].Now)  // print date of post
			fmt.Println(thread.Posts[j].No)   // print ID of post (for comment links)

			// convert the html markup of the post into
			// markdown for easier reading
			fmt.Println("Post Content:")
			content, err := converter.ConvertString(thread.Posts[j].Com)
			if err != nil {
				panic("Issue with converting post" + string(j) + "'s content into markdown")
			}
			fmt.Println(content)

			// bar for cleanliness
			fmt.Println("---------------------")

		}
		// bar for cleanliness
		fmt.Print("-------------END CONTENT #")
		fmt.Println(i)
	}

}

// search a thread for a certain post/comment (by id)
func getPostById(op int, id int, converter md.Converter) {

	// if the op is the id, then just print the entire thread
	if op == id {
		thread := apiHandler.GetThread(fBoard, op)
		for i := 0; i < len(thread.Posts); i++ {

			// print title and date of post
			fmt.Println(thread.Posts[i].Name) // print author name (90% of the time anonymous)
			fmt.Println(thread.Posts[i].Now)  // print date of post
			fmt.Println(thread.Posts[i].No)   // print ID of post (for comment links)

			// convert html markup to markdown for post body
			content, err := converter.ConvertString(thread.Posts[i].Com)
			if err != nil {
				panic("Issue with converting post" + string(i) + "'s content into markdown")
			}
			fmt.Println(content)
		}
		// search for a post of id <id> in op thread
	} else {
		thread := apiHandler.GetThread(fBoard, op)
		for i := 0; i < len(thread.Posts); i++ {
			if thread.Posts[i].No == id {

				// print title and date of post
				fmt.Println(thread.Posts[i].Name) // print author name (90% of the time anonymous)
				fmt.Println(thread.Posts[i].Now)  // print date of post
				fmt.Println(thread.Posts[i].No)   // print ID of post (for comment links)

				// convert html markup to markdown for post body
				content, err := converter.ConvertString(thread.Posts[i].Com)
				if err != nil {
					panic("Issue with converting post" + string(i) + "'s content into markdown")
				}
				fmt.Println(content)
			}
		}
	}
}

func main() {
	// get and parse command line flags
	flag.StringVar(&fBoard, "b", "g", "Board to grab results from")
	flag.IntVar(&fPageNum, "p", 1, "Page of threads")
	flag.IntVar(&fPostId, "i", 0, "Id of individual thread")
	flag.IntVar(&fPostOP, "op", 0, "Id of OP to search for a comment")
	flag.BoolVar(&fHelp, "h", false, "Print help dialouge")
	flag.Parse()

	// print help dialouge
	if fHelp == true {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// panic if the fPageNum is above the actual returned amount of
	// boards
	if fPageNum > MaxPages {
		panic("Too large of a page number! (range: 1-10)")
	}

	// make struct from board request
	request := apiHandler.GetBoardThreads(fBoard)

	// decalre a html -> markdown converter for
	// displaying content
	mdcov := md.NewConverter("", true, nil)

	// if a postid query is passed
	if fPostId != 0 {
		// check if an opriginal post id is passed
		if fPostOP != 0 {
			getPostById(fPostOP, fPostId, *mdcov)
		} else {
			fmt.Println("-i requires specifing an OP post with --op!")
		}
	} else if fThreadId != 0 {
		getPage(request, *mdcov)

		// fetch latest threads
	} else {
		getPage(request, *mdcov)
	}

	// list the threads of the specified page number
	// fmt.Println(request[fPageNum].Threads)

}
