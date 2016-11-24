package main

import (
  "./src/loading"
  "fmt"
  "flag"
  "strconv"
)

func main() {
  //load users
  users := loading.LoadUsers()

  //load participants
  parts := loading.LoadParts()

  //review command
  flag.Parse()

  //command
  com := flag.Arg(0)
  spec := flag.Arg(1)

  switch (com) {
    //User Commands
    case "user":
      switch (spec) {
        case "list":
          for _, user := range users.USERS {
            //Print Name
            fmt.Println("- " + user.FIRSTNAME + " " + user.SECONDNAME)
          }
        break
      }
    break

    //Participant Commands
    case "part":
      switch (spec) {
        case "list":
          for i, _ := range parts.PARTICIPANTS {
            //Get User
            user := users.USERS[i]

            //Print Name
            fmt.Println("- " + user.FIRSTNAME + " " + user.SECONDNAME)
          }
        break
        case "add":
          //Get User Id
          userId, _ := strconv.Atoi(flag.Arg(2))

          //Get User from users
          user := users.USERS[userId]

          //Add User to parts
          parts.PARTICIPANTS = append(parts.PARTICIPANTS, userId)

          //save Parts
          loading.SaveParts(parts)

          //Print Success
          fmt.Println("Added " + user.FIRSTNAME)
        break
        case "clear":
          //Clear Participants
          parts.PARTICIPANTS = make([]int, 0)

          //Save Parts
          loading.SaveParts(parts)

          //Succes Msg
          fmt.Println("Cleared all")
        break
        case "addall":
          //Clear Participants
          parts.PARTICIPANTS = make([]int, 0)

          //Add all Users
          for i, user := range users.USERS {
            //Add User
            parts.PARTICIPANTS = append(parts.PARTICIPANTS, i)

            //Print Name
            fmt.Println("- Added " + user.FIRSTNAME + " " + user.SECONDNAME)
          }

          //Save Parts
          loading.SaveParts(parts)

          //Succes Msg
          fmt.Println("Added all")
        break
        case "remove":
          //Get User Id
          userId, _ := strconv.Atoi(flag.Arg(2))

          //Get User from users
          user := users.USERS[userId]

          //Remove elements with value 'userId'
          for i := len(parts.PARTICIPANTS) - 1; i >= 0; i-- {
            part := parts.PARTICIPANTS[i]
            if part == userId {
              parts.PARTICIPANTS = append(parts.PARTICIPANTS[:i],
              parts.PARTICIPANTS[i+1:]...)
            }
          }

          //Save Parts
          loading.SaveParts(parts)

          //Succes Msg
          fmt.Println("Removed " + user.FIRSTNAME)
        break
      }
    break
  }

}
