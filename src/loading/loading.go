package loading

import (
  "./yaml.v2"
  "io/ioutil"
  "fmt"
)

const UserFile = "./dat/Users.yml"
const PartsFile = "./dat/Parts.yml"
const CompFile = "./dat/Comp.yml"

type Users struct {
  USERS []struct {
    FIRSTNAME string
    SECONDNAME string
    SKILL int
    BYTIME int
  }
}

type Parts struct {
  PARTICIPANTS []int
}

type Comp struct {
  TYPE string
  PARTICIPANTS []int
  ROUNDNUM int
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

func LoadComp() Comp {
  //Read Comp File
  bts, err := ioutil.ReadFile(CompFile)
  if err != nil {fmt.Println(err)}

  //Create 'Comp' data structure
  comp := Comp{}

  //Deserialize
  err = yaml.Unmarshal(bts, &comp)
  if err != nil {fmt.Println(err)}

  return comp
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

func SaveComp(comp Comp) {
  //serialize
  srl, _ := yaml.Marshal(&comp)

  //Write to file
  ioutil.WriteFile(CompFile, srl, 0644)
}

func SaveUsers(users Users) {
  //serialize
  srl, _ := yaml.Marshal(&users)

  //Write to file
  ioutil.WriteFile(UserFile, srl, 0644)
}
