package main

import (
  "./src/loading"
  "./src/cmdlutil"
  "fmt"
  "flag"
  "strconv"
  "math/rand"
  "time"
  "sort"
)

var users loading.Users = loading.LoadUsers()
var parts loading.Parts = loading.LoadParts()
var comp loading.Comp = loading.LoadComp()




/////////////////////////////////////////////////////////////////////////////////////////////////////////////




//Create types for sorting by time since last by
type ByTime []int

func (s ByTime) Len() int {
    return len(s)
}

func (s ByTime) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByTime) Less(i, j int) bool {
    //Get User1
    user1 := users.USERS[s[i]]

    //Get User2
    user2 := users.USERS[s[j]]

    return (user1.BYTIME > user2.BYTIME)
}

//Create types for sorting by skill
type BySkill []int

func (s BySkill) Len() int {
    return len(s)
}

func (s BySkill) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s BySkill) Less(i, j int) bool {
    //Get User1
    user1 := users.USERS[s[i]]

    //Get User2
    user2 := users.USERS[s[j]]

    return (user1.SKILL < user2.SKILL)
}





////////////////////////////////////////////////////////////////////////////////////////////////////////////





func main() {

  //review command
  flag.Parse()

  //command
  com := flag.Arg(0)
  spec := flag.Arg(1)

  switch (com) {
    //create new comp listing
    case "comp":
      switch (spec) {
        case "won":
          //Addit info
          matchNum, _ := strconv.Atoi(flag.Arg(2))
          userNum, _ := strconv.Atoi(flag.Arg(3))

          //Subtract One (because of array indexing)
          matchNum -= 1
          userNum -= 1

          //Get Winners Name
          winnerId := comp.PARTICIPANTS[matchNum * 2 + userNum]
          winner := users.USERS[winnerId]
          winnerName := winner.FIRSTNAME + " " + winner.SECONDNAME

          //Get Losers Name
          loserId := comp.PARTICIPANTS[matchNum * 2 + (1 - userNum)]
          loser := users.USERS[loserId]
          loserName := loser.FIRSTNAME + " " + loser.SECONDNAME

          //Are you sure?
          if !cmdlutil.AskForConfirmation(winnerName + " beat " + loserName + "?") {
            return
          }

          //Effects of winning
          switch (comp.TYPE) {
            case "undefined":
              //Increase Skill of winner
              users.USERS[winnerId].SKILL += 1
            break
            case "elimination":
              //Increase Skill of winner
              users.USERS[winnerId].SKILL += 1

              //Remove loser from participants
              for i := len(parts.PARTICIPANTS) - 1; i >= 0; i-- {
                part := parts.PARTICIPANTS[i]
                if part == loserId {
                  parts.PARTICIPANTS = append(parts.PARTICIPANTS[:i],
                  parts.PARTICIPANTS[i+1:]...)
                }
              }

              //Save Parts
              loading.SaveParts(parts)

              //Succes Msg
              fmt.Println("Eliminated " + loserName)
            break
          }

          //Save Users
          loading.SaveUsers(users)

          //Has anyone won?
          if len(parts.PARTICIPANTS) == 1 {
            //Get Winner
            compWinner := users.USERS[parts.PARTICIPANTS[0]]

            //Print Winner
            fmt.Println("\n")
            fmt.Println(compWinner.FIRSTNAME + " " + compWinner.SECONDNAME + " has won the competition!")
            fmt.Println("\n")
          }
        break
        case "new":
          //what type?
          comp.TYPE = flag.Arg(2)

          //Reset Round Number
          comp.ROUNDNUM = 0

          //Save Comp
          loading.SaveComp(comp)
        break
        case "round":
          //Increase Round Num
          comp.ROUNDNUM += 1

          //Possible participants
          possibles := parts.PARTICIPANTS

          //Seed Random
          rand.Seed(time.Now().UnixNano())

          //is there an odd amount of participants?
          if len(parts.PARTICIPANTS) % 2 != 0 {

            //Random Sort Participants
            for i := range parts.PARTICIPANTS {
                j := rand.Intn(i + 1)
                parts.PARTICIPANTS[i], parts.PARTICIPANTS[j] = parts.PARTICIPANTS[j], parts.PARTICIPANTS[i]
            }

            //Actually Sort By Time
            sort.Sort(ByTime(parts.PARTICIPANTS))

            //Increase all bytimes
            for i, _ := range parts.PARTICIPANTS {
              //Print Name
              users.USERS[i].BYTIME += 1
            }

            //Reset By's bytime
            users.USERS[parts.PARTICIPANTS[0]].BYTIME = 0

            //Save Users
            loading.SaveUsers(users)

            //Remove
            possibles = possibles[1:]
          }

          //Sort Randomly
          for i := range possibles {
              j := rand.Intn(i + 1)
              possibles[i], possibles[j] = possibles[j], possibles[i]
          }

          //Sort by Skill Level
          sort.Sort(BySkill(possibles))

          //Title
          fmt.Println("\n")
          fmt.Println("\t  Round " + strconv.Itoa(comp.ROUNDNUM))

          //print
          i := 0
          for {
            //Break At End
            if i >= len(possibles)-1 {
             break
            }

            //Print match
            p1 := users.USERS[possibles[i]]
            p2 := users.USERS[possibles[i+1]]
            fmt.Println(strconv.Itoa(i/2 + 1) + " - " + p1.FIRSTNAME + " " + p1.SECONDNAME + " : " + p2.FIRSTNAME + " " + p2.SECONDNAME )

            //Skip
            i += 2
          }

          //Gap
          fmt.Println("\n")

          //Add possibles to comp DS
          comp.PARTICIPANTS = possibles

          //Save Comp
          loading.SaveComp(comp)
        break
      }
    break

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




/////////////////////////////////////////////////////////////////////////////////////////////////////////////
