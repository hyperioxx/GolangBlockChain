package peering

import (
    "context"
    "fmt"
    "github.com/perlin-network/noise"
)




type PeeringCLient struct {

}


func (pc *PeeringCLient) Run(){
	node, err := noise.NewNode()
    if err != nil{
		fmt.Println()
	}

	defer node.Close()

    node.Handle(func(ctx noise.HandlerContext) error {
        if !ctx.IsRequest() {
            return nil
        }

        fmt.Printf("Got a message: '%s'\n", string(ctx.Data()))

        return ctx.Send([]byte("Hello World"))
    })


    node.Listen()

}




    

 