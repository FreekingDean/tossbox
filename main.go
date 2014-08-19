package main

import (
  "fmt"
  "../reddit"
)

var comms chan *reddit.Comment
func main() {
  r, err := reddit.NewLoginSession("FuckTheCowboysBot", "im1coolguy", "FuckTheCowboysBot/v0.2 /u/bananaboydean /u/harkins")
  if err != nil { fmt.Println(err) }
  r_eagles, err := r.AboutSubreddit("eagles")
  if err != nil { fmt.Println(err) }
  comments, err := r_eagles.Comments(&r.Session, 5, "")
  if err != nil { fmt.Println(err) }
  for {
    for _, comment := range comments {
      fmt.Println(comment.Body)
      last = comment.time
    }
    comments, err := r_eagles.Comments(&r.Session, 5, last)
  }
}

    comments, err := subreddit.Comments(session, 10, last)
    if err != nil { fmt.Println(err) }
    for _, comment := range comments {
      comms <- comment
    }
  }
}
