package main

import (
  "fmt"
  "../reddit"
  "time"
  "strings"
)

func main() {
  r, err := reddit.NewLoginSession("FuckTheCowboysBot", "im1coolguy", "FuckTheCowboysBot/v0.2 /u/bananaboydean /u/harkins")
  if err != nil { fmt.Println(err) }
  bt, err := r.AboutSubreddit("nfl")
  if err != nil { fmt.Println(err) }
  last := ""
  body := ""
  for {
    comments, err := bt.Comments(&r.Session, 100, "", last)
    if err != nil { fmt.Println(err) }
    if len(comments) > 0 {
      last = comments[0].FullID
    }
    for _, comment := range comments {
      body = strings.ToLower(comment.Body)
      if strings.Contains(body, "cowboys") && !strings.Contains(comment.Author, "FuckTheCowboysBot") {
        fmt.Println(comment.Body)
        fmt.Println(comment.Author)
        r.Reply(comment, "Fuck the Cowboys")
        fmt.Println("Replied")
      }
    }
    time.Sleep(time.Second*15)
  }
}
