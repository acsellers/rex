package rex

import "fmt"

type RexContext struct {
  Nodes  []*RexNode
  Collections []*RexCollection
}

type RexCollection struct {
  Name string
  NodeName string
  Nodes []*RexNode
}

type RexNode struct {
  Tag         string
  Content     string
  Nodes       []*RexNode
  Attrs       map[string]interface{}
  Collections []*RexCollection
}

func GetContext() *RexContext {
  rc := new(RexContext)
  rc.Nodes = make([]*RexNode, 0)
  return rc
}

func (rc *RexContext)Item(name string, attrs ...interface{}) *RexNode {
  rn := NewRexNode(name,attrs)
  rc.Nodes = append(rc.Nodes,rn)
  return rn
}

func (rn *RexNode)Item(name string, attrs ...interface{}) *RexNode {
  new_rn := NewRexNode(name, attrs)
  rn.Nodes = append(rn.Nodes,new_rn)
  return new_rn
}

func (rn *RexContext)Collection(name,instances string) *RexCollection {
  rc := NewRexCollection(name,instances)
  rn.Collections = append(rn.Collections,rc)
  return rc
}

func (rn *RexNode)Collection(collection_name, collection_type string) *RexCollection {
  rc := NewRexCollection(collection_name,collection_type)
  rn.Collections = append(rn.Collections,rc)
  return rc
}

func (rc *RexCollection)Instance(attr ...interface{}) *RexNode {
  rn := NewRexNode(rc.NodeName,attr)
  rc.Nodes = append(rc.Nodes, rn)
  return rn
}

func NewRexNode(tag string, attrs []interface{}) *RexNode {
  rn      := new(RexNode)
  rn.Nodes = make([]*RexNode,0)
  rn.Attrs = make(map[string]interface{})
  rn.Tag   = tag

  switch {
  case len(attrs) == 1:
    rn.Content = fmt.Sprint(attrs[0])
  case len(attrs) % 2 == 0:
    for i := 0; i + 1 < len(attrs); i += 2 {
      rn.Attrs[fmt.Sprint(attrs[i])] = attrs[i+1]
    }
  case len(attrs) % 2 == 1:
    rn.Content = fmt.Sprint(attrs[0])
    for i := 1; i + 1 < len(attrs); i += 2 {
      rn.Attrs[fmt.Sprint(attrs[i])] = attrs[i+1]
    }
  }

  return rn
}

func NewRexCollection(name,instance string) *RexCollection {
  rc := new(RexCollection)
  rc.Nodes = make([]*RexNode,0)
  rc.Name = name
  rc.NodeName = instance

  return  rc
}
