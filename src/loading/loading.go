package loading

import (
  "./yaml.v2"
  "io/ioutil"
  "fmt"
)

const UserFile = "./dat/Users.yml"
const PartsFile = "./dat/Parts.yml"

type Users struct {
  USERS []struct {
    FIRSTNAME string
    SECONDNAME string
    SKILL int
  }
}

type Parts struct {
  PARTICIPANTS []int
}

func LoadUsers() Users {
  //Read Users File
  bts, err := ioutil.ReadFile(UserFile)
  if err != nil {fmt.Println(err)}

  //Create 'Users' data structure
  users := Users{}

  //Deserialize
  err = yaml.Unmarshal(bts, &users)
  if err != nil {fmt.Println(err)}

  return users
}

func LoadParts() Parts {
  //Read Parts File
  bts, err := ioutil.ReadFile(PartsFile)
  if err != nil {fmt.Println(err)}

  //Create 'Parts' data structure
  parts := Parts{}

  //Deserialize
  err = yaml.Unmarshal(bts, &parts)
  if err != nil {fmt.Println(err)}

  return parts
}

func SaveParts(parts Parts) {
  //serialize
  srl, _ := yaml.Marshal(&parts)

  //Write to file
  ioutil.WriteFile(PartsFile, srl, 0644)
}

func SaveUsers(users Users) {
  //serialize
  srl, _ := yaml.Marshal(&users)

  //Write to file
  ioutil.WriteFile(UserFile, srl, 0644)
}
