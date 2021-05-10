package htmllinkparser

import (
    "io"
    "golang.org/x/net/html"
    "fmt"
    "strings"
)

type Link struct {
    Href string
    Text string
}

func Parse(fileReader io.Reader) ([]Link, error){
    doc, err := html.Parse(fileReader)
    if err != nil {
        return nil, err
    }

    linkNodes := findLinkNodes(doc)
    for _, node := range linkNodes {
        link := buildLink(node)
        fmt.Printf("\n---\nlink = %+v\n---\n", link)
    }
   
    return nil, nil
}

func findLinkNodes(node *html.Node) []*html.Node {
    if node.Type == html.ElementNode && node.Data == "a" {
        return []*html.Node{node}
    }

    var linkNodes []*html.Node
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        linkNodes = append(linkNodes, findLinkNodes(c)...)
    }

    return linkNodes
}

func getTextNodes(node *html.Node) string {
    if node.Type == html.TextNode {
        return node.Data
    }

    if node.Type != html.ElementNode {
        return ""
    }

   var text string
   for c := node.FirstChild; c != nil; c = c.NextSibling {
       text += getTextNodes(c) + " "
   }

   textFields := strings.Fields(text)
   return strings.Join(textFields, " ")
}

func buildLink(node *html.Node) Link {
    for _, attr := range node.Attr{
        if attr.Key == "href" {
            return Link{
                Href: attr.Val,
                Text: getTextNodes(node),
            }
        }
    }

    return Link{}
}
