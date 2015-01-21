// You can edit this code!
// Click here and start typing.
package main

import ( 
	"fmt"
	"strings"
	"regexp"
)


func main() {
	ps:=`It was seven o'clock of a very warm evening in the Seeonee hills when Father Wolf woke up from his day's rest, scratched himself, yawned, and spread out his paws one after the other to get rid of the sleepy feeling in their tips. Mother Wolf lay with her big gray nose dropped across her four tumbling, squealing cubs, and the moon shone into the mouth of the cave where they all lived. "Augrh!" said Father Wolf. "It is time to hunt again." He was going to spring down hill when a little shadow with a bushy tail crossed the threshold and whined: "Good luck go with you, O Chief of the Wolves. And good luck and strong white teeth go with noble children that they may never forget the hungry in this world."`
	fmt.Printf("%q\n",ps)
	re := regexp.MustCompile(`(([.?!])\s)`)
	fmt.Printf("%q\n",strings.Split(re.ReplaceAllString(ps, "$2|"),"|"))
}
