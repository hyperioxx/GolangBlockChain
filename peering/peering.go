package peering

import (
    "context"
    "fmt"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
)


var Client *PeeringCLient


func init(){
	Client = new(PeeringCLient)
	go Client.Run()
}




type PeeringCLient struct {
	node *noise.Node
	k *kademlia.Protocol
}


func (pc *PeeringCLient) Run(){
	pc.node, _ = noise.NewNode()
	pc.k = kademlia.New()
	pc.node.Bind(pc.k.Protocol())
	
    

	defer pc.node.Close()

    pc.node.Handle(func(ctx noise.HandlerContext) error {
        if !ctx.IsRequest() {
            return nil
        }

        fmt.Printf("Got a message: '%s'\n", string(ctx.Data()))

        return ctx.Send([]byte("Hello World"))
	})
	
	err := pc.node.Listen()
	if err != nil {
		fmt.Println(err)
	}

}


func (pc *PeeringCLient) Send(msg string){
	fmt.Println(pc.node.Addr())

	pc.node.Request(context.TODO(),pc.node.Addr() , []byte(msg))

}




    

 