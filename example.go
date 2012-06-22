package main

import (
  "rex"
  "fmt"
)

func main(){
  users := []string{"andrew","josh","aaron","mario","anthony","ron"}
  xml := rex.GetContext()

  school_section := xml.Item("school","user_count",6)
  school_section.Item("name","University of Learningness")
  users_section := school_section.Collection("users","user")

  for _,v := range users {
    user_node := users_section.Instance()
    user_node.Item("name",v)
  }

  fmt.Println(xml.String())

}
