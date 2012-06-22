package rex

import (
  "strings"
  "fmt"
  "io"
  "encoding/xml"
)

func (rc *RexContext)String() string {
  lines := make([]string,len(rc.Nodes)+1+len(rc.Collections))
  lines[0] = xml.Header
  for i,v := range rc.Nodes {
    lines[i+1] = v.String()
  }
  for i,v := range rc.Collections {
    lines[i+len(rc.Nodes)+1] = v.String()
  }
  return strings.Join(lines,"")
}

func (rc *RexContext)Write(output io.Writer){
  fmt.Fprint(output, xml.Header)
  for _,v := range rc.Nodes {
    fmt.Fprint(output,v.String())
  }
  for _,v := range rc.Nodes {
    fmt.Fprint(output, v.String())
  }
}

func (rn *RexNode)String() string {
  attr_text := ""
  if len(rn.Attrs) > 0 {
    attrs := make([]string, len(rn.Attrs))
    for attr,value := range rn.Attrs {
      attrs = append(attrs, fmt.Sprintf(`%s="%s"`,attr, fmt.Sprint(value)))
    }
    attr_text = strings.Join(attrs, " ")
  }

  switch {
  case len(rn.Nodes) == 0 && rn.Content == "":
    return fmt.Sprintf("<%s%s />",rn.Tag,attr_text)
  case len(rn.Nodes) == 0 && rn.Content != "":
    return fmt.Sprintf("<%s%s>%s</%s>",rn.Tag,attr_text,rn.Content,rn.Tag)
  case len(rn.Nodes)+len(rn.Collections) > 0 && rn.Content == "":
    nodes := make([]string, len(rn.Nodes)+len(rn.Collections))
    for i,node := range rn.Nodes {
      nodes[i] = node.String()
    }

    for i := 0; i < len(rn.Collections); i += 1 {
      nodes[i+len(rn.Nodes)] = rn.Collections[i].String()

    }

    return fmt.Sprintf("<%s%s>%s</%s>",rn.Tag,attr_text,strings.Join(nodes, "\n"),rn.Tag)
  }
  return ""
}

func (rc *RexCollection)String() string {
  nodes := make([]string, len(rc.Nodes))
  for i,v := range rc.Nodes {
    nodes[i] = v.String()
  }

  return fmt.Sprintf("<%s>%s</%s>",rc.Name,strings.Join(nodes,"\n"),rc.Name)
}
